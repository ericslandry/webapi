FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app
COPY bin/myapiserver .
EXPOSE 8080
CMD ["./myapiserver"]