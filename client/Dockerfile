FROM node:latest

WORKDIR /app/client

COPY package*.json ./

RUN npm install -g pnpm
RUN pnpm install

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev"]
