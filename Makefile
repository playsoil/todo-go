up:
	docker compose up -d
stop:
	docker compose 
stats:
	docker stats 

production_up:
	 COMPOSE_PROJECT_NAME=prod-todo-app IMAGE_TAG=latest DB_USER=prod_todos_user DB_PASSWORD=prod_todos_pass DB_NAME=prod_todos APP_PORT=8030 docker compose -f docker-compose-prod.yml up -d

production_stop:
	 COMPOSE_PROJECT_NAME=prod-todo-app IMAGE_TAG=latest DB_USER=prod_todos_user DB_PASSWORD=prod_todos_pass DB_NAME=prod_todos APP_PORT=8030 docker compose -f docker-compose-prod.yml stop

production_stats:
	 COMPOSE_PROJECT_NAME=prod-todo-app  docker compose -f docker-compose-prod.yml stats 
