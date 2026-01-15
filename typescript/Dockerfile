FROM node:18-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install --only=production

FROM node:18-alpine
WORKDIR /app

RUN apk add --no-cache \
    tesseract-ocr \
    tesseract-ocr-data-eng \
    tesseract-ocr-data-tha

COPY --from=builder /app/node_modules ./node_modules
COPY . .


EXPOSE 8080
CMD ["npm", "start"]