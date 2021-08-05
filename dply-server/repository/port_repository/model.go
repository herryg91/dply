package port_repository

import (
	"encoding/json"
	"time"

	"github.com/herryg91/dply/dply-server/entity"
)

type PortModel struct {
	Id        int        `gorm:"column:id"`
	Env       string     `gorm:"column:env"`
	Name      string     `gorm:"column:name"`
	Ports     []byte     `gorm:"column:ports"`
	CreatedBy int        `gorm:"column:created_by"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (pm *PortModel) ToPortEntity() *entity.Port {
	if pm == nil {
		return nil
	}

	ports := &entity.Port{
		Env:  pm.Env,
		Name: pm.Name,
	}
	json.Unmarshal(pm.Ports, &ports)
	return ports
}

func (PortModel) FromPortEntity(p entity.Port) *PortModel {
	portJson := map[string]interface{}{
		"access_type": string(p.AccessType),
		"external_ip": p.ExternalIP,
		"ports":       p.Ports,
	}

	portJsonMarshalled, _ := json.Marshal(&portJson)
	return &PortModel{
		Env:       p.Env,
		Name:      p.Name,
		Ports:     portJsonMarshalled,
		CreatedBy: p.CreatedBy,
	}
}

type PortTemplateModel struct {
	Id           int        `gorm:"column:id"`
	TemplateName string     `gorm:"column:template_name"`
	Ports        []byte     `gorm:"column:ports"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
}

func (pm *PortTemplateModel) ToPortTemplateEntity() *entity.PortTemplate {
	if pm == nil {
		return nil
	}

	resp := &entity.PortTemplate{
		TemplateName: pm.TemplateName,
	}
	json.Unmarshal(pm.Ports, &resp)
	return resp
}
func (PortTemplateModel) FromPortTemplateEntity(p entity.PortTemplate) *PortTemplateModel {
	portJson := map[string]interface{}{
		"access_type": string(p.AccessType),
		"external_ip": p.ExternalIP,
		"ports":       p.Ports,
	}

	portJsonMarshalled, _ := json.Marshal(&portJson)
	// ports, _ := json.Marshal(&p.Ports)
	return &PortTemplateModel{
		TemplateName: p.TemplateName,
		Ports:        portJsonMarshalled,
	}
}
