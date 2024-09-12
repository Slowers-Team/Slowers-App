FROM node:20

WORKDIR /frontend

COPY package*.json ./
RUN npm ci

COPY . ./

EXPOSE 5173

RUN npm run build
CMD ["npm", "run", "dev"] 
