FROM harbor.nicleary.com/dockerhub/golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /discord-metrics-server

EXPOSE 8080

CMD ["/discord-metrics-server"]