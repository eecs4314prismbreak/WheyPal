FROM golang:latest
WORKDIR /go/src/ak-rest-api/
RUN go get -d -v golang.org/x/net/html  
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/ak-rest-api .
# Expose port 8080 to the outside world
EXPOSE 8080
CMD ["./app"]  
