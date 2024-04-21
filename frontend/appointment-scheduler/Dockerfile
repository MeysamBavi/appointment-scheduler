FROM node:20-alpine

WORKDIR /app

COPY package*.json /app/

RUN npm i --legacy-peer-deps

COPY . .

CMD ["npm", "run", "dev"]