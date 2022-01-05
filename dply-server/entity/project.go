package entity

import (
	"errors"
	"regexp"
)

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var regex_name = regexp.MustCompile("^[a-zA-Z0-9_-]*$")

func (p *Project) ValidateName() error {
	ok := regex_name.MatchString(p.Name)
	if !ok {
		return errors.New("name should be alphanumeric / underscore (_) / dash (-)")
	}
	return nil
}
