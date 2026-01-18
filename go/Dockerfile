# Stage 1: Builder
FROM golang:1.24-bookworm AS builder

# ติดตั้ง libtesseract และ wget
RUN apt-get update && apt-get install -y \
    libtesseract-dev \
    wget \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

RUN mkdir -p /app/my_tessdata

RUN wget https://github.com/tesseract-ocr/tessdata_fast/raw/main/eng.traineddata -O /app/my_tessdata/eng.traineddata && \
    wget https://github.com/tesseract-ocr/tessdata_fast/raw/main/tha.traineddata -O /app/my_tessdata/tha.traineddata

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    tesseract-ocr \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

ENV OMP_THREAD_LIMIT=1
ENV MALLOC_TRIM_THRESHOLD_=131072
RUN mkdir -p /usr/local/share/tessdataฆ

COPY --from=builder /app/my_tessdata/*.traineddata /usr/local/share/tessdata/

ENV TESSDATA_PREFIX=/usr/local/share/tessdata/

WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8081
CMD ["./main"]