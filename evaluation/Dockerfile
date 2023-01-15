FROM golang:1.19-buster as builder

WORKDIR /app/code

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /app/build

FROM gcr.io/distroless/base-debian10 AS app

WORKDIR /

COPY --from=builder /app/build /app

USER nonroot:nonroot

ENTRYPOINT ["/app"]