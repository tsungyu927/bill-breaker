# Build stage
FROM golang:1.23.2 as build_stage

WORKDIR /app
COPY . .

# Fetch dependencies and build the app
# Set CGO_ENABLED=0 to be able to run in scratch image below
RUN go mod download 
RUN CGO_ENABLED=0 go build -o main .

# Runtime stage
FROM scratch as final_stage

COPY --from=build_stage /app/main /main

# Expose port 8080
EXPOSE 8080

# Start the service
CMD ["/main"]
