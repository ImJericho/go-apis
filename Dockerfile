FROM golang:1.23-alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]

#To build the image
#docker build -t vivek-go-apis .

#To run the image
#docker run -p 8080:8080 vivek-go-apis