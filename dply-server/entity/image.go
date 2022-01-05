package entity

import (
	"time"
)

type Image struct {
	Id          int        `json:"id"`
	Digest      string     `json:"digest"`
	Image       string     `json:"image"`
	Project     string     `json:"project"`
	Repository  string     `json:"repository"`
	Description string     `json:"description"`
	CreatedBy   int        `json:"created_by"`
	CreatedAt   *time.Time `json:"created_at"`
}
