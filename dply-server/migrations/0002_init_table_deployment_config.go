package migration_file

import "gorm.io/gorm"

type migration_0002 struct {
	db *gorm.DB
}

func NewMigration0002(db *gorm.DB) MigrationFile {
	return &migration_0002{db}
}

func (m *migration_0002) GetName() string {
	return "0002_init_table_deployment_config"
}

func (m *migration_0002) Down() error {
	err := m.db.Exec(`DROP TABLE IF EXISTS deployment_config;`).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *migration_0002) Up() error {
	err := m.db.Exec(`
	CREATE TABLE deployment_config (
		id INT NOT NULL AUTO_INCREMENT,
		project VARCHAR(128) NOT NULL,
		env VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		config JSON NOT NULL,
		created_by INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY ( id ),
		UNIQUE(project, env, name)
	);
	`).Error
	if err != nil {
		return err
	}
	return nil
}
