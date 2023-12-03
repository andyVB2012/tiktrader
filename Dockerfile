FROM golang:1.18


WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./
RUN go build -o /TIKTRADER

EXPOSE 3000

CMD [ "/TIKTRADER" ]
