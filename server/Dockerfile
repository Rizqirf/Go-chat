# syntax=docker/dockerfile:1

FROM golang:1.20


# set /app to be the active directory
WORKDIR /app

# copy all files from outside container, into the container
COPY . .

# download dependencies
RUN go mod tidy

# build binary
RUN go build -o binary

# set the entry point of the binary
ENTRYPOINT ["/app/binary"]

EXPOSE 8080