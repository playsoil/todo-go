# Run project in your local machine

make a copy of .env.example and fill values.

```bash
cp .env.example .env
```

then:

For Linux/Mac users:

```
make up
```

windows user:

```bash
docker compose up -d
```

Then go to:  
[localhost:8000](http://localhost:8000)

# Run in production

To deploy backend project in production you need:

- SSH access
- a VPS with docker installed

ssh to your vps and follow this commands:

```bash
wget wget https://raw.githubusercontent.com/playsoil/todo-go/refs/heads/master/docker-compose-prod.yml

# change variables with your desired values
COMPOSE_PROJECT_NAME=prod-todo-app IMAGE_TAG=latest DB_USER=prod_todos_user DB_PASSWORD=prod_todos_pass DB_NAME=prod_todos APP_PORT=8080 docker compose -f docker-compose-prod.yml up -d

```

now visit

# Do you want a UI?!

```bash
wget wget https://raw.githubusercontent.com/playsoil/todo-ui/refs/heads/master/docker-compose-prod.yml -O  docker-compose-prod-ui.yml


# change variables with your desired values
COMPOSE_PROJECT_NAME=prod-todo-app IMAGE_TAG=latest docker compose -f docker-compose-prod-ui.yml up -d

```
