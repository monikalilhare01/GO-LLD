package models

type Questions struct {
	ID           int    `json:"id"`
	SurveyID     int    `json:"survey_id"`
	Questions    string `json:"question"`
	QuestionType string `json:"question_type"`
	IsRequired   bool   `json:"is_required"`
	Options      string `json:"options"`
}
