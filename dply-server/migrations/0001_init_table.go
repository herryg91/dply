package migration_file

import "gorm.io/gorm"

type migration_0001 struct {
	db *gorm.DB
}

func NewMigration0001(db *gorm.DB) MigrationFile {
	return &migration_0001{db}
}

func (m *migration_0001) GetName() string {
	return "0001_init_table"
}

func (m *migration_0001) Down() error {
	err := m.db.Exec(`DROP TABLE IF EXISTS deployment;`).Error
	if err != nil {
		return err
	}
	err = m.db.Exec(`DROP TABLE IF EXISTS image;`).Error
	if err != nil {
		return err
	}
	err = m.db.Exec(`DROP TABLE IF EXISTS affinity_template;`).Error
	if err != nil {
		return err
	}
	err = m.db.Exec(`DROP TABLE IF EXISTS affinity;`).Error
	if err != nil {
		return err
	}
	err = m.db.Exec(`DROP TABLE IF EXISTS port_template;`).Error
	if err != nil {
		return err
	}
	err = m.db.Exec(`DROP TABLE IF EXISTS port;`).Error
	if err != nil {
		return err
	}
	err = m.db.Exec(`DROP TABLE IF EXISTS scale;`).Error
	if err != nil {
		return err
	}
	err = m.db.Exec(`DROP TABLE IF EXISTS envar;`).Error
	if err != nil {
		return err
	}
	err = m.db.Exec(`DROP TABLE IF EXISTS user;`).Error
	if err != nil {
		return err
	}
	err = m.db.Exec(`DROP TABLE IF EXISTS project;`).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *migration_0001) Up() error {
	err := m.db.Exec(`
	CREATE TABLE project (
		id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(128) NOT NULL UNIQUE DEFAULT 'default',
		description TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY ( id )
	);
	`).Error
	if err != nil {
		return err
	}

	err = m.db.Exec(`
	INSERT INTO project (name, description) values ('default', '');
	`).Error
	if err != nil {
		return err
	}

	err = m.db.Exec(`
	CREATE TABLE user (
		id INT NOT NULL AUTO_INCREMENT,
		email VARCHAR(128) NOT NULL UNIQUE,
		password VARCHAR(128) NOT NULL,
		usertype ENUM('admin', 'user') DEFAULT 'user' NOT NULL,
		name VARCHAR(128) NOT NULL,
		token VARCHAR(128) NOT NULL UNIQUE,
		status boolean NOT NULL DEFAULT TRUE,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY ( id )
	);
	`).Error
	if err != nil {
		return err
	}

	err = m.db.Exec(`
	CREATE TABLE envar (
		id INT NOT NULL AUTO_INCREMENT,
		project VARCHAR(128) NOT NULL DEFAULT 'default',
		env VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		variables JSON NOT NULL,
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

	err = m.db.Exec(`
	CREATE TABLE scale (
		id INT NOT NULL AUTO_INCREMENT,
		project VARCHAR(128) NOT NULL DEFAULT 'default',
		env VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
	
		min_replica INT NOT NULL DEFAULT 1,
		max_replica INT NOT NULL DEFAULT 3,
		min_cpu INT NOT NULL DEFAULT 100,
		max_cpu INT NOT NULL DEFAULT 250,
		min_memory INT NOT NULL DEFAULT 100,
		max_memory INT NOT NULL DEFAULT 250,
		target_cpu INT NOT NULL DEFAULT 70,
	
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

	err = m.db.Exec(`
	CREATE TABLE port (
		id INT NOT NULL AUTO_INCREMENT,
		project VARCHAR(128) NOT NULL DEFAULT 'default',
		env VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		ports JSON NOT NULL,
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

	err = m.db.Exec(`
	CREATE TABLE port_template (
		id INT NOT NULL AUTO_INCREMENT,
		template_name VARCHAR(255) NOT NULL UNIQUE,
		ports JSON NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY ( id )
	);
	`).Error
	if err != nil {
		return err
	}

	err = m.db.Exec(`
	INSERT INTO port_template (template_name, ports) values ('default', '{"access_type":"ClusterIP","external_ip":"","ports":[{"name":"http","port":80,"remote_port":80,"protocol":"TCP"}]}');
	`).Error
	if err != nil {
		return err
	}

	err = m.db.Exec(`
	CREATE TABLE affinity (
		id INT NOT NULL AUTO_INCREMENT,
		project VARCHAR(128) NOT NULL DEFAULT 'default',
		env VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		affinity JSON NOT NULL,
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

	err = m.db.Exec(`
	CREATE TABLE affinity_template (
		id INT NOT NULL AUTO_INCREMENT,
		template_name VARCHAR(255) NOT NULL UNIQUE,
		affinity JSON NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY ( id )
	);
	`).Error
	if err != nil {
		return err
	}

	err = m.db.Exec(`
	INSERT INTO affinity_template (template_name, affinity) values ('default', '{"node_affinity":[],"pod_affinity":[],"pod_anti_affinity":[]}');
	`).Error
	if err != nil {
		return err
	}

	err = m.db.Exec(`
	CREATE TABLE image (
		id INT NOT NULL AUTO_INCREMENT,
		digest VARCHAR(32) NOT NULL UNIQUE,
		image VARCHAR(255) NOT NULL UNIQUE,
		project VARCHAR(128) NOT NULL DEFAULT 'default',
		repository VARCHAR(128) NOT NULL,
		description TEXT NOT NULL,
		created_by INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY ( id )
	);
	`).Error
	if err != nil {
		return err
	}

	err = m.db.Exec(`CREATE INDEX image_project_repository_created_at ON image (project, repository, created_at);`).Error
	if err != nil {
		return err
	}

	err = m.db.Exec(`
	CREATE TABLE deployment (
		id INT NOT NULL AUTO_INCREMENT,
		project VARCHAR(128) NOT NULL DEFAULT 'default',
		env VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		image_digest VARCHAR(32) NOT NULL,
		variables JSON NOT NULL,
		ports JSON NOT NULL,
		created_by VARCHAR(128) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY ( id )
	);
	`).Error
	if err != nil {
		return err
	}

	err = m.db.Exec(`CREATE INDEX deployment_project_env_name_created_at ON deployment (project, env, name, created_at);`).Error
	if err != nil {
		return err
	}
	return nil
}
