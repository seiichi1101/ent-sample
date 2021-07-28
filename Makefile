.PHONY: gremlin
gremlin:
	@docker-compose up graph
console:
	@docker-compose run --rm console
gen:
	@go generate ./ent