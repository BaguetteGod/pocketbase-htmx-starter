live/server:
	@air

live/tailwind:
	@npx tailwindcss -i ./app/internal/assets/input.css -o ./app/internal/assets/dist/output.css --watch

live/templ:
	@templ generate --watch --proxy="http://localhost:8090" --open-browser=false -v

live/sync_assets:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "app/internal/assets" \
	--build.include_ext "js,css"

live:
	@make -j3 live/templ live/server live/tailwind live/sync_assets

init:
	@git clone git@github.com:pocketbase/js-sdk.git app/internal/assets/dist/pocketbase
	@yarn
	@go mod download
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/air-verse/air@latest
	@go mod tidy

