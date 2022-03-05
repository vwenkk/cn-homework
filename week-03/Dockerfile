FROM golang:1.17-alpine3.15 as build

COPY ./ /go/src/project/

WORKDIR /go/src/project/

RUN go build -o /bin/project

FROM alpine

COPY --from=build /bin/project /bin/project

ENTRYPOINT ["/bin/project"]