APP_NAME := "My Go Desktop App"

clean:
	rm -rf $(APP_NAME).app || true
	rm app || true

build: clean
	go generate
	go build -o app
	bin/appify app $(APP_NAME)
	open -a $(APP_NAME).app
