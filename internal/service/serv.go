package service

import (
	"backend/client"
	"backend/internal/repository"
	"backend/internal/service/student"
)

type Service struct {
	Student student.IStudentService
}

func NewService(repo repository.Repository, openaiToken string) *Service {
	return &Service{
		Student: student.NewStudentService(repo, client.NewOpenaiClient(openaiToken)),
	}
}
