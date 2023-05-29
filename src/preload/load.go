package preload

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Scyllahost      []string ""
	Scyllakeyspace  string   ""
	Scyllamigration string   ""
}

func Load() (Env, error) {
	var env Env
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	env.Scyllahost = append(env.Scyllahost, os.Getenv("SCYLLA_HOST"))
	env.Scyllakeyspace = os.Getenv("SCYLLA_KEYSPACE")
	env.Scyllamigration = os.Getenv("SCYLLA_MIGRATIONS")

	return env, err
}
