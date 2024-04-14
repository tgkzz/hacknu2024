package dto

type GetStudentQuestion struct {
	UserID   string `json:"user_id"`
	Question string `json:"question"`
}
