FROM node:18-alpine

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build

ENV HOST=0.0.0.0
ENV PORT=3000

EXPOSE 3000
CMD ["node", ".output/server/index.mjs"]