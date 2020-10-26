FROM golang:alpine AS build
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o gobadge cmd/*.go

FROM scratch
COPY --from=build /app /app
CMD ["/app/gobadge"]