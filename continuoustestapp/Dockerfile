FROM alpine:latest

RUN apk update && apk add --no-cache curl
RUN apk add --no-cache --upgrade bash

WORKDIR /app

COPY continuous_test.sh /app/continuous_test.sh

RUN chmod +x /app/continuous_test.sh

CMD ["/bin/sh", "-c", "/app/continuous_test.sh"]
