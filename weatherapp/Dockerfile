FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o weatherapp .

FROM alpine:latest

# Set the working directory in the new stage
WORKDIR /app

COPY --from=builder /app/weatherapp .

EXPOSE 8080

# Command to run the application
CMD ["./weatherapp"]