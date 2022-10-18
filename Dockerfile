FROM golang:1.18 as build-stage
ENV APP /go/src/github.com/paoloo/correios
ADD . $APP/
WORKDIR $APP
RUN go get && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o correios .

FROM scratch
COPY --from=build-stage /go/src/github.com/paoloo/correios/correios /correios
CMD ["/correios"]
