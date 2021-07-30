package entity

type Envar struct {
	Env       string                 `json:"env"`
	Name      string                 `json:"name"`
	Variables map[string]interface{} `json:"variables"`
	CreatedBy int                    `json:"created_by"`
}

func (Envar) DefaultEnvar(env, name string) *Envar {
	return &Envar{
		Env:       env,
		Name:      name,
		Variables: map[string]interface{}{},
		CreatedBy: 0,
	}
}
