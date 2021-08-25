### entities:
- #image: jpg/jpeg/png/webp/gif

- #action:
	- a subject does an action on an object
	- structure:
		- name: short text
		- code: short text

- actions:
	1. post:
		1. create
		2. update
		3. delete
	2. thread:
		1. create
		2. update
		3. delete
	3. board:
		1. create
		2. update
		3. delete
	4. site
	5. op:
		1. create
		2. update
		3. delete
	6. board admin:
		1. create
		2. update
		3. delete
	7. site admin:
		1. create
		2. update
		3. delete

- #group: a list of actions

- default_groups(new groups can be created, but these are unchangeable):
	- anon:
		- action_list:
			- 1.1
			- 2.1
	- op:
		- +anon
		- action_list:
			- 1.2, 1.3
			- 2.2, 2.3
	- board admin:
		- +op
		- action_list:
			- 5.1, 5.2, 5.3
	- site admin:
		- +board admin
		- action_list:
			- 6.1, 6.2, 6.3
			- 3.1, 3.2, 3.3
	- root:
		- +site admin
		- action_list:
			- 7.1, 7.2, 7.3

- #user:
	- username: short text
	- email: email
	- password: short text
	- type: char
	- status: char
	- group: int

- #post:
	- info:
		- creating a post generates an id stored in a cookie,
		as long as you have the cookie, you can delete your post
	- metadata:
		- thread: int
		- status: char
		- name: short text
		- media:
			- image:
				- jpg
				- jpeg
				- png
				- webp
				- gif
			- video:
				- webm:
					- 120s
					- 2048x2048
					- 3MB
				- mp4
	- message:
		- type: text
		- plain text
		- can contain ">"(quote), ">>"(link), ">>>/x/123121"(cross-link) are allowed
		- can contain tags "[tag]content[/tag]"
- #thread:
	- a thread needs at least one post with an image
	- page limit: 10
	- post per page limit: 50
	- posts: []#post
- #board
	- info:
		- thread limit: 10
	- name: short text
	- code: short text
	- descriptions: text
	- threads: []#thread
	- image: #image


### db_schema:
- every entity contains:
	- deleted_at int DEFAULT 0

- log_actions:
	- log_id:
		- serial(this should be the highest int size available)
		- primary key
	- action: varchar(255) not null(insert, update, delete)
	- subject: int not null
	- object: int not null
	- object_type: varchar(255) not null
	- data: large text null(has values only on insert and update)

- user_type:
	- user_type_id: serial primary key
	- name: varchar(64)
	- code: varchar(64)
	- values on db creation:
		- R: root
		- SA: site_admin
		- BA: board_admin
		- BOP: board_operator

- user_status:
	- user_status_id: serial primary key
	- name: varchar(64)
	- code: varchar(64) DEFAULT 'D'
	- values on db creation:
		- D: disabled
		- A: active

- post_status:
	- post_status_id: serial primary key
	- name: varchar(64)
	- code: varchar(64) DEFAULT 'D'
	- values on db creation:
		- D: disabled
		- H: hidden
		- A: active

- action:
	- action_id: serial primary key
	- name: varchar(64)
	- code: varchar(64)

- role:
	- role_id: serial primary key
	- name: varchar(64)

- media_type:
	- media_type_id: serial primary key
	- name: varchar(64)
	- code: varchar(64)

- media:
	- media_id: serial primary key
	- media_type_id: fk to media_type.media_type_id
	- path: text

- action_to_role:
	- atg_id: serial primary key
	- action_id: fk to action.action_id
	- role_id: fk to role.role_id

- user:
	- user_id: serial primary key
	- username: varchar(255)
	- email: varchar(255)
	- password: varchar(255)
	- salt: varchar(255)
	- user_type_id: fk to user_type.user_type_id
	- user_status_id: fk to user_status.user_status_id
	- role_id: fk to role.role_id

- board:
	- board_id: serial primary key
	- name: varchar(255)
	- code: varchar(64)
	- description: text
	- media_id: fk to media.media_id

- thread:
	- thread_id: serial primary key
	- board_id: fk to board.board_id

- post:
	- post_id: serial primary key
	- thread_id: fk to thread.thread_id
	- name: varchar(255)
	- status: fk to post_status.post_status_id
	- media: fk to media.media_id
	- content: text

- post_to_post:
	- ptp_id: serial primary key
	- post_from_id: fk to post.post_id
	- post_to_id: fk to post.post_id


### sitemap:
- frontend:
	- /: site index
	- /<board_name:string>/: board index
	- /<board_name:string>/thread/<thread_id:int>/: thread index
	- /media/<media_id:int>.<extension:string>/: media view
- backend:
	- /admin?dispatch=auth.login: login form
	- /admin?dispatch=auth.reset: reset password form
	- /admin?dispatch=auth.logout: logout endpoint

	- /admin?dispatch=user.manage: user list
	- /admin?dispatch=user.create: user creation
	- /admin?dispatch=user.update: user updating
	- /admin?dispatch=user.delete: user deletion endpoint

	- /admin?dispatch=index.index: dashboard

	- /admin?dispatch=boards.manage: board list
	- /admin?dispatch=boards.create: board creation
	- /admin?dispatch=boards.update: board updating

	- /admin?dispatch=threads.manage: thread list
	- /admin?dispatch=threads.update: thread updating

	- /admin?dispatch=posts.manage: post list


### tech:
- db: postgresql
- backend: golang
- frontend: golang templates + react


### MVP:
- working imageboard with no users except for anons:
	- [ ] serve sitemap pages
	- [x] serve media
	- [ ] post:
		- [ ] name
		- [ ] media
		- [ ] content
	- [ ] create thread:
		- [ ] enforce thread rules