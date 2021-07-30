package entity

type Scale struct {
	Env                  string `json:"-"`
	Name                 string `json:"-"`
	MinReplica           int    `json:"min_replica"`
	MaxReplica           int    `json:"max_replica"`
	MinCpu               int    `json:"min_cpu"`
	MaxCpu               int    `json:"max_cpu"`
	MinMemory            int    `json:"min_memory"`
	MaxMemory            int    `json:"max_memory"`
	TargetCPUUtilization int    `json:"target_cpu"`
}
