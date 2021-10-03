FROM golang:1.17 as  build


ADD . /dockerdev
WORKDIR /dockerdev

RUN go build -o /easytunnel ./cmd

# Final stage
FROM debian:buster

EXPOSE 8000

RUN mkdir "/easytunnel"
WORKDIR "/easytunnel"

COPY --from=build /easytunnel ./

CMD ["./easytunnel"]