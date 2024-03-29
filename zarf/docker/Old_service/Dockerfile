# build the Go Binary

FROM golang:1.20 as build_sales-api
ENV CGO_ENABLED 0
ARG BUILD_REF

RUN mkdir /service
COPY go.* /service/

WORKDIR /service
RUN go mod download

# Copy the source code into the container
COPY .. /service

# Build the service binary
#WORKDIR /service
RUN go build -ldflags "-X main.build=${BUILD_REF}" -o go_sales_test

# Run the Go Binary in Alpine.
FROM alpine:3.17
ARG BUILD_DATE
ARG BUILD_REF
COPY --from=build_sales-api /service /service/
WORKDIR /service
CMD ["./go_sales_test"]
#CMD ["ls"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="go_sales_test" \
      org.opencontainers.image.authors="William Kennedy <bill@ardanlabs.com>" \
      org.opencontainers.image.source="https://github.com/ardanlabs/service/app/sales-api" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Ardan Labs"