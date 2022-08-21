# go-redis-chat-app
A basic chat application built with Golang, Websocket &amp; Redis

### The app could be run in 2 ways:
1. locally on the host OS with seperate services for redis, UI & API
2. through docker-compose, with 2 containers


### Running the app Through Docker
1. Rebuild the UI project
    - cd into `client`
    - clean the `node_modules` directory
    - run `yarn install`
    - run `yarn build`
2. cd into the project root
3. run `docker-compose up --build`
4. The app is now up & running. Go to http://localhost:5555 in the browser. 5555 is the SERVER_ADDRESS defined in the `.env` file.


## Reference template from https://github.com/redis-developer