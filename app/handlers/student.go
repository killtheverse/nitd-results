package handlers

import (
	// "context"
	"encoding/json"
	"fmt"
	"net/http"

	// "time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/go-playground/validator/v10"

	logger "github.com/killtheverse/nitd-results/app/logging"
	"github.com/killtheverse/nitd-results/app/models"
	"github.com/killtheverse/nitd-results/app/utils"
)

func GetStudents(db *mongo.Database, rw http.ResponseWriter, request *http.Request) {
	studentList := []int{1, 2,3 }
	utils.ResponseWriter(rw, http.StatusOK, "", studentList)
}

func CreateStudent(db *mongo.Database, rw http.ResponseWriter, request *http.Request) {
	student := new(models.Student)
	err := json.NewDecoder(request.Body).Decode(student)
	if err != nil {
		utils.ResponseWriter(rw, http.StatusBadRequest, "Invalid JSON body", nil)
		logger.Write("[ERROR]: %v", err)
		return
	}
	fmt.Println(student)
	utils.ResponseWriter(rw, http.StatusCreated, "Created Student", student)
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// result, err := db.Collection("students").InsertOne(ctx, student)
	// if err != nil {
	// 	switch err.(type) {
	// 	case mongo.WriteException:
	// 		ResponseWriter(response_writer, http.StatusNotAcceptable, "Can't write in the database", nil)
	// 	default:
	// 		ResponseWriter(response_writer, http.StatusInternalServerError, "Error occured", nil)
	// 	}
	// }
}

func UpdateStudent(db *mongo.Database, rw http.ResponseWriter, request *http.Request) {
	// Parse request body and extract the student
	student := new(models.Student)
	err := json.NewDecoder(request.Body).Decode(student)
	if err != nil {
		utils.ResponseWriter(rw, http.StatusBadRequest, "Invalid JSON body", nil)
		logger.Write("[ERROR]: %v", err)
		return
	}

	validate := validator.New()
	err = validate.Struct(student)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		responseBody := map[string]string{"error": validationErrors.Error()}
		utils.ResponseWriter(rw, http.StatusUnprocessableEntity, "Errors in validation", responseBody)
	}

	// filter := bson.D{
	// 	{"roll_no": student.Roll}
	// }

	// // Find the student and update it
	// ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	// defer cancel()
	// err = db.Collection("students").FindOneAndReplace(ctx, )
}