FROM node:20-alpine

WORKDIR /app

COPY package*.json /app/

RUN npm i --legacy-peer-deps

COPY ./.eslintrc.cjs ./.eslintrc.cjs
COPY ./index.html ./index.html
COPY ./package.json ./package.json
COPY ./package-lock.json ./package-lock.json
COPY ./vite.config.js ./vite.config.js
COPY ./public ./public
COPY ./src ./src

CMD ["npm", "run", "dev"]