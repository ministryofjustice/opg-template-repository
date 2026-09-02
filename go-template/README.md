# {{APP_TITLE}}

### Major dependencies

- [Go](https://golang.org/) (1.26.x)
- [docker compose](https://docs.docker.com/compose/install/) (>= 2.26.0)
- [Yarn](https://yarnpkg.com/) for frontend assets


## Local development

Build frontend assets and start the application:

```zsh
yarn install --frozen-lockfile
yarn build
make up
```

The application is available at `http://localhost:{{PORT}}{{URL_PREFIX}}/`.

To enable debugging and hot-reloading of Go files:

```zsh
make dev-up
```

Hot-reloading is managed by Air.
