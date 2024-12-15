FROM node:18 AS build_stage

WORKDIR /app

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
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]