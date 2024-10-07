FROM nginx:alpine

# Remove the default Nginx configuration
RUN rm /etc/nginx/conf.d/default.conf

# Copy custom Nginx configuration
COPY frontend/nginx.conf /etc/nginx/conf.d

# Copy frontend files to Nginx's public directory
COPY frontend/public /usr/share/nginx/html

# Copy static assets
COPY frontend/src/css /usr/share/nginx/html/static/css
COPY frontend/src/js /usr/share/nginx/html/static/js

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
