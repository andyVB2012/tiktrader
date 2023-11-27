FROM golang:1.16.3-alpine3.13

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./
RUN go build -o /TIKTRADER

EXPOSE 3000

CMD [ "/TIKTRADER" ]
