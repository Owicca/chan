CREATE DATABASE chan;

USE chan;

CREATE TABLE log_actions(
	id bigint NOT NULL AUTO_INCREMENT,
	action varchar(255) NOT NULL,-- insert, update, delete, virtual_delete
	subject bigint NOT NULL,-- subject that does the action
	object bigint NOT NULL,-- the subject acts on an object
	object_type varchar(255) NOT NULL,-- type of the object acted on
	data text NULL,-- data from insert or update
	PRIMARY KEY(id)
);

-- CREATE TABLE user_types(
-- 	user_type_id serial primary key,
-- 	deleted_at int NOT NULL DEFAULT 0,
-- 	name varchar(64) NOT NULL,
-- 	code varchar(64) NOT NULL
-- );

-- INSERT INTO user_types(user_type_id, name, code) VALUES
-- (1, 'R', 'root'),
-- (2, 'SA', 'site_admin'),
-- (3, 'BA', 'board_admin'),
-- (4, 'BOP', 'board_operator');

-- ---

-- CREATE TABLE user_statuses(
-- 	user_status_id serial primary key,
-- 	deleted_at int NOT NULL DEFAULT 0,
-- 	name varchar(64) NOT NULL,
-- 	code varchar(64) NOT NULL
-- );

-- INSERT INTO user_statuses(user_status_id, name, code) VALUES
-- (1, 'Disabled', 'D'),
-- (2, 'Active', 'A');

-- ---

-- CREATE TABLE post_statuses(
-- 	post_status_id serial primary key,
-- 	deleted_by int NOT NULL DEFAULT 0,
-- 	name varchar(64) NOT NULL,
-- 	code varchar(64) NOT NULL
-- );

-- INSERT INTO post_statuses(post_status_id, name, code) VALUES
-- (1, 'Disabled', 'D'),
-- (2, 'Hidden', 'H'),
-- (3, 'Active', 'A');

-- ---

-- CREATE TABLE actions(
-- 	action_id serial primary key,
-- 	deleted_at int NOT NULL DEFAULT 0,
-- 	name varchar(64) NOT NULL,
-- 	code varchar(64) NOT NULL
-- );

-- INSERT INTO actions(action_id, name, code) VALUES
-- (11, 'post_create', '1.1'),
-- (12, 'post_update', '1.2'),
-- (13, 'post_delete', '1.3'),

-- (21, 'thread_create', '2.1'),
-- (22, 'thread_update', '2.2'),
-- (23, 'thread_delete', '2.3'),

-- (31, 'board_create', '3.1'),
-- (32, 'board_update', '3.2'),
-- (33, 'board_delete', '3.3'),

-- (51, 'op_create', '5.1'),
-- (52, 'op_update', '5.2'),
-- (53, 'op_delete', '5.3'),

-- (61, 'board_create', '6.1'),
-- (62, 'board_update', '6.2'),
-- (63, 'board_delete', '6.3'),

-- (71, 'site_create', '7.1'),
-- (72, 'site_update', '7.2'),
-- (73, 'site_delete', '7.3');

-- ---

-- CREATE TABLE roles(
-- 	role_id serial primary key,
-- 	deleted_at int NOT NULL DEFAULT 0,
-- 	name varchar(64) NOT NULL
-- );

-- INSERT INTO roles(role_id, name) VALUES
-- (1, 'anon'),
-- (2, 'op'),
-- (3, 'board_admin'),
-- (4, 'site_admin'),
-- (5, 'root');

-- ---

-- CREATE TABLE action_to_role(
-- 	atg_id serial primary key,
-- 	action_id integer REFERENCES actions(action_id),
-- 	role_id integer REFERENCES roles(role_id)
-- );

-- INSERT INTO action_to_role(atg_id, action_id, role_id) VALUES
-- (1, 11, 1),
-- (2, 21, 1),

-- (3, 11, 2),
-- (4, 12, 2),
-- (5, 13, 2),
-- (6, 21, 2),
-- (7, 22, 2),
-- (8, 23, 2),

-- (9, 11, 3),
-- (10, 12, 3),
-- (11, 13, 3),
-- (12, 21, 3),
-- (13, 22, 3),
-- (14, 23, 3),
-- (15, 51, 3),
-- (16, 52, 3),
-- (17, 53, 3),

-- (18, 11, 4),
-- (19, 12, 4),
-- (20, 13, 4),
-- (21, 21, 4),
-- (22, 22, 4),
-- (23, 23, 4),
-- (24, 31, 4),
-- (25, 32, 4),
-- (26, 33, 4),
-- (27, 51, 4),
-- (28, 52, 4),
-- (29, 53, 4),
-- (30, 61, 4),
-- (31, 62, 4),
-- (32, 63, 4),

-- (33, 11, 5),
-- (34, 12, 5),
-- (35, 13, 5),
-- (36, 21, 5),
-- (37, 22, 5),
-- (38, 23, 5),
-- (39, 31, 5),
-- (40, 32, 5),
-- (41, 33, 5),
-- (42, 51, 5),
-- (43, 52, 5),
-- (44, 53, 5),
-- (45, 61, 5),
-- (46, 62, 5),
-- (47, 63, 5),
-- (48, 71, 5),
-- (49, 72, 5),
-- (50, 73, 5);

-- ---

-- CREATE TABLE medias(
-- 	media_id serial primary key,
-- 	deleted_at int NOT NULL DEFAULT 0,
-- 	name varchar(64) NOT NULL,-- seo name
-- 	code varchar(64) NOT NULL,-- img, vid
-- 	path text NOT NULL
-- );

-- ---

-- CREATE TABLE users(
-- 	user_id serial primary key,
-- 	deleted_at int NOT NULL DEFAULT 0,
-- 	username varchar(255) NOT NULL,
-- 	email varchar(255) NOT NULL,
-- 	password varchar(255) NOT NULL,
-- 	salt varchar(255) NOT NULL,
-- 	user_type_id integer REFERENCES user_types(user_type_id),
-- 	user_status_id integer REFERENCES user_statuses(user_status_id),
-- 	role_id integer REFERENCES roles(role_id)
-- );

-- INSERT INTO users(user_id, username, email, password, salt, user_type_id, user_status_id, role_id) VALUES
-- (1, 'root', 'root@root.com', '$2a$10$KI4EmNCFlvYteYeI//1s2OhR5jNmIJlEdrgLOzYINyuf8MrUNbaAC', 'salt', 1, 2, 5);-- password: password; salt: salt

-- ---

-- CREATE TABLE boards(
-- 	board_id serial primary key,
-- 	deleted_at int NOT NULL DEFAULT 0,
-- 	name varchar(255) NOT NULL,
-- 	code varchar(64) NOT NULL,
-- 	description text NOT NULL,
-- 	media_id integer REFERENCES medias(media_id)
-- );

-- ---

-- CREATE TABLE threads(
-- 	thread_id serial primary key,
-- 	deleted_at int NOT NULL DEFAULT 0,
-- 	board_id integer REFERENCES boards(board_id)
-- );

-- ---

-- CREATE TABLE posts(
-- 	post_id serial primary key,
-- 	deleted_at int NOT NULL DEFAULT 0,
-- 	thread_id integer REFERENCES threads(thread_id),
-- 	status integer REFERENCES post_statuses(post_status_id),
-- 	media integer REFERENCES medias(media_id),
-- 	name varchar(255) NOT NULL,
-- 	content text NOT NULL
-- );

-- ---

-- CREATE TABLE post_to_post(
-- 	ptp_id serial primary key,
-- 	post_from_id integer REFERENCES posts(post_id),
-- 	post_to_id integer REFERENCES posts(post_id)
-- )