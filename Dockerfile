#
# Build Buzz for BMI
#
# To build image:
# 	docker build -t exerp/bmi:latest .
#
# To run the built image in container
# 	docker run --rm -d --name bmi_server -p 8080:8080 exerp/bmi:latest serve
#
# Where
#	run     - Ask the Docker engine to run image
#	--rm    - Remove container again after shutdown - exclude this option if you want your container to "stay around"
#	-d      - Detach from execution (run in background) - exclude this option if you want the console output in your terminal
#	--name  - Give the running container a name to easily find it in th UIs
#	-p      - Make internal port 8080 available from the outside (also) on port 8080
#	<image> - exerp/bmi:latest is the image name that corresponds to to the tag given above when building the image
#

# To make the build process a bit smoother a layer is created all dependencies downloaded, and make Docker cache the result for later use. This layer only gets updated if
# `go.mod` or `go.sum` changes.
FROM golang:alpine as build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of source files needed for build
COPY . .

# Need to set this to allow using "scratch" for runtime
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go test ./...

RUN go build

FROM scratch
WORKDIR /app
COPY --from=build /build/bmi .

ENTRYPOINT ["/app/bmi"]
EXPOSE 8080