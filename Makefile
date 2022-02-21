fl=-f ./container.yml
bin=./chan.exe
dev_reload=CompileDaemon \
					 -build="go build -o ./chan.exe ./main.go" \
					 -command="./chan.exe" \
					 -exclude-dir=".git" \
					 -exclude-dir="log" \
					 -exclude-dir="./db" \
					 -exclude="(.*\.exe)" \
					 -pattern="(.*)" \
					 -verbose


all: run

run: build
	$(bin)

clean:
	rm $(bin)

build:
	go build -o $(bin) main.go

dev_reload:
	$(dev_reload)
