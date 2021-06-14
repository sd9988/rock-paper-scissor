# Name the node stage "builder"
FROM node:15 AS builder

# Set working directory
WORKDIR /app

# Copy our node module specification
COPY package.json package.json

RUN npm install -g parcel-bundler

RUN npm install --production


# Copy all files from current directory to working dir in image
# Except the one defined in '.dockerignore'
COPY . .


CMD [ "npm","start" ]