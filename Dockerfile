FROM alpine:3.6
WORKDIR /root

ADD ./cmd/main .
ADD ./.env .
ADD ./public ./public
ADD ./view ./view

EXPOSE 8000

# herokuではENTRYPOINTではなくCMDを使う
CMD ["./main"]
