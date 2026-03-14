FROM nginx:alpine

# Remove default nginx static assets
RUN rm -rf /usr/share/nginx/html/*

# Copy the vanilla frontend files to Nginx public directory
COPY . /usr/share/nginx/html

# Expose port 80 (will be mapped to 3000 in docker-compose)
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
