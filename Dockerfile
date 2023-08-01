# BUILD
FROM golang:1.20 As build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o app cmd/app/main.go
RUN go build -o migrate cmd/migrate/main.go

# DEPLOY
FROM ubuntu:22.04
WORKDIR /app
COPY --from=build /app/app ./
COPY --from=build /app/migrate ./
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
CMD ["./app"]
