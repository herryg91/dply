package entity

type Envar struct {
	Project   string                 `json:"-"`
	Env       string                 `json:"-"`
	Name      string                 `json:"-"`
	Variables map[string]interface{} `json:"variables"`
}
