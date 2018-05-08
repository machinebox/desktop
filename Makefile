APP_NAME := "My Go Desktop App"

clean:
	rm -rf $(APP_NAME).app || true
	rm app || true
	rm bindata.go || true

build: clean
	go generate
	go build -o app
	appify -name=$(APP_NAME) -icon=icon.png app

run: clean build
	sleep 1
	open -a $(APP_NAME).app

debug: clean build
	sleep 1
	MB_DESKTOP_DEBUG=true open -a $(APP_NAME).app

install:
	go get github.com/machinebox/appify
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/elazarl/go-bindata-assetfs/...
