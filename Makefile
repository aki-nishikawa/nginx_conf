all:
	make down
	make up

up:
	docker-compose up -d

down:
	docker-compose down
	docker rmi nginx_conf_app

restart:
	docker-compose restart

log-n:
	docker logs nginx_conf_nginx_1

log-a:
	docker logs nginx_conf_app_1

shell:
	docker exec -it nginx_conf_nginx_1 bash
