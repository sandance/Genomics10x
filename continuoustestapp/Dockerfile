FROM ubuntu:latest

RUN apt-get update && apt-get install -y curl

WORKDIR /app

COPY continuous_test.sh /app/continuous_test.sh

RUN chmod +x /app/continuous_test.sh

CMD ["/bin/bash", "-c", "/app/continuous_test.sh"]
