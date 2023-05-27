package main

import (
	"context"

	"src/config"

	scylla "src/scyllaDB"

	"github.com/scylladb/gocqlx/v2/migrate"
)

func main() {
	ctx := context.Background()

	config, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	manager := scylla.CreateManager(config)

	if err = manager.CreateKeySpace(config.ScyllaKeySpace); err != nil {
		panic(err)
	}

	session, err := manager.Connect()
	if err != nil {
		panic(err)
	}

	if err = migrate.Migrate(ctx, session, config.ScyllaMigrations); err != nil {
		panic(err)
	}
}
