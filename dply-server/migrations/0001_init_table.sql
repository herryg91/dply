DROP TABLE IF EXISTS deployment;
DROP TABLE IF EXISTS image;
DROP TABLE IF EXISTS affinity_template;
DROP TABLE IF EXISTS affinity;
DROP TABLE IF EXISTS port_template;
DROP TABLE IF EXISTS port;
DROP TABLE IF EXISTS scale;
DROP TABLE IF EXISTS envar;
DROP TABLE IF EXISTS user;

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

CREATE TABLE envar (
    id INT NOT NULL AUTO_INCREMENT,
    env VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    variables JSON NOT NULL,
    created_by INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( id ),
    UNIQUE(env, name)
);

CREATE TABLE scale (
    id INT NOT NULL AUTO_INCREMENT,
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
    UNIQUE(env, name)
);

CREATE TABLE port (
    id INT NOT NULL AUTO_INCREMENT,
    env VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    ports JSON NOT NULL,
    created_by INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( id ),
    UNIQUE(env, name)
);

CREATE TABLE port_template (
    id INT NOT NULL AUTO_INCREMENT,
    template_name VARCHAR(255) NOT NULL UNIQUE,
    ports JSON NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( id )
);
INSERT INTO port_template (template_name, ports) values ('default', '[{"name":"http","port":80,"protocol":"TCP"}]');

CREATE TABLE affinity (
    id INT NOT NULL AUTO_INCREMENT,
    env VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    affinity JSON NOT NULL,
    created_by INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( id ),
    UNIQUE(env, name)
);

CREATE TABLE affinity_template (
    id INT NOT NULL AUTO_INCREMENT,
    template_name VARCHAR(255) NOT NULL UNIQUE,
    affinity JSON NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( id )
);
INSERT INTO affinity_template (template_name, affinity) values ('default', '{"node_affinity":[],"pod_affinity":[],"pod_anti_affinity":[]}');

CREATE TABLE image (
    id INT NOT NULL AUTO_INCREMENT,
    digest VARCHAR(32) NOT NULL UNIQUE,
    image VARCHAR(255) NOT NULL UNIQUE,
    repository VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    created_by INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( id )
);
CREATE INDEX image_repository_created_at ON image (repository, created_at);

CREATE TABLE deployment (
    id INT NOT NULL AUTO_INCREMENT,
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
CREATE INDEX deployment_env_name_created_at ON deployment (env, name, created_at);
