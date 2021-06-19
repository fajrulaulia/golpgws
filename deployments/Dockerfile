FROM golang:1.14.1-alpine as build
ENV CGO_ENABLED=0
ENV GO111MODULE on
WORKDIR /build
RUN apk add bash
COPY . .
RUN go build cmd/main.go 

FROM alpine
ARG SERVER_PORT=8081
RUN apk add bash
WORKDIR /prod
COPY --from=build /build/main  .
COPY --from=build /build/migrations  ./migrations
COPY --from=build /build/depedencies/migrate  .
COPY --from=build /build/scripts/init.up.sh  .
RUN ["chmod", "+x", "init.up.sh"]
RUN ls

EXPOSE $SERVER_PORT
ENTRYPOINT [ "init.up.sh" ]