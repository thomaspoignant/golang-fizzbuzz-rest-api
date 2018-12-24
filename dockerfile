#Build distribution from golang:1.11
FROM golang:1.11 as build
ADD . /go/src/github.com/thomaspoignant/golang-fizzbuzz-rest-api
RUN go install /go/src/github.com/thomaspoignant/golang-fizzbuzz-rest-api

# Copy to distroless image to have a more secure container
FROM gcr.io/distroless/base
COPY --from=build /go/bin/golang-fizzbuzz-rest-api /
CMD ["/golang-fizzbuzz-rest-api"]
EXPOSE 8080