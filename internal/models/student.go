package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Student struct {
	Id         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `bson:"name" json:"name" faker:"first_name"`
	TeacherIds []string           `bson:"teacher_ids" json:"teacher_ids"`
	Subjects   []Subject          `bson:"subjects" json:"subjects"`
}

type Subject struct {
	Name    string   `bson:"name" json:"name"`
	Modules []Module `bson:"modules" json:"modules"`
}

type Module struct {
	Name       string     `bson:"name" json:"name"`
	Activities []Activity `bson:"activities" json:"activities"`
}

type Activity struct {
	ActivityTime        float64   `bson:"activity_time" json:"activity_time"`
	Date                time.Time `bson:"date" json:"date"`
	ConfidenceScore     float64   `bson:"conf_score" json:"conf_score"`
	СumulativeСonfScore float64   `bson:"cumulative_conf_score" json:"cumulative_conf_score"`
}
