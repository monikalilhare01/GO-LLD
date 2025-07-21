package models

import "time"

type Response struct {
	ID             int              `json:"id"`
	SurveyID       int              `json:"surver_id"`
	QuestionID     int              `json:"question_id"`
	UserID         int              `json:"user_id"`
	ResponseAnswer []ResponseAnswer `json:"response_answer"`
	SubmittedAt    time.Time        `json:"submitted_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}
