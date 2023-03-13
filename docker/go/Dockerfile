FROM golang:1.19.7-alpine3.17
RUN go get github.com/cosmtrek/air
WORKDIR go/src
CMD ["air", "-c", "cdm/public/.air.toml"]
EXPOSE 8000
