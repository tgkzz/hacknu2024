package repository

import (
	"backend/internal/repository/student"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	student.IStudentRepo
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		IStudentRepo: student.NewStudentRepo(db),
	}
}
