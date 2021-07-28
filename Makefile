.PHONY: gremlin
gremlin:
	@docker-compose up --build graph
console:
	@docker-compose run --rm console
gen:
	@go generate ./ent