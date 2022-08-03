package deployment_config_usecase

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

func (uc *usecase) Get(project, env, name string) (*entity.DeploymentConfig, error) {
	resp, err := uc.repo.GetDeploymentConfig(project, env, name)
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return nil, fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return resp, nil
}

func (uc *usecase) Upsert(data entity.DeploymentConfig) error {
	err := uc.repo.UpsertDeploymentConfig(data)
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

	current := map[string]interface{}{
		"liveness_probe":  getResp.LivenessProbe,
		"readiness_probe": getResp.ReadinessProbe,
		"startup_probe":   getResp.StartupProbe,
	}

	currentData, _ := json.MarshalIndent(current, "", "    ")

	// Get Updated Data via Editor
	updatedData, err := editor.Open(editorApp, "tmp_dc_edit", currentData)
	if err != nil {
		return false, fmt.Errorf("%w: %v", ErrUnexpected, "Error on editor: "+err.Error())
	}

	// if nothing to change
	if string(currentData) == string(updatedData) {
		return false, nil
	}

	data := entity.DeploymentConfig{Project: project, Env: env, Name: name, LivenessProbe: nil, ReadinessProbe: nil, StartupProbe: nil}
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
