# Start by building the application.
FROM golang:1.18-bullseye as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...
RUN go build -o /go/bin/app

# Now copy it into our base image.
FROM golang:1.18-bullseye
WORKDIR /app
COPY --from=build /go/src/app/ /go/src/app/
COPY --from=build /go/src/app/start.sh /go/bin/app ./
CMD ["bash", "/app/start.sh"]
