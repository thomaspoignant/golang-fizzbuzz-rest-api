#Build distribution from golang:1.11
FROM golang:1.11 as build
ADD . /go/src/gitlab.com/thomas.poignant/fizzbuzz
RUN go install /go/src/gitlab.com/thomas.poignant/fizzbuzz

# Copy to distroless image to have a more secure container
FROM gcr.io/distroless/base
COPY --from=build /go/bin/fizzbuzz /
CMD ["/fizzbuzz"]
EXPOSE 8080