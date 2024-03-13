FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
ARG APPNAME
ENV APPNAME=${APPNAME}
COPY bin/${APPNAME} .
EXPOSE 8080
CMD ["./${APPNAME}"]