FROM golang:1.19-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o crwl .
RUN chmod a+x ./crwl
ENTRYPOINT ["/app/crwl"]