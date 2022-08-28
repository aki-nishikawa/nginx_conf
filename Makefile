
up:
	docker-compose up -d

down:
	docker-compose down

restart:
	docker-compose restart

log:
	docker logs nginx_conf_nginx_1

shell:
	docker exec -it nginx_conf_nginx_1 bash
