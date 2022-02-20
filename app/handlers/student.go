package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	logger "github.com/killtheverse/nitd-results/app/logging"
	"github.com/killtheverse/nitd-results/app/models"
	"github.com/killtheverse/nitd-results/app/utils"
)

// swagger:route GET /students students listStudents
// Returns a list of students filtered by parameters
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http, https
//
// Responses:
//	default: ErrorResponse
//	200: PaginatedResponse

// GetStudents returns a list of students based on the search query parameters
func GetStudents(db *mongo.Database, rw http.ResponseWriter, request *http.Request) {
	// Extract query parameters
	var params = request.URL.Query()
	limitString := params.Get("limit")
	offsetString := params.Get("offset")
	batch := params.Get("batch")
	branch := params.Get("branch")
	program := params.Get("program")
	
	// If limit and offset not present or invalid, then assign default values
	limit, err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		limit = 100
	}
	if limit>100 {
		limit = 100
	}

	offset, err := strconv.ParseInt(offsetString, 10, 64)
	if err != nil {
		offset = 0
	}

	// Create filter
	batchString := "^" + batch
	branchString := "^" + branch
	programString := "^" + program


	filter := bson.D{
	primitive.E{Key: "roll_no", Value: bson.D{primitive.E{Key: "$regex", Value: batchString}}},
	primitive.E{Key: "branch", Value: bson.D{primitive.E{Key: "$regex", Value: branchString}, {Key: "$options", Value: "i"}}},
	primitive.E{Key: "program", Value: bson.D{{Key: "$regex", Value: programString}, {Key: "$options", Value: "i"}}},
	}

	opts := options.FindOptions{
		Skip: &offset,
		Limit: &limit,
		Sort: bson.M{
			"roll_no": 1,
		},
	}

	// Query the database
	var students []models.Student
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	cursor, err := db.Collection("students").Find(ctx, filter, &opts)
	if err != nil {
		logger.Write("[ERROR]: While quering collection - %s", err)
		utils.ErrorResponseWriter(rw, http.StatusInternalServerError, "Error occured while reading data", nil)
		return
	}
	err = cursor.All(context.Background(), &students)
	if err != nil {
		logger.Write("[ERROR] in cursor - %s", err)
		utils.ErrorResponseWriter(rw, http.StatusInternalServerError, "Error occured while reading data", nil)
		return
	}

	// Write a paginated response
	var nextOffset int64 = offset + limit
	var nextLimit int64 = limit
	var nextParams url.Values = params
	var prevParams url.Values = params
	nextParams.Set("offset", strconv.Itoa(int(nextOffset)))
	nextParams.Set("limit", strconv.Itoa(int(nextLimit)))
	request.URL.RawQuery = params.Encode()
	var next string = request.URL.String()
	
	var prevOffset int64 = offset - limit
	if prevOffset < 0 {
		prevOffset = 0
	}
	var prevLimit int64 = limit
	prevParams.Set("offset", strconv.Itoa(int(prevOffset)))
	prevParams.Set("limit", strconv.Itoa(int(prevLimit)))
	request.URL.RawQuery = params.Encode()
	var prev string = request.URL.String()
	
	count := len(students)
	utils.PaginatedResponseWriter(rw, http.StatusOK, "Retrieved students list", count, next, prev, students)
}

// swagger:route GET /students/{roll_number} students studentDetail
// Returns information about a particular student
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http, https
//
// Responses:
//	default: ErrorResponse
//	200: Response

// GetStudent returns information about a particular student
func GetStudent(db *mongo.Database, rw http.ResponseWriter, request *http.Request) {
	// Extract roll number from request URI
	var params = mux.Vars(request)
	roll_no := params["roll_number"]
	
	// Find the student document
	var student models.Student
	filter := bson.M{"roll_no": roll_no}
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err := db.Collection("students").FindOne(ctx, filter).Decode(&student)
	if err == mongo.ErrNoDocuments {
		logger.Write("No document found for roll number: %s.", student.Roll)
		utils.ErrorResponseWriter(rw, http.StatusNotFound, "No document exists", nil)
		return
	} else if err == nil{
		logger.Write("Document found for %s", student.Roll)
		utils.ResponseWriter(rw, http.StatusOK, "Found the document", student)
		return
	} else {
		logger.Write("[ERROR]: %s", err)
		utils.ErrorResponseWriter(rw, http.StatusInternalServerError, "Error occured while looking up for document", nil)
		return
	}
}

// swagger:route PUT /students/{roll_number} students updateStudent
// Updates the student if it exists, otherwise creates a new entry
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http, https
//
// Responses:
//	default: ErrorResponse
//	201: Response

// UpdateStudent updates or creates a student in the database
func UpdateStudent(db *mongo.Database, rw http.ResponseWriter, request *http.Request) {
	// Extract roll number from request URI
	var params = mux.Vars(request)
	roll_no := params["roll_number"]

	// Parse request body and extract the student
	var student models.Student
	err := json.NewDecoder(request.Body).Decode(&student)
	if err != nil {
		utils.ErrorResponseWriter(rw, http.StatusBadRequest, "Invalid JSON body", nil)
		logger.Write("[ERROR]: Error in unmarshalling student object - %v", err)
		return
	}

	// Validate the parsed student structure
	validate := validator.New()
	err = validate.Struct(student)
	if err != nil {
		logger.Write("[ERROR]: Errors in validating student")
		validationErrors := err.(validator.ValidationErrors)
		responseBody := map[string]string{"error": validationErrors.Error()}
		utils.ErrorResponseWriter(rw, http.StatusUnprocessableEntity, "Errors in validation", responseBody)
		return
	}

	if student.Roll != roll_no {
		logger.Write("Roll number in payload does not match roll number in URI")
		utils.ErrorResponseWriter(rw, http.StatusBadRequest, "Invalid roll number in payload", nil)
		return
	}

	var existing_student models.Student
	var existing_semesters []models.Semester
	filter := bson.M{"roll_no": student.Roll}

	// Check if the student already exists in record
	var exists bool = true
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err = db.Collection("students").FindOne(ctx, filter).Decode(&existing_student)
	if err == mongo.ErrNoDocuments {
		exists = false
		logger.Write("No document found for roll number: %s. Creating new document", student.Roll)
		student.CreatedAt = time.Now()
	} else if err == nil{
		logger.Write("Document already exists for roll number: %s. Updating", student.Roll)
		student.CreatedAt = existing_student.CreatedAt
		existing_semesters = existing_student.Semesters
	} else {
		logger.Write("[ERROR]: %s", err)
		utils.ErrorResponseWriter(rw, http.StatusInternalServerError, "Error occured while looking up for existing documents", nil)
		return
	}
	student.UpdatedAt = time.Now()

	// Include any semester which already exists in the database but is missing in new body
	for _, existing_sem := range existing_semesters {
		var present bool = false
		for _, new_sem := range student.Semesters {
			if existing_sem.Number == new_sem.Number {
				present = true
				break
			}
		}
		if !present {
			student.Semesters = append(student.Semesters, existing_sem)
		}
	}

	// Update the document for student
	ctx, cancel = context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if exists {
		_, err = db.Collection("students").ReplaceOne(ctx, filter, student)
	} else {
		_, err = db.Collection("students").InsertOne(ctx, student)
	}
	
	if err != nil {
		logger.Write("[ERROR] Updating docuement - %s", err)
		utils.ErrorResponseWriter(rw, http.StatusInternalServerError, "Error occured while updating document", nil)
	} else {
		var message string
		if exists {
			message = "Successfully updated student"
		} else {
			message = "Successfully created student"
		}
		logger.Write("Successfully updated document for %s", student.Roll)
		utils.ResponseWriter(rw, http.StatusCreated, message, student)
	}
}
