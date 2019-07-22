default:
	@echo "=============building Local API============="
	docker build -t numberwang . --target=local

up: default
	@echo "=============starting api locally============="
	docker-compose up

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	rm -f api
	docker system prune -f
	docker volume prune -f