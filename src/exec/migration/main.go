package main

import (
	"context"

	"src/preload"

	scylla "src/scyllaDB"

	"github.com/scylladb/gocqlx/v2/migrate"
)

func main() {
	ctx := context.Background()

	env, err := preload.Load()
	if err != nil {
		panic(err)
	}

	manager := scylla.CreateManager(env)

	if err = manager.CreateKeySpace(env.Scyllakeyspace); err != nil {
		panic(err)
	}

	session, err := manager.Connect()
	if err != nil {
		panic(err)
	}

	if err = migrate.Migrate(ctx, session, env.Scyllamigration); err != nil {
		panic(err)
	}
}
