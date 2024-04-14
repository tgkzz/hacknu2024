package main

import (
	"backend/config"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/service"
	"context"
	"github.com/go-faker/faker/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx := context.Background()

	cfg, err := config.LoadConfig("app.env")
	if err != nil {
		log.Print(err)
		return
	}

	log.Print("connecting to db...")
	db, err := repository.LoadDB(cfg.MongoDB)
	if err != nil {
		log.Print(err)
		return
	}

	students := []models.Student{
		{
			Id:         primitive.NewObjectID(),
			Name:       faker.FirstName(),
			TeacherIds: []string{},
			Subjects: []models.Subject{
				{
					Name: "Math",
					Modules: []models.Module{
						generateModuleForWeek("Algebra", 2024, time.April),
						generateModuleForWeek("Geometry", 2024, time.April),
						generateModuleForWeek("Calculus", 2024, time.April),
					},
				},
				{
					Name: "Chemistry",
					Modules: []models.Module{
						generateModuleForWeek("Organic Chemistry", 2024, time.April),
						generateModuleForWeek("BioChemistry", 2024, time.April),
					},
				},
				{
					Name: "Music",
					Modules: []models.Module{
						generateModuleForWeek("Classical Music", 2024, time.April),
						generateModuleForWeek("Bethoven", 2024, time.April),
					},
				},
			},
		},
		{
			Id:         primitive.NewObjectID(),
			Name:       faker.FirstName(),
			TeacherIds: []string{},
			Subjects: []models.Subject{
				{
					Name: "Physics",
					Modules: []models.Module{
						generateModuleForWeek("Mechanic", 2024, time.April),
						generateModuleForWeek("AstroPhysics", 2024, time.April),
					},
				},
				{
					Name: "PE",
					Modules: []models.Module{
						generateModuleForWeek("Basketball", 2024, time.April),
						generateModuleForWeek("Football", 2024, time.April),
					},
				},
			},
		},
		{
			Id:         primitive.NewObjectID(),
			Name:       faker.FirstName(),
			TeacherIds: []string{},
			Subjects: []models.Subject{
				{
					Name: "Informatics",
					Modules: []models.Module{
						generateModuleForWeek("C++", 2024, time.April),
						generateModuleForWeek("go", 2024, time.April),
					},
				},
				{
					Name: "History",
					Modules: []models.Module{
						generateModuleForWeek("Kazakhstan history", 2024, time.April),
						generateModuleForWeek("Usa history", 2024, time.April),
						generateModuleForWeek("Modern history", 2024, time.April),
					},
				},
			},
		},
		{
			Id:         primitive.NewObjectID(),
			Name:       faker.FirstName(),
			TeacherIds: []string{},
			Subjects: []models.Subject{
				{
					Name: "Design",
					Modules: []models.Module{
						generateModuleForWeek("Pablo Picasso", 2024, time.April),
						generateModuleForWeek("PostModern", 2024, time.April),
					},
				},
			},
		},
	}

	r := repository.NewRepository(db)

	s := service.NewService(*r, cfg.OpenAiToken)

	for _, student := range students {
		studentId, err := s.Student.CreateNewStudent(ctx, student)
		if err != nil {
			log.Print(err)
		}

		log.Printf("%s\n", studentId)
	}

	log.Print("action has been end")
}

func generateModuleForWeek(moduleName string, year int, month time.Month) models.Module {
	rand.Seed(time.Now().UnixNano())

	activities := []models.Activity{}
	totalScore := 0.0

	for day := 1; day <= 7; day++ {
		date := time.Date(year, month, day, 13, 0, 0, 0, time.UTC)
		activityTime := rand.Float64() * 8
		if rand.Float64() < 0.33 {
			activityTime /= 8
		}
		confidenceScore := randomizeConfScore(activityTime)
		activities = append(activities, models.Activity{
			ActivityTime:        activityTime,
			Date:                date,
			ConfidenceScore:     confidenceScore,
			СumulativeСonfScore: totalScore + confidenceScore,
		})
		totalScore += confidenceScore
	}
	return models.Module{
		Name:       moduleName,
		Activities: activities,
	}
}

func randomizeConfScore(activityTime float64) float64 {

	if activityTime < 1 {
		return -5 - rand.Float64()*(15-5)
	} else {
		return 10 + rand.Float64()*(70-10)
	}
}
