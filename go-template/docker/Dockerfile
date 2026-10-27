FROM golang:1.27.3-alpine@sha256:d200d2dd76b03a99b52651fd1325aff7dcd35422c35c01bac56571b0507f3deb AS base

RUN adduser \
--disabled-password \
--gecos "" \
--home "/nonexistent" \
--shell "/sbin/nologin" \
--no-create-home \
--uid 65532 \
app-user

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux

RUN update-ca-certificates
RUN apk add --no-cache tzdata

FROM node:24.22.1-alpine@sha256:06d7adf3e0ef61a13f1d28682726366f97a15aa1888c934d3ecec86c33918b53 AS asset-env
WORKDIR /app

RUN mkdir -p web/static

COPY web/assets web/assets
COPY package.json .
COPY .npmrc .

RUN apk add --no-cache git
RUN npm ci --ignore-scripts
RUN npm run build

FROM base AS healthcheck-build
WORKDIR /app

COPY healthcheck healthcheck

WORKDIR /app/healthcheck

RUN go build -a -installsuffix cgo -o /go/bin/healthcheck

FROM base AS dev
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY internal internal
COPY web/template web/template
COPY --from=asset-env /app/web/static web/static
COPY --from=healthcheck-build /go/bin/healthcheck healthcheck

RUN go install github.com/air-verse/air@latest && go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE {{PORT}}
EXPOSE 2345

ENTRYPOINT ["air"]

FROM base AS build-env
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY internal internal
COPY web/template web/template
COPY --from=asset-env /app/web/static web/static

RUN go build -a -installsuffix cgo -o /go/bin/{{APP_NAME}} .

FROM scratch AS prod

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ca-certificates* /etc/
COPY --from=base /etc/ssl/* /etc/
COPY --from=base /usr/share/ca-certificates/ /usr/share/ca-certificates/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=build-env /go/bin/{{APP_NAME}} {{APP_NAME}}
COPY --from=build-env /app/web/template web/template
COPY --from=asset-env /app/web/static web/static
COPY --from=healthcheck-build /go/bin/healthcheck healthcheck

USER app-user:app-user

HEALTHCHECK --interval=5s --timeout=5s --start-period=5s --retries=5 CMD [ "/healthcheck" ]

ENTRYPOINT ["./{{APP_NAME}}"]

