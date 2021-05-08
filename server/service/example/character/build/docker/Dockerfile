# build image
FROM golang:1.15.7-alpine3.12 as builder

# enable https:// repositories
RUN apk update \
    && apk add curl \
    && apk add bash \
    && apk add git

# app home
ENV APP_SERVER_HOME /go-mono-api-boilerplate

# copy application
WORKDIR $APP_SERVER_HOME
COPY . .

# build API server
WORKDIR $APP_SERVER_HOME/service/character
RUN go build -o /go/bin/gomo-character-server ./cmd/server
RUN go build -o /go/bin/gomo-character-cli ./cmd/cli

# final image
FROM golang:1.15.7-alpine3.12

RUN apk update \
    && apk add curl \
    && apk add bash

COPY --from=builder /go/bin/gomo-character-server /go/bin
COPY --from=builder /go/bin/gomo-character-cli /go/bin
COPY --from=builder /go-mono-api-boilerplate/service/character/build/docker/entrypoint.sh .
COPY --from=builder /go-mono-api-boilerplate/schema/docs ./schema/docs

RUN pwd
RUN ls -la

# entrypoint
ENTRYPOINT [ "./entrypoint.sh" ]
