package repository

type DeployRepository interface {
	Deploy(project, env, name, digest string) error
	Redeploy(project, env, name string) error
}
