package entity

type Envar struct {
	Env       string                 `json:"-"`
	Name      string                 `json:"-"`
	Variables map[string]interface{} `json:"variables"`
}
