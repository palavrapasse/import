FROM import:latest AS import
FROM node:18.14-alpine AS builder
WORKDIR /app

COPY --from=import /app/import .

COPY index.js ./
COPY package.json ./

RUN npm install

CMD [ "index.js" ]