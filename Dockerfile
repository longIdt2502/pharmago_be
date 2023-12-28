FROM golang
WORKDIR /add
COPY . .
RUN go get .
ENTRYPOINT go run main.go
EXPOSE 8080