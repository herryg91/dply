package port_usecase

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply/app/repository"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/editor"
)

type usecase struct {
	repo repository.SpecRepository
}

func New(repo repository.SpecRepository) UseCase {
	return &usecase{repo: repo}
}

func (uc *usecase) Get(project, env, name string) (*entity.Port, error) {
	resp, err := uc.repo.GetPort(project, env, name)
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return nil, fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return resp, nil
}

func (uc *usecase) Upsert(data entity.Port) error {
	err := uc.repo.UpsertPort(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}

func (uc *usecase) UpsertViaEditor(project, env, name string, editorApp editor.EditorApp) (bool, error) {
	// Get Current Data
	getResp, err := uc.Get(project, env, name)
	if err != nil {
		return false, err
	}
	currentData, _ := json.MarshalIndent(getResp, "", "    ")

	// Get Updated Data via Editor
	updatedData, err := editor.Open(editorApp, "tmp_port_edit", currentData)
	if err != nil {
		return false, fmt.Errorf("%w: %v", ErrUnexpected, "Error on editor: "+err.Error())
	}

	// if nothing to change
	if string(currentData) == string(updatedData) {
		return false, nil
	}

	data := entity.Port{Project: project, Env: env, Name: name, Ports: []entity.PortSpec{}}
	err = json.Unmarshal(updatedData, &data)
	if err != nil {
		return false, fmt.Errorf("%w: %v", ErrUnexpected, "Error unmarshal: "+string(updatedData))
	}

	// Upsert to server
	err = uc.Upsert(data)
	if err != nil {
		return false, err
	}
	return true, nil
}
