FROM golang:1.23.0-alpine AS build

# Set destination for COPY
WORKDIR /app

# download dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy the files to the container
COPY . .

# Compile the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /non-sei-solo ./cmd/main.go


# Stage 2: Create a minimal production image
FROM alpine:latest

# Install dependencies for the application
RUN apk add --no-cache sqlite-libs

# create a directory for the application
WORKDIR /app

# Copy the application from stage 1
COPY --from=build /non-sei-solo /non-sei-solo

# Copy .env file
COPY .env .env

EXPOSE 1323

# run the application
CMD ["/non-sei-solo"]
