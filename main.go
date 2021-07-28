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

	user1, err := client.User.Create().SetName("alice").Save(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	user2, err := client.User.Create().SetName("bob").Save(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, err = client.User.UpdateOne(user1).AddFriend(user2, "20210728").Save(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
