# Go Web App Skeleton Template
This folder is a copy-ready template for a small Go web app using:
- Go HTTP server
- Go HTML templates
- GOV.UK/MOJ frontend assets built with Yarn, Sass and esbuild
- Docker multi-stage build
- Docker Compose local development
- A small GitHub Actions skeleton workflow

## Placeholders to replace
After copying this template into a new repository, replace these placeholders:
| Placeholder | Example | Description |
| --- | --- | --- |
| `{{MODULE_NAME}}` | `github.com/ministryofjustice/my-new-service` | Go module path |
| `{{SERVICE_NAME}}` | `opg-sirius-my-new-service` | Logging/observability service name |
| `{{APP_NAME}}` | `my-new-service` | Docker Compose service, image and binary name |
| `{{APP_TITLE}}` | `My New Service` | Visible page heading |
| `{{PAGE_TITLE}}` | `OPG My New Service` | Browser page title |
| `{{URL_PREFIX}}` | `/my-new-service` | URL path prefix used behind routing/proxy layers |
| `{{PORT}}` | `1234` | Local app port |


## First-time setup after rendering
```zsh
go mod tidy
yarn install --frozen-lockfile
yarn build
docker compose build {{APP_NAME}} or make build
```
## What is intentionally excluded
Do not commit generated or local dependency folders:
- `node_modules/`
- `web/static/`
- `tmp/`
- `test-results/`

The app builds `web/static/` from `web/assets/` during `yarn build` and in the Docker build.

## CI/CD note
The included `.github/workflows/build.yml` is intentionally minimal for a skeleton app. It runs:
- `go test ./...`
- `yarn build`
- `docker compose build {{APP_NAME}}`
Deployment, image publishing, tagging, Snyk, SSM and Jenkins steps should be added later when the service is ready for the full pipeline.
