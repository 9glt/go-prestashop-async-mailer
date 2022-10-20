FROM golang:1.17-alpine as compiler

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /mailer-app


FROM scratch
WORKDIR /
COPY --from=compiler /mailer-app /

EXPOSE 80

ENTRYPOINT ["/mailer-app"]