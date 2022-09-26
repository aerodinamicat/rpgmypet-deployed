serve:
	docker-compose build server && docker-compose up server

data:
	docker-compose -f docker-compose.yml stop db
	docker-compose -f docker-compose.yml rm -f -v db
	docker-compose -f docker-compose.yml up --remove-orphans -d db
	go run cmd/data/main.go

deploy:
	docker build -t rpgmypet_server .
	docker tag rpgmypet_server:latest eu.gcr.io/rpgmypet/rpgmypet_server
	docker push eu.gcr.io/rpgmypet/rpgmypet_server