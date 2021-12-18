FROM alpine:latest
RUN apk update && rm -rf /var/cache/apk/*
COPY ./sellpump.crt /usr/local/share/ca-certificates/sellpump.crt
RUN update-ca-certificates
COPY user user
COPY config config
COPY logs logs
CMD ["./user"]
