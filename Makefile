bin=./chan.exe
dev_reload=CompileDaemon \
					 -build="go build -o ./chan.exe ./main.go" \
					 -command="./chan.exe" \
					 -exclude-dir=".git" \
					 -exclude-dir="log" \
					 -exclude-dir="./static/media" \
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

clean_media:
	-rm static/media/*.png
	-rm static/media/*.jpg
	-rm static/media/*.jpeg
	-rm static/media/*.webp
	-rm static/media/*.gif
	-rm static/media/*.mp4
	-rm static/media/*.webm
