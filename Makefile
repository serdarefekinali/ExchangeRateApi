build:  
	cd er-api && docker build -t erapi_img .
	cd er-api-consumer && docker build -t erapiconsumer_img .
	cd er-rabbit-consumer && docker build -t errabbitconsumer_img .

postgres:
	docker-compose up -d postgres
rabbitmq:
	docker-compose up -d rabbitmq
	timeout 3
erapi:
	docker-compose up -d erapi
erapiconsumer:
	docker-compose up -d erapiconsumer
errabbitconsumer:
	docker-compose up -d errabbitconsumer
up:
	docker-compose -f docker-compose.yml up
down:
	docker-compose -f docker-compose.yml down -v
	
run: rabbitmq postgres erapi erapiconsumer errabbitconsumer

buildandrun: build rabbitmq postgres erapi erapiconsumer errabbitconsumer

startpostgres: 
	docker-compose start postgres
startrabbitmq:
	docker-compose start rabbitmq
	timeout 3
starterapi:
	docker-compose start erapi
starterapiconsumer:
	docker-compose start erapiconsumer
starterrabbitconsumer:
	docker-compose start errabbitconsumer

start: startrabbitmq startpostgres starterapi starterapiconsumer starterrabbitconsumer

stop:
	docker-compose -f docker-compose.yml stop