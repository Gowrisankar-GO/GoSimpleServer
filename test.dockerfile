FROM alpine:latest

WORKDIR /app

# RUN apk update

COPY simple_go_server .

# RUN chmod +x ./simple_go_server

CMD [ "./simple_go_server" ]

# EXPOSE 8080