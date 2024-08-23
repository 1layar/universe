###############################################################################
#                                INSTALLATION
###############################################################################
FROM golang:1.22-alpine AS builder

ARG APP_DIR
ARG SERVICE_NAME

WORKDIR /app

# Copy the source code
COPY ${APP_DIR} .

# Sync the dependencies
RUN go work sync

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./internal/$SERVICE_NAME


###############################################################################
#                                   START
###############################################################################
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]