# build stage
FROM node:16-alpine as build

WORKDIR /app

COPY . .
RUN npm i
RUN npm run build

# run stage
FROM node:16-alpine as run

WORKDIR /app

COPY --from=build /app/package*.json ./
COPY --from=build /app/build ./

EXPOSE 3000
CMD ["node", "./index.js"]
