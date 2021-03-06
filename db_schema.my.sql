CREATE TABLE log_actions(
	id bigint NOT NULL AUTO_INCREMENT,
	action varchar(255) NOT NULL,-- insert, update, delete, virtual_delete
	subject bigint NOT NULL,-- subject that does the action
	object bigint NOT NULL,-- the subject acts on an object
	object_type varchar(255) NOT NULL,-- type of the object acted on
	data text NULL,-- data from insert or update
	PRIMARY KEY(id)
);

CREATE TABLE actions(
	id bigint NOT NULL AUTO_INCREMENT,
	deleted_at int NOT NULL DEFAULT 0,
	name varchar(64) NOT NULL,
	PRIMARY KEY(id)
);

INSERT INTO actions(id, name) VALUES
(1, 'read'),
(2, 'create'),
(3, 'update'),
(4, 'delete');

CREATE TABLE objects(
	id bigint NOT NULL AUTO_INCREMENT,
	name varchar(64) NOT NULL,
	PRIMARY KEY(id)
);

INSERT INTO objects(id, name) VALUES
(1, "post"),
(2, "thread"),
(3, "board"),
(4, "op"),
(5, "site");

CREATE TABLE action_to_object (
	id bigint NOT NULL AUTO_INCREMENT,
	obj_id bigint,
	action_id bigint,
	PRIMARY KEY (id),
	CONSTRAINT fk_action_to_object_object FOREIGN KEY (obj_id) REFERENCES objects(id),
	CONSTRAINT fk_action_to_object_action FOREIGN KEY (action_id) REFERENCES actions(id)
);

INSERT INTO action_to_object (id, action_id, obj_id) VALUES
(11, 1,	1),
(12, 2,	1),
(13, 3,	1),
(14, 4,	1),
(21, 1,	2),
(22, 2,	2),
(23, 3,	2),
(24, 4,	2),
(31, 1,	3),
(32, 2,	3),
(33, 3,	3),
(34, 4,	3),
(41, 1,	4),
(42, 2,	4),
(43, 3,	4),
(44, 4,	4),
(51, 1,	5),
(52, 2,	5),
(53, 3,	5),
(54, 4,	5);

CREATE TABLE roles(
	id bigint NOT NULL AUTO_INCREMENT,
	deleted_at int NOT NULL DEFAULT 0,
	name varchar(64) NOT NULL,
	PRIMARY KEY(id)
);

INSERT INTO roles(id, name) VALUES
(1, 'root'),
(2, 'site_admin'),
(3, 'board_admin'),
(4, 'op');

CREATE TABLE pair_to_role(
	id bigint NOT NULL AUTO_INCREMENT,
	ato_id bigint NOT NULL,
	role_id bigint NOT NULL,
	PRIMARY KEY(id),
	FOREIGN KEY(ato_id) REFERENCES action_to_object(id),
	FOREIGN KEY(role_id) REFERENCES roles(id)
);

INSERT INTO pair_to_role(ato_id, role_id) VALUES
(11, 1),
(12, 1),
(13, 1),
(14, 1),
(21, 1),
(22, 1),
(23, 1),
(24, 1),
(31, 1),
(32, 1),
(33, 1),
(34, 1),
(41, 1),
(42, 1),
(43, 1),
(44, 1),
(51, 1),
(52, 1),
(53, 1),
(54, 1),

(11, 2),
(12, 2),
(13, 2),
(21, 2),
(22, 2),
(23, 2),
(31, 2),
(32, 2),
(33, 2),
(41, 2),
(42, 2),
(43, 2),
(44, 2),
(51, 2),
(52, 2),
(53, 2),

(11, 3),
(12, 3),
(13, 3),
(21, 3),
(22, 3),
(23, 3),
(31, 3),
(32, 3),
(33, 3),
(41, 3),
(42, 3),
(43, 3),
(44, 3),

(11, 4),
(12, 4),
(13, 4),
(21, 4),
(22, 4),
(23, 4);

CREATE TABLE users(
	id bigint NOT NULL AUTO_INCREMENT,
	deleted_at int NOT NULL DEFAULT 0,
	username varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	password varchar(255) NOT NULL,-- bcrypt(pepper + pass)
	pepper varchar(255) NOT NULL,
	status varchar(2) NOT NULL,
	role_id bigint NOT NULL,
	FOREIGN KEY(role_id) REFERENCES roles(id),
	PRIMARY KEY(id)
);

INSERT INTO users(id, username, email, password, pepper, status, role_id) VALUES
(1, 'root', 'root@root.com', '$2a$10$ILRgDxBKyNyBdGP9969PuO7Egb5naBWgQeAJwjrHvw39mzaFwtDU2', 'pepper', "A", 1);-- password: password; pepper: pepper

-- ---

CREATE TABLE media (
	object_id bigint NOT NULL,
	object_type varchar(64) NOT NULL,
	deleted_at int NOT NULL DEFAULT 0,
	name varchar(64) NOT NULL,-- seo name
	type varchar(64) NOT NULL DEFAULT 'img',-- img, vid
	path text NOT NULL,
	thumb TEXT NOT NULL,
	x SMALLINT NOT NULL,
	y SMALLINT NOT NULL,
	size BIGINT NOT NULL
);

-- ---

CREATE TABLE topics(
	id bigint NOT NULL AUTO_INCREMENT,
	deleted_at int NOT NULL DEFAULT 0,
	name varchar(255) NOT NULL,
	PRIMARY KEY(id)
);

-- ---

CREATE TABLE boards(
	id bigint NOT NULL AUTO_INCREMENT,
	deleted_at int NOT NULL DEFAULT 0,
	topic_id bigint NOT NULL,
	name varchar(255) NOT NULL,
	code varchar(64) NOT NULL,
	description text NOT NULL,
	FOREIGN KEY(topic_id) REFERENCES topics(id),
	PRIMARY KEY(id)
);

-- ---

CREATE TABLE threads (
	id bigint NOT NULL AUTO_INCREMENT,
	deleted_at int NOT NULL DEFAULT 0,
	board_id bigint NOT NULL,
	subject varchar(255) NOT NULL,
	primary_post_id bigint NOT NULL,
	FOREIGN KEY(board_id) REFERENCES boards(id),
	PRIMARY KEY(id)
);

-- ---

CREATE TABLE posts (
	id bigint NOT NULL AUTO_INCREMENT,
	created_at bigint NOT NULL,
	deleted_at int NOT NULL DEFAULT '0',
	tripcode varchar(255) DEFAULT NULL,
	secure_tripcode varchar(255) DEFAULT NULL,
	status varchar(10) NOT NULL DEFAULT 'A',
	thread_id bigint NOT NULL,
	name varchar(255) DEFAULT NULL,
	content text NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (thread_id) REFERENCES threads (id)
);

-- ---

CREATE TABLE links (
	src bigint NOT NULL,
	dest bigint NOT NULL,
	PRIMARY KEY (src, dest),
	FOREIGN KEY (src) REFERENCES posts(id),
	FOREIGN KEY (dest) REFERENCES posts(id)
);

-- ---

CREATE TABLE sessions (
	id bigint NOT NULL,
	data text NULL,
	PRIMARY KEY(id)
);
