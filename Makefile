run:
	@air

tailwind:
	@npx tailwindcss -i ./app/internal/assets/input.css -o ./app/internal/assets/dist/output.css --watch

templ:
	@templ generate --watch --proxy="http://localhost:3000"

init:
	@git clone git@github.com:pocketbase/js-sdk.git app/internal/assets/dist/pocketbase
	@yarn
	@go mod download
	@go install github.com/a-h/templ/cmd/templ@latest
	@go mod tidy
