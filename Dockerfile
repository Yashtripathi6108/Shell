# Step 1: Use an official Go image (Go 1.22) as the base image
FROM golang:1.22-alpine as builder

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Step 4: Copy the source code to the container
COPY . .

# Step 5: Build the executable file for your shell (Shell.exe)
RUN go build -o Shell.exe main.go

# Step 6: Create a minimal final image based on Alpine for smaller size
FROM alpine:latest

# Step 7: Set the working directory for the final image
WORKDIR /app

# Step 8: Copy the built executable from the builder image to the final image
COPY --from=builder /app/Shell.exe .

# Step 9: Make the executable file executable
RUN chmod +x Shell.exe

# Step 10: Define the command to run when the container starts
ENTRYPOINT ["./Shell.exe"]
