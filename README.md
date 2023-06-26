```
classDiagram
    Topics "1" <-- "0..n" Boards
    Boards "1" <-- "0..n" Threads
    Threads "1" <-- "0..n" Posts
    Boards "1" <-- "0..n" Media
    Posts "1" <-- "0..n" Media
    class Topics{
        +int ID
        +string Name
        +int Deleted_at
    }
    class Boards{
        +int ID
        +int Deleted_at
        +string Name
        +string Code
        +string Description
        +int Topic_id
    }
    class Threads{
        +int ID
        +int Deleted_at
        +int Board_id
        +int Primary_post_id
        +subject Subject
        +string Content
        +[]posts.Post Preview
    }
    class Posts{
        +int ID
        +int Created_at
        +int Deleted_at
        +string Tripcode
        +string SecureTripcode
        +string Status
        +int Thread_id
        +string Name
        +string Content
        +media.Media Media
        +[]Link LinkList
    }
    class Media {
        +int Object_id
        +string Object_type
        +int Deleted_at
        +string Name
        +string Type
        +string Path
        +string Thumb
        +int X
        +int Y
        +int Size
    }
```

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



### tech:
- db: mysql
- backend: golang
- frontend: golang templates


### Setup:
1. `go get github.com/githubnemo/CompileDaemon`
2. `CompileDaemon.exe -build="go build -o ..\main.exe .\main.go" -command="..\main.exe" -exclude-dir=".git" -exclude-dir="log" -exclude="(.*\.exe)$" -pattern="(.*)$" -verbose`


### media potential requirements:
- type:
	- image:
		- gif:
			- thumb: extract frame and resize
		- png:
			- thumb: copy and resize
		- jpeg:
			- thumb: copy and resize
		- jpg:
			- thumb: copy and resize
		- webp:
			- thumb: copy and resize
	- video:
		- webm:
			- thumb: extract frame and resize
		- mp4:
			- thumb: extract frame and resize
- size:
	- min size
	- max size
	- min resolution
	- max resolution


closed threads?

### reporting:
- board: # of threads in a daterange
- topic: # of threads in a daterange


### todo:
- [] abstract methods in `/models/`
- [] test at least 50% of `/models/` methods
- [] document EVERY method
