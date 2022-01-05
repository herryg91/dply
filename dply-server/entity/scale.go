package entity

type Scale struct {
	Project              string `json:"project"`
	Env                  string `json:"env"`
	Name                 string `json:"name"`
	MinReplica           int    `json_name:"min_replica"`
	MaxReplica           int    `json_name:"max_replica"`
	MinCpu               int    `json_name:"min_cpu"`
	MaxCpu               int    `json_name:"max_cpu"`
	MinMemory            int    `json_name:"min_memory"`
	MaxMemory            int    `json_name:"max_memory"`
	TargetCPUUtilization int    `json_name:"target_cpu"`
	CreatedBy            int    `json:"created_by"`
}

func (Scale) DefaultScale(project, env, name string) *Scale {
	s := &Scale{}
	s.Project = project
	s.Env = env
	s.Name = name
	s.MinReplica = 1
	s.MaxReplica = 3
	s.MinCpu = 64
	s.MaxCpu = 64
	s.MinMemory = 64
	s.MaxMemory = 64
	s.TargetCPUUtilization = 70
	return s
}
