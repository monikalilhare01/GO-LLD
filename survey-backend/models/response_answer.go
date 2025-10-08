package models

type ResponseAnswer struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	QuestionID int    `json:"question_id"`
	ResponseID int    `json:"response_id"`
	Response   string `json:"response"`
}
