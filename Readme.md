# *Exchange Rate API*
-------------------
This project allows the user to see the current exchange rates by entering the currency parities that the user wants through the API. **(USD/EUR/TRY)**.

### Features
There are three different applications in this project, two of which are microservices and one is the REST API. The application has a cron job triggered per minute and used the RabbitMQ message queue to put it into a queue structure. The PostgreSQL database was used for the database.

- The first application is the 'er-api-consumer' application. This application first pulls the current currency exchange rate information from the "www.exchangerate-api.com" site. It calculates the parities and writes the current exchange rates to the RabbitMQ message queue every minute.

- The second application is the 'er-rabbit-consumer' application. This application takes the current exchange rate information that our first microservice 'er-api-consumer' wrote to the message queue in RabbitMQ, from this message queue and writes it to the database table established via PostgreSQL.

- Our third application is the 'er-api' application. This application is the application that opens the data on PostgreSQL to users using the REST API. Parity in database with users currency parameters can inquire about their information.

#### Prerequisites and Installing
Primarily, you need to install Go on your system since the project was developed with Go language. Since this system is built with Docker, you need to install the Docker application on your system in order to create images and containers. In order to enable us to up and down our project easily, you need to install the Make tool on your operating system.

#### Build and Launch
>     To launch our project, we first go to the folder location of our project by typing cd . ./ExchangeRateApi on our command line to enter our commands on the default command line and start the project. Firstly To start the application, you must first enter the first of the following commands (make build). This will allow us to build our images to docker. Next, let's get our project up and running by entering our second command (make run). To run both at the same time, we can enter the (make buildandrun) command.

The command needed to create build images:
```sh
make build
```
The command needed to up the entire project:
```sh
make run
```
The command needed to create build images then up the entire project:
```sh
make buildandrun
```
The command needed to stop after up the entire project:
```sh
make stop
```
The command needed to start after stopping the whole project:
```sh
make start
```
The command needed to down after up the entire project:
```sh
make down
```

Warning, the application will not run without the database and rabbitmq services running. After starting these services first, other applications should be launched. There are more make commands in the Makefile.

#### API Requests
http://localhost:3000/api/rates/all
http://localhost:3000/api/rates/usd
http://localhost:3000/api/rates/eur
http://localhost:3000/api/rates/try
http://localhost:3000/api/rates/usd/eur
http://localhost:3000/api/rates/usd/try
http://localhost:3000/api/rates/usd/usd
http://localhost:3000/api/rates/eur/usd
http://localhost:3000/api/rates/eur/try
http://localhost:3000/api/rates/eur/eur
http://localhost:3000/api/rates/try/usd
http://localhost:3000/api/rates/try/eur
http://localhost:3000/api/rates/try/try

Optionally, Postman and Rest Client like this can be used.

#### Build With
- Go - https://go.dev
- Docker - www.docker.com
- PostgreSQL - www.postgresql.org
- RabbitMQ - www.rabbitmq.com

#### Frameworks Used
- Fiber - github.com/gofiber/fiber/v2
- Gorm - github.com/jinzhu/gorm
- Amqp - github.com/streadway/amqp
- Environment - github.com/joho/godotenv


##### Serdar Efe KINALI
*September, 2022*