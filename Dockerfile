FROM golang:1.17.1-buster
WORKDIR /src/nitd-results/
ADD . .
RUN ["go", "mod", "tidy"]
RUN ["go", "mod", "vendor"]
CMD ["go", "run", "main.go"]