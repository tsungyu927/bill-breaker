FROM node:18 AS build_stage

WORKDIR /app

# Env for app build
ARG VITE_API_HOST
ENV VITE_API_HOST=$VITE_API_HOST

# Install pnpm
RUN npm install -g pnpm

# Install dependencies and build the app
COPY pnpm-lock.yaml ./
COPY package.json ./
RUN pnpm install --frozen-lockfile
COPY . .
RUN pnpm run build

# Use Nginx for serving static files
FROM nginx:alpine
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
COPY --from=build_stage /app/dist /usr/share/nginx/html
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]