FROM node:22.5.1-alpine

WORKDIR /app

COPY .output /app/

EXPOSE 3000

ENTRYPOINT ["node", "/app/server/index.mjs"]