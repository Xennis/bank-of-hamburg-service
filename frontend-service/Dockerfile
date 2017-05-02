FROM node

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Install app dependencies
COPY package.json /usr/src/app/
RUN npm install
RUN npm install -g serve

# Bundle app source
COPY . /usr/src/app

# Build and optimize react app
RUN npm run build


EXPOSE 5000
CMD serve -s build
