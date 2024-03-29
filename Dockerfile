FROM golang:1.19-alpine AS build-stage

#Set up dependencies
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

#Copy and compile project
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /hosttech-ddns

#Move binary to a leaner image
FROM gcr.io/distroless/base-debian11 AS release-page
WORKDIR /
COPY --from=build-stage /hosttech-ddns /hosttech-ddns
USER nonroot:nonroot

#Start application
CMD ["/hosttech-ddns"]