FROM ubuntu

WORKDIR /gateway

COPY config/config.yaml config/
COPY gateway /gateway

EXPOSE 8000

CMD ["./gateway"]