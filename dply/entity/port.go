package entity

type Port struct {
	Env        string     `json:"-"`
	Name       string     `json:"-"`
	AccessType AccessType `json:"access_type"`
	ExternalIP string     `json:"external_ip"`
	Ports      []PortSpec `json:"ports"`
}

type PortSpec struct {
	Name       string   `json:"name"`
	Port       int      `json:"port"`
	TargetPort int      `json:"target_port"`
	Protocol   PortType `json:"protocol"`
}

type PortType string

const (
	Port_TCP  PortType = "TCP"
	Port_UDP  PortType = "UDP"
	Port_SCTP PortType = "SCTP"
)

type AccessType string

const (
	Access_Type_ClusterIP    AccessType = "ClusterIP"
	Access_Type_LoadBalancer AccessType = "LoadBalancer"
)
