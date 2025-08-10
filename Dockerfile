FROM golang:1.23.7-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o sqldot

FROM alpine:3.22
RUN apk add --no-cache graphviz && \
    adduser -D -H sqldotuser && \
    mkdir -p /sqldot && \
    chown sqldotuser:sqldotuser /sqldot && \
    chmod 700 /sqldot
COPY --from=build /src/sqldot /usr/local/bin
USER sqldotuser
WORKDIR /sqldot
ENTRYPOINT [ "sqldot" ]
