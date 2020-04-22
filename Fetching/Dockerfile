FROM node:latest
WORKDIR /app
ADD . /app
COPY package*.json ./
RUN npm install
CMD ["npm", "run", "start"]