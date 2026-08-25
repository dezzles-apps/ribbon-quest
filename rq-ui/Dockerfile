FROM node:lts-alpine AS build-stage
WORKDIR /
COPY package*.json ./
RUN export NODE_OPTIONS=--openssl-legacy-provider
RUN npm install
COPY . .
RUN npm run build

# production stage
FROM nginx:stable-alpine AS production-stage
COPY --from=build-stage /dist /usr/share/nginx/html
#COPY nginx.conf /etc/nginx/nginx.conf
COPY conf/prd.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]