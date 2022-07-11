#get a base image
FROM golang:1.18

LABEL vincent monneger <mvincem@gmail.com>

WORKDIR /go/src/github.com/vmonneger/FinalProject

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go mod tidy
RUN go get -d -v
RUN go build -o /final-project

EXPOSE 8000

CMD ["/final-project"]