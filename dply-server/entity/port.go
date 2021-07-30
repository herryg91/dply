package entity

type Port struct {
	Env       string     `json:"env"`
	Name      string     `json:"name"`
	Ports     []PortSpec `json:"ports"`
	CreatedBy int        `json:"created_by"`
}

type PortSpec struct {
	Name     string   `json:"name"`
	Port     int      `json:"port"`
	Protocol PortType `json:"protocol"`
}

type PortType string

const (
	Port_TCP  PortType = "TCP"
	Port_UDP  PortType = "UDP"
	Port_SCTP PortType = "SCTP"
)

func (Port) DefaultPort(env, name string) *Port {
	return &Port{
		Env:  env,
		Name: name,
		Ports: []PortSpec{
			{Name: "http", Port: 80, Protocol: Port_TCP},
		},
	}
}

type PortTemplate struct {
	TemplateName string     `json:"template_name"`
	Ports        []PortSpec `json:"ports"`
}

func (PortTemplate) DefaultPortTemplate() *PortTemplate {
	return &PortTemplate{
		TemplateName: "default",
		Ports: []PortSpec{
			{Name: "http", Port: 80, Protocol: Port_TCP},
		},
	}
}
func (pt *PortTemplate) ToPortEntity(env, name string) *Port {
	return &Port{
		Env:   env,
		Name:  name,
		Ports: pt.Ports,
	}
}
