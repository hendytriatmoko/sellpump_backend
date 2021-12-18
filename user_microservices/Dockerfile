FROM alpine:latest
RUN apk update && rm -rf /var/cache/apk/*
COPY ./dna.crt /usr/local/share/ca-certificates/dna.crt
RUN update-ca-certificates
COPY user user
COPY config config
COPY logs logs
CMD ["./user"]
