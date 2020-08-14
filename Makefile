test:
	docker-compose run --entrypoint="go test" test-web-backend
dev:
	CompileDaemon -command="./tahor"