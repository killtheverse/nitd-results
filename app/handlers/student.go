package handlers

import (
	// "context"
	"encoding/json"
	"fmt"
	"net/http"
	// "time"

	"go.mongodb.org/mongo-driver/mongo"

	logger "github.com/killtheverse/nitd-results/app/logging"
	"github.com/killtheverse/nitd-results/app/models"
	"github.com/killtheverse/nitd-results/app/utils"
)

func GetStudents(db *mongo.Database, response_writer http.ResponseWriter, request *http.Request) {
	studentList := []int{1, 2,3 }
	utils.ResponseWriter(response_writer, http.StatusOK, "", studentList)
}

func CreateStudent(db *mongo.Database, response_writer http.ResponseWriter, request *http.Request) {
	student := new(models.Student)
	err := json.NewDecoder(request.Body).Decode(student)
	if err != nil {
		utils.ResponseWriter(response_writer, http.StatusBadRequest, "Invalid JSON body", nil)
		logger.Write("[ERROR]: %v", err)
		return
	}
	fmt.Println(student)
	utils.ResponseWriter(response_writer, http.StatusCreated, "Created Student", student)
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