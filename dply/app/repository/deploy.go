package repository

type DeployRepository interface {
	Deploy(env, name, digest string) error
	Redeploy(env, name string) error
}
