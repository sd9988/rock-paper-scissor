# Build Rock, Paper and Scissors Game with React JS

This repository contains code for rock,paper and scissors game in reactjs.

Forked From 
https://github.com/codebucks27/ReactJs-rock-paper-scissors-game

## Building and running on localhost

First install dependencies:

```sh
npm install
```

To run in hot module reloading mode:

```sh
npm start
```

To create a production build:

```sh
npm run build-prod
```

## Runningg

Open the file `dist/index.html` in your browser


## With Docker 

Run 
`docker build -t frontend .`

Run 

`docker run -p 3000:1234 frontend` 

Visit 

`localhost:3000` for frontend. 

Backend and Frontend need to be run together. Use `docker-compose` to run project as a whole.
