destroy:
	docker-compose down --rmi local

serve:
	docker-compose build && docker-compose up -d

deploy:
	docker build -t rpgmypet_server .
	docker tag rpgmypet_server:latest eu.gcr.io/rpgmypet/rpgmypet_server
	docker push eu.gcr.io/rpgmypet/rpgmypet_server

doc:
	swagger generate spec -o ./internal/swaggerui/swagger.json
