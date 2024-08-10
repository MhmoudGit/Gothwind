tailwind:
	@tailwindcss -i ./cmd/web/assets/css/input.css -o ./cmd/web/assets/css/output.css --watch

template:
	@templ generate --watch --proxy=http://localhost:8000

run:
	@air