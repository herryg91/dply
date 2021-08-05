package entity

type Port struct {
	Env        string     `json:"env"`
	Name       string     `json:"name"`
	AccessType AccessType `json:"access_type"`
	ExternalIP string     `json:"external_ip"`
	Ports      []PortSpec `json:"ports"`
	CreatedBy  int        `json:"created_by"`
}

type PortSpec struct {
	Name       string           `json:"name"`
	Port       int              `json:"port"`
	RemotePort int              `json:"remote_port"`
	Protocol   PortProtocolType `json:"protocol"`
}

type AccessType string

const (
	Access_Type_ClusterIP    AccessType = "ClusterIP"
	Access_Type_LoadBalancer AccessType = "LoadBalancer"
)

type PortProtocolType string

const (
	Port_TCP  PortProtocolType = "TCP"
	Port_UDP  PortProtocolType = "UDP"
	Port_SCTP PortProtocolType = "SCTP"
)

func (Port) DefaultPort(env, name string) *Port {
	return &Port{
		Env:        env,
		Name:       name,
		AccessType: Access_Type_ClusterIP,
		ExternalIP: "",
		Ports: []PortSpec{
			{Name: "http", Port: 80, RemotePort: 80, Protocol: Port_TCP},
		},
	}
}

type PortTemplate struct {
	TemplateName string     `json:"-"`
	AccessType   AccessType `json:"access_type"`
	ExternalIP   string     `json:"external_ip"`
	Ports        []PortSpec `json:"ports"`
}

func (PortTemplate) DefaultPortTemplate() *PortTemplate {
	return &PortTemplate{
		TemplateName: "default",
		AccessType:   Access_Type_ClusterIP,
		ExternalIP:   "",
		Ports: []PortSpec{
			{Name: "http", Port: 80, RemotePort: 80, Protocol: Port_TCP},
		},
	}
}
func (pt *PortTemplate) ToPortEntity(env, name string) *Port {
	return &Port{
		Env:        env,
		Name:       name,
		AccessType: pt.AccessType,
		ExternalIP: pt.ExternalIP,
		Ports:      pt.Ports,
	}
}
