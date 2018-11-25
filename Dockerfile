FROM golang:1.11.2 as dep

ENV GO111MODULE=on

WORKDIR /work

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY tools.go .

RUN go run tools.go

FROM golang:1.11.2 as build

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest as app

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /work/app .
COPY --from=build /work/bin .

EXPOSE 8080

CMD ["./app"]
