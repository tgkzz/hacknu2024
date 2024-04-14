package student

import (
	"backend/internal/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const StudentCollection = "student"

type StudentRepo struct {
	DB *mongo.Database
}

type IStudentRepo interface {
	CreateStudent(ctx context.Context, student models.Student) (string, error)
	GetUserIdByName(ctx context.Context, name string) (string, error)
	GetUserById(ctx context.Context, id string) (models.Student, error)
	GetByFilterAndProjection(ctx context.Context, filter, projection bson.M) (models.Student, error)
}

func NewStudentRepo(db *mongo.Database) *StudentRepo {
	return &StudentRepo{
		DB: db,
	}
}

func (s *StudentRepo) CreateStudent(ctx context.Context, student models.Student) (string, error) {
	coll := s.DB.Collection(StudentCollection)

	res, err := coll.InsertOne(ctx, student)
	if err != nil {
		return "", err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", errors.New("could not return student id")
}

func (s *StudentRepo) GetUserById(ctx context.Context, id string) (models.Student, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Student{}, err
	}

	coll := s.DB.Collection(StudentCollection)

	filter := bson.M{"_id": oid}

	var res models.Student
	if err := coll.FindOne(ctx, filter).Decode(&res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Student{}, models.ErrNoDocument
		}
		return models.Student{}, err
	}

	return res, nil
}

func (s *StudentRepo) GetUserIdByName(ctx context.Context, name string) (string, error) {
	var res struct {
		Id string `bson:"_id"`
	}

	coll := s.DB.Collection(StudentCollection)

	filter := bson.M{"name": name}

	findOptions := options.FindOne().SetProjection(bson.M{"_id": 1})

	if err := coll.FindOne(ctx, filter, findOptions).Decode(&res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", models.ErrNoDocument
		}
		return "", err
	}

	return res.Id, nil
}

func (s *StudentRepo) GetByFilterAndProjection(ctx context.Context, filter, projection bson.M) (models.Student, error) {
	coll := s.DB.Collection(StudentCollection)

	opts := options.FindOne().SetProjection(projection)

	var res models.Student
	if err := coll.FindOne(ctx, filter, opts).Decode(&res); err != nil {
		return models.Student{}, err
	}

	return res, nil
}
