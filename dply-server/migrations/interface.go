package migration_file

type MigrationFile interface {
	GetName() string
	Up() error
	Down() error
}
