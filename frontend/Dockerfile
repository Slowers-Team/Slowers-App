FROM node:18

WORKDIR /app
COPY . .

RUN npm clean-install
RUN npm run build
RUN npm install http-server

CMD npx http-server --port 5173 --proxy http://localhost:5173? ./dist
