FROM golang:alpine AS build
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -a -installsuffix cgo -o gobadge cmd/*.go

FROM scratch
COPY --from=build /app/gobadge gobadge
CMD ["gobadge"]