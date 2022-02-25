FROM golang:1.17.1-buster

ENV APP_NAME nitd-results
ENV CMD_PATH main.go

WORKDIR /src/nitd-results/
ADD . .

RUN make install_swagger
RUN make swagger
RUN CGO_ENABLED=0 go build -v -o ./cmd/$APP_NAME

CMD ./cmd/$APP_NAME