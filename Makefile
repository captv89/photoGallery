run:
	templ generate -path web/tfs
	npx tailwindcss -i tailwind-input.css -o ./web/static/css/tail-style.css
	go run main.go

build:
	templ generate -path web/tfs
	npx tailwindcss -i tailwind-input.css -o ./web/static/css/tail-style.css
	go build -o bin/app main.go