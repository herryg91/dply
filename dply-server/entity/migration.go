package entity

import "time"

type Migration struct {
	Id        int        `json_name:"id"`
	Name      string     `json_name:"name"`
	CreatedAt *time.Time `json_name:"created_at"`
	UpdatedAt *time.Time `json_name:"updated_at"`
}
