package main

import (
	"context"
	"ent-sample/ent"
	"log"

	"entgo.io/ent/dialect"
)

func main() {
	var err error
	client, err := ent.Open(dialect.Gremlin, "http://localhost:8182")
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	user, err := client.User.Create().Save(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("created user: ", user)
	}
}
