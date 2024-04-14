package student

import (
	"backend/internal/handler/dto"
	"backend/internal/models"
	"backend/internal/repository/student"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"regexp"
	"strings"
)

type StudentService struct {
	repo         student.IStudentRepo
	openaiClient *openai.Client
}

type IStudentService interface {
	CreateNewStudent(ctx context.Context, student models.Student) (string, error)
	GetUserIdByName(ctx context.Context, name string) (string, error)
	AnswerStudentReq(ctx context.Context, req dto.GetStudentQuestion) (interface{}, error)
}

func NewStudentService(repo student.IStudentRepo, client *openai.Client) *StudentService {
	return &StudentService{
		repo:         repo,
		openaiClient: client,
	}
}

func (s *StudentService) CreateNewStudent(ctx context.Context, student models.Student) (string, error) {
	return s.repo.CreateStudent(ctx, student)
}

func (s *StudentService) GetUserIdByName(ctx context.Context, name string) (string, error) {
	return s.repo.GetUserIdByName(ctx, name)
}

func (s *StudentService) GetUserById(ctx context.Context, id string) (models.Student, error) {
	return s.repo.GetUserById(ctx, id)
}

func (s *StudentService) AnswerStudentReq(ctx context.Context, req dto.GetStudentQuestion) (interface{}, error) {
	sampleData, err := s.GetUserById(ctx, "661ab8fb43099c09847e29a8")
	if err != nil {
		return nil, err
	}

	systemQuery := fmt.Sprintf("you are expert in golang and mongo driver for golang. for my next requests you should return only projection and filter of data by which i will find required data, and no natural language.  given the sample document from student table %v. it represents student table and their academic perfomance in representation of conf_score for each date. In the collection itself there are many other documents that have same structure. return me projection and filter and no natural language in format filter := bson.M{} projection := bson.M{}", sampleData)

	userQuery := fmt.Sprintf("In my table I have next fields name represents name and surname of the student, teacher_ids represents teacher ids that he takes class in, subjects represents subject that he is studying, modules represent modules that this subject have, and activity_time is time that he have spent for this lesson, and conf_score represents the score that he has been given on platform")

	resp, err := s.openaiClient.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemQuery,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userQuery + ". " + req.Question,
				},
			},
		})
	if err != nil {
		return nil, err
	}
	log.Print(resp.Choices[0].Message.Content)

	filter, projection, err := s.getFilterAndFilter(resp.Choices[0].Message.Content)
	if err != nil {
		return nil, err
	}

	res, err := s.repo.GetByFilterAndProjection(ctx, filter, projection)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *StudentService) getFilterAndFilter(text string) (bson.M, bson.M, error) {
	re := regexp.MustCompile(`(?m)(filter|projection) *:= *bson\.M\{([^}]*)\}`)
	matches := re.FindAllStringSubmatch(text, -1)
	if len(matches) < 2 {
		log.Print(text)
		return nil, nil, fmt.Errorf("filter or projection not found in the text")
	}

	var filterStr, projectionStr string
	for _, match := range matches {
		key := match[1]
		value := "{" + strings.TrimSpace(match[2]) + "}"
		if key == "filter" {
			filterStr = value
		} else if key == "projection" {
			projectionStr = value
		}
	}

	var filter bson.M
	if err := bson.UnmarshalExtJSON([]byte(filterStr), true, &filter); err != nil {
		return nil, nil, fmt.Errorf("failed to parse filter: %v", err)
	}

	var projection bson.M
	if err := bson.UnmarshalExtJSON([]byte(projectionStr), true, &projection); err != nil {
		return nil, nil, fmt.Errorf("failed to parse projection: %v", err)
	}

	return filter, projection, nil
}
