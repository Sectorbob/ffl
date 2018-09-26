# Use an alpine image to source certificate data via COPY
FROM alpine:3.6 as alpine

# Ensure we have the standard CA's
RUN apk add -U --no-cache ca-certificates

# Now create the image from scratch
FROM scratch

# Add in the certificates
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Add in the golang zoneinfo packet
COPY zoneinfo.zip /usr/local/Cellar/go/1.11/libexec/lib/time/zoneinfo.zip

# Copy in the service binaries
COPY target/service /service

WORKDIR /
ENTRYPOINT [ "/service" ]