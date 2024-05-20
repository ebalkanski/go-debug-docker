FROM golang:1.22.1

# build watcher
RUN go install github.com/ysmood/kit/cmd/guard@v0.25.11

WORKDIR /microservices/debug_docker

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /server ./cmd/...

ENV TZ=Europe/Sofia

EXPOSE 80

ENTRYPOINT ["sh", "-c", "/go/bin/guard -w '**/*.go' -- ./run.sh"]
