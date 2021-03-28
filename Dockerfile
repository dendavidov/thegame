FROM golang:1.15-alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -v ./cmd/thegame

FROM alpine

RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/thegame /app/
WORKDIR /app

CMD ["./thegame", "-s", "f2fDroCRnFqF!&tPZF&Qr&jRAP9qBDMP#iJ4L%#3G7*wAb6juGBTGJGar!zy8GWVjT55q"]

# QUEST 2:
#CMD ["./thegame", "-s", "SECRET PHRASE FROM THE QUEST 2"]
