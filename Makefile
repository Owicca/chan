bin=/chan.exe
dev_reload=CompileDaemon \
					 -build="go build -o /chan.exe ./cmd/chan/main.go" \
					 -command="/chan.exe" \
					 -pattern="(.*)\.(go|tpl|js|css|json|sql)" \
					 -exclude-dir=./log/ \
					 -exclude-dir=./static/media/ \
					 -exclude="./chan.exe" \
					 -polling=true \
					 -polling-interval=300 \
					 -color \
					 -verbose

dev_reload_no_pool=CompileDaemon \
					 -build="go build -o /chan.exe ./cmd/chan/main.go" \
					 -command="/chan.exe" \
					 -pattern="(.*)\.(go|tpl|js|css|json|sql)" \
					 -exclude-dir=./log/ \
					 -exclude-dir=./static/media/ \
					 -exclude="./chan.exe" \
					 -color \
					 -verbose

all: run

run: build
	$(bin)

clean:
	rm $(bin)

build:
	go build -o $(bin) ./cmd/chan/main.go

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
