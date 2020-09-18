#
# Build Docker image for BMI
#
# To build image:
# 	docker build -t snexerp/bmi:latest --build-arg VERSION=snapshot --build-arg GIT_COMMIT=$(git rev-parse -q --verify HEAD) .
#
# To run the built image in container
# 	docker run --rm -d --name bmi_server -p 8080:8080 snexerp/bmi:latest serve
#
# Where
#	run     - Ask the Docker engine to run image
#	--rm    - Remove container again after shutdown - exclude this option if you want your container to "stay around"
#	-d      - Detach from execution (run in background) - exclude this option if you want the console output in your terminal
#	--name  - Give the running container a name to easily find it in th UIs
#	-p      - Make internal port 8080 available from the outside (also) on port 8080
#	<image> - exerp/bmi:latest is the image name that corresponds to to the tag given above when building the image
#

# To make the build process a bit smoother a layer is created for all the dependencies to be downloaded. Docker caches the 
# result for later use. This layer only gets updated if `go.mod` or `go.sum` changes.
ARG VERSION=
ARG GIT_COMMIT=

FROM golang:alpine as build
ARG VERSION
ARG GIT_COMMIT

ENV CGO_ENABLED=0
WORKDIR /build

RUN go version
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

# Copy rest of source files needed for build
COPY . .
RUN go build \
	-ldflags="-s -w -X 'github.com/sorennielsen/bmi/internal/system.Version=$VERSION' -X 'github.com/sorennielsen/bmi/internal/system.GitCommit=$GIT_COMMIT'"

RUN go test ./...

FROM scratch
ARG VERSION
LABEL version=$VERSION
WORKDIR /app
COPY --from=build /build/bmi .

ENTRYPOINT ["/app/bmi"]
EXPOSE 8080