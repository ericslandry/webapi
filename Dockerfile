FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY bin/myapiserver .
EXPOSE 8080
CMD ["./myapiserver"]