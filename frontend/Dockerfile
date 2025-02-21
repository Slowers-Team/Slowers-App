FROM node:18

WORKDIR /app
COPY . .

RUN npm clean-install
RUN npm run build
RUN npm install http-server

EXPOSE 5173

CMD npx http-server dist --port 5173 --proxy $BACKEND_URL
