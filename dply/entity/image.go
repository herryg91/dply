package entity

import "time"

type ContainerImage struct {
	Id             int        `json:"id"`
	Digest         string     `json:"digest"`
	Image          string     `json:"image"`
	Project        string     `json:"project"`
	RepositoryName string     `json:"repository_name"`
	Description    string     `json:"description"`
	CreatedBy      int        `json:"created_by"`
	CreatedAt      *time.Time `json:"created_at"`
	Notes          string     `json:"notes"`
}
