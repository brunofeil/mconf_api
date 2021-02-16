## We specify the base image we need for our
## go application
FROM golang:1.15.8
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /api
## We copy everything in the root directory
## into our /app directory
ADD . /api
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /api
# Download all the dependencies
RUN go get -d -v ./...
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .
EXPOSE 8080
## Our start command which kicks off
## our newly created binary executable
CMD ["/api/main"]