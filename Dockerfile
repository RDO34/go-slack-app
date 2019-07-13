FROM alpine:3.4

RUN apk add --no-cache ca-certificates

ADD go-slack-app testbot
RUN chmod +x testbot

CMD ["./testbot"]
