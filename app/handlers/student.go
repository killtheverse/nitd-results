package handlers

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/killtheverse/nitd-results/app/models"
)

func GetStudents(db *mongo.Database, rw http.ResponseWriter, r *http.Request) {
	var studentList []models.Student
	ResponseWriter(rw, http.StatusOK, "", studentList)
}