FROM golang:alpine AS build
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -a -installsuffix cgo -o gobadge cmd/*.go

FROM --platform=linux/arm64 alpine:latest
COPY --from=build /app/gobadge /usr/bin/
CMD ["gobadge"]