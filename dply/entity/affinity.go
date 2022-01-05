package entity

type Affinity struct {
	Project         string         `json:"-"`
	Env             string         `json:"-"`
	Name            string         `json:"-"`
	NodeAffinity    []AffinityTerm `json:"node_affinity"`
	PodAffinity     []AffinityTerm `json:"pod_affinity"`
	PodAntiAffinity []AffinityTerm `json:"pod_anti_affinity"`
}

type AffinityTerm struct {
	Mode        AffinityMode     `json:"mode"`
	Key         string           `json:"key"`
	Operator    AffinityOperator `json:"operator"`
	Values      []string         `json:"values"`
	Weight      int              `json:"weight"`
	TopologyKey string           `json:"topology_key"`
}

type AffinityOperator string

const (
	AffinityOperator_In           AffinityOperator = "In"
	AffinityOperator_NotIn        AffinityOperator = "NotIn"
	AffinityOperator_Exists       AffinityOperator = "Exists"
	AffinityOperator_DoesNotExist AffinityOperator = "DoesNotExist"
	AffinityOperator_Gt           AffinityOperator = "Gt"
	AffinityOperator_Lt           AffinityOperator = "Lt"
)

type AffinityMode string

const (
	AffinityMode_Required  = "required"
	AffinityMode_Preferred = "preferred"
)
