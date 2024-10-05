package models

type RecordDTO struct {
	Email  string `json:"email" binding:"required"`
	Record int    `json:"record" binding:"required"`
}
