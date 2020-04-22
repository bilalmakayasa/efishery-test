# Installation Guide

Welcome to my efishery-test apps, there are 2 framework that I Use:

1.  Golang

   This apps is handling register, login and checking claims endpoint, It runs on port 8081

2. Node

   This apps is handling fetch, aggregate and checking claims endpoint, It runs on port 8082

   

   So, before you run this apps, plase make sure your machine could run GO and NPM

## Running By Docker

Please install and runing Docker Daemon tools if you decide to run this app with docker, If you are sure the Docker daemon is running, make sure you are in the main directory and simply run: 

```bash
Docker-compose up
```

If your display are showing  this

```
Starting efishery-test_fetching_1 ... done
Starting efishery-test_auth_1     ... done
```

Then, you could use both of the apps from localhost port 8081(Auth) and 8082(Fetching Data)

## Conventional Running

There are 2 apps that you should run, so you will need 2 terminal windows / tab to run it simultaneously

1. Auth Apps (GO)

   First, you could access the /Auth directory, after that, simply run:

   ```bash
   go run main.go
   ```

   The Auth apps will be running on localhost:8081

2. Fetching Data Apps (NODE)

   With your second terminal window/tab access the directiory of /Fetching , then run:

   ```
   npm install
   ```

   After the Installation process done, simply run:

   ```
   npm run start
   ```

   The Fetching apps will be running on localhost:8082

