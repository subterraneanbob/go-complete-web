set windows-shell := ["pwsh", "-c"]

default: watch

watch:
  @wgo \
    -file=".go" \
    -file=".templ" \
    -xfile="_templ.go" \
    templ generate -lazy -log-level=warn \
    :: \
    go run ./cmd/

run: generate-templates run-web

generate-templates:
  @templ generate

run-web:
  @go run ./cmd/