FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . ./
RUN go mod tidy
RUN  go build -v --installsuffix cgo --ldflags="-s" -o hrm
RUN ./hrm
