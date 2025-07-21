package models

import "time"

type Survey struct {
	ID                int         `json:"id"`
	SurveyName        string      `json:"survey_name"`
	SurveyDescription string      `json:"survey_description"`
	Questions         []Questions `json:"questions"`
	CreatedAt         time.Time   `json:"created_at"`
}
