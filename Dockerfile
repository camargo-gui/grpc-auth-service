FROM golang:latest

# set working directory
WORKDIR /go/src/app

# Copy the source code
COPY . . 

#EXPOSE the port
EXPOSE 50051

# Build the Go app
RUN go build -o main cmd/main.go

# Run the executable
CMD ["./main"]