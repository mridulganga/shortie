FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git make

WORKDIR /app
COPY . .

ARG GITHUB_TOKEN

RUN git config --global url.https://$GITHUB_TOKEN@github.com/.insteadOf https://github.com/

RUN make vendor
RUN make build
RUN go build GOOS=linux CGO_ENABLED=0 -tags netgo -a -v -installsuffix cgo -o bin/shortie main.go


FROM alpine:3
RUN apk update \
    && apk add --no-cache curl wget \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates 2>/dev/null || true

COPY --from=builder /app/bin/shortie /shortie

CMD ["/shortie"]

ENV PORT=8000
EXPOSE $PORT