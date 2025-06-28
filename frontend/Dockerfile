# Use official Node image as base
FROM node:20

# Create app directory
WORKDIR /app

# Copy package.json and pnpm-lock.yaml (or package-lock.json)
COPY package.json pnpm-lock.yaml ./

# Install dependencies
RUN npm install -g pnpm && pnpm install

# Copy source code
COPY . .

# Start dev server (for development)
CMD ["pnpm", "dev"]

# Or for production, build and serve static files
