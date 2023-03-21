package entity

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

type Affinity struct {
	Project         string               `json:"project"`
	Env             string               `json:"env"`
	Name            string               `json:"name"`
	NodeAffinity    []AffinityTerm       `json:"node_affinity"`
	PodAffinity     []AffinityTerm       `json:"pod_affinity"`
	PodAntiAffinity []AffinityTerm       `json:"pod_anti_affinity"`
	Tolerations     []AffinityToleration `json:"tolerations"`
	CreatedBy       int                  `json:"created_by"`
}

type AffinityTerm struct {
	Mode        AffinityMode     `json:"mode"`
	Key         string           `json:"key"`
	Operator    AffinityOperator `json:"operator"`
	Values      []string         `json:"values"`
	Weight      int              `json:"weight"`
	TopologyKey string           `json:"topology_key"`
}

type AffinityToleration struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
	Effect   string `json:"effect"`
}

func (Affinity) DefaultAffinity(env, name string) *Affinity {
	return &Affinity{
		Env:             env,
		Name:            name,
		NodeAffinity:    []AffinityTerm{},
		PodAffinity:     []AffinityTerm{},
		PodAntiAffinity: []AffinityTerm{},
		Tolerations:     []AffinityToleration{},
	}
}

type AffinityTemplate struct {
	TemplateName    string               `json:"template_name"`
	NodeAffinity    []AffinityTerm       `json:"node_affinity"`
	PodAffinity     []AffinityTerm       `json:"pod_affinity"`
	PodAntiAffinity []AffinityTerm       `json:"pod_anti_affinity"`
	Tolerations     []AffinityToleration `json:"tolerations"`
}

func (AffinityTemplate) DefaultAffinityTemplate() *AffinityTemplate {
	return &AffinityTemplate{
		TemplateName:    "default",
		NodeAffinity:    []AffinityTerm{},
		PodAffinity:     []AffinityTerm{},
		PodAntiAffinity: []AffinityTerm{},
		Tolerations:     []AffinityToleration{},
	}
}
func (at *AffinityTemplate) ToAffinityEntity(env, name string) *Affinity {
	return &Affinity{
		Env:             env,
		Name:            name,
		NodeAffinity:    at.NodeAffinity,
		PodAffinity:     at.PodAffinity,
		PodAntiAffinity: at.PodAntiAffinity,
		Tolerations:     at.Tolerations,
	}
}
