FROM golang:1.14.1-alpine as build
ENV CGO_ENABLED=0
ENV GO111MODULE on
RUN apk add bash
WORKDIR /build
COPY . .
RUN go build cmd/main.go 

FROM alpine
RUN apk add bash
WORKDIR /prod
COPY --from=build /build/main  .
COPY --from=build /build/migrations  ./migrations
COPY --from=build /build/bin/migrate  .
RUN cd migrations/ && ls
RUN ls
EXPOSE $SERVER_PORT

ENTRYPOINT ["/bin/sh","-c", "chmod +x migrate && ./migrate && sleep 5 && ./main" ]