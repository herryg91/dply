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
	json.Unmarshal(pm.Ports, &ports.Ports)
	return ports
}

func (PortModel) FromPortEntity(p entity.Port) *PortModel {
	ports, _ := json.Marshal(&p.Ports)
	return &PortModel{
		Env:       p.Env,
		Name:      p.Name,
		Ports:     ports,
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

	portSpecs := []entity.PortSpec{}
	json.Unmarshal(pm.Ports, &portSpecs)
	return &entity.PortTemplate{
		TemplateName: pm.TemplateName,
		Ports:        portSpecs,
	}
}
func (PortTemplateModel) FromPortTemplateEntity(p entity.PortTemplate) *PortTemplateModel {
	ports, _ := json.Marshal(&p.Ports)
	return &PortTemplateModel{
		TemplateName: p.TemplateName,
		Ports:        ports,
	}
}
