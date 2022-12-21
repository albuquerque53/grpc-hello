up:
	docker-compose -f build/docker-compose.yml up -d || docker compose -f build/docker-compose.yml up -d
down:
	docker-compose -f build/docker-compose.yml down || docker compose -f build/docker-compose.yml down
app:
	docker exec -it grpc_api /bin/bash
api:
	go run ./cmd/main.go -op=serve
call:
	go run ./cmd/main.go -op=call $(filter-out $@,$(MAKECMDGOALS))
%:
	@: