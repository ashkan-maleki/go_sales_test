# build the Go Binary
FROM golang:1.20 as build_sales-api
ENV CGO_ENABLED 0
ARG BUILD_REF

# Copy the source code into the container
COPY . /service

# Build the service binary
WORKDIR /service/app/services/sales-api
RUN go build -ldflags "-X main.build=${BUILD_REF}"

# Run the Go Binary in Alpine.
FROM alpine:3.17
ARG BUILD_DATE
ARG BUILD_REF
COPY --from=build_sales-api /service/app/services/sales-api/sales-api /service/sales-api
WORKDIR /service
CMD ["./sales-api"]
#CMD ["ls"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="go-sales-api" \
      org.opencontainers.image.authors="Ashkan Maleki <me@ashkan.com>" \
      org.opencontainers.image.source="https://github.com/mamalmaleki/go_sales_test/app/sales-api" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Ashkan Maleki"