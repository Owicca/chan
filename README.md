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

- #role: a list of actions

- default_roles(new roles can be created, but these are unchangeable):
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
- db: mysql
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