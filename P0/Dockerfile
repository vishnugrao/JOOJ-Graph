# Uses golang latest version as base image
FROM golang:latest

# Goes to the app directory 
WORKDIR /app

# Copy package.json and package-lock.json (if available)
# COPY package*.json ./

# install app dependencies
# RUN npm install

# Copy rest of the App into the container
COPY . .

# Set port environment (if needed)
# ENV PORT = 9000

# Expose the port so that the computer can access it
# EXPOSE 9000

RUN go build -o main .

# Run the app
CMD ["./main"]