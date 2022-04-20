FROM golang:1.17 as builder

COPY / /build

WORKDIR /build

RUN make build

FROM gcr.io/distroless/base

ENV PORT 8080
ENV HOST 0.0.0.0

COPY --from=builder /build/bin/feature-evaluation-api /app/feature-evaluation-api

USER 65532
CMD [ "/app/feature-evaluation-api"]


