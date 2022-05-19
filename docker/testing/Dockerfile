FROM golang:1.18-alpine AS builder
RUN apk add bash g++
#RUN apk add make

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY Makefile ./

COPY ./src ./src

#RUN make go-build

#ENTRYPOINT ["tail", "-f", "/dev/null"]

#FROM alpine:latest
#RUN apk add bash
#WORKDIR /app
#COPY --from=builder /app/main /app/basket-go-count

COPY ./utils/* ./

ENV DB_HOST=postgres
ENV DB_PORT=5432

CMD ./test.sh
