APP_NAME := "My Go Desktop App"

clean:
	rm -rf $(APP_NAME).app || true
	rm app || true

build: clean
	packr build -o app
	appify -name=$(APP_NAME) -icon=icon.png app
	sleep 1

run: clean build
	open -a $(APP_NAME).app

debug: clean
	MB_DESKTOP_DEBUG=true go run *.go

install:
	go get github.com/machinebox/appify
	go get -u github.com/gobuffalo/packr/...
