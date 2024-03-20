FROM golang as builder
ARG VERSION
ARG APPNAME
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -x -ldflags "-w -X main.version=${VERSION}" -o bin/${APPNAME} ./cmd/${APPNAME}/

FROM alpine
ARG APPNAME
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser
COPY --chown=appuser:appuser --from=builder /app/bin/${APPNAME} .
USER appuser
EXPOSE 8080
RUN echo "./${APPNAME}"
CMD ./${APPNAME}