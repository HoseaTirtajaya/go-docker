FROM golang:latest

WORKDIR /app/server

COPY . .

RUN go get ./... 
RUN go build 

CMD [ "go", "run", "main.go" ]
