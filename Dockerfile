FROM golang:1.22.4 AS base
WORKDIR /app
ADD . /app/
RUN go mod download
RUN make build
FROM alpine AS release
WORKDIR /app
COPY --from=base /app/bin/main /app/
ENTRYPOINT [ "./main" ]
