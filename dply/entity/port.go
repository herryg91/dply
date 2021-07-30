package entity

type Port struct {
	Env   string     `json:"-"`
	Name  string     `json:"-"`
	Ports []PortSpec `json:"ports"`
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
