DROP TABLE IF EXISTS deployment_config;

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
