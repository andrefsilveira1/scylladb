package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadEnv(t *testing.T) {
	require := require.New(t)
	os.Setenv("SCYLLA_HOST", "host1.com,host2.com")
	os.Setenv("SCYLLA_MIGRATIONS", "./cql")
	os.Setenv("SCYLLA_KEYSPACE", "reporting")

	actual, err := LoadEnv()
	require.NoError(err)

	expected := Config{
		ScyllaHost:       []string{"host1.com", "host2.com"},
		ScyllaMigrations: "./cql",
		ScyllaKeySpace:   "reporting",
	}

	require.Equal(expected, actual)
}
