# build stage #0
# define base image
FROM golang:latest AS stage0
# create work directory
RUN mkdir /app
# switch to work directory
WORKDIR /app
# copy all files
ADD . .
RUN go mod download
# build app
RUN CGO_ENABLED=0 GOOS=linux GODEBUG=netdns=1 GOARCH=amd64 go build -o app .
# build stage #1
FROM alpine:latest AS stage1
# create work directory
RUN mkdir /app
# switch to work directory
WORKDIR /app
# copy app artifacts to current directory
COPY --from=stage0 /app/app .
# run app
CMD ["sh", "-c", "./app"]