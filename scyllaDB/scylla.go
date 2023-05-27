package scyllaDB

import (
	"github.com/amin-mir/reporting/config"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type Manager struct {
	cfg config.Config
}

func createManager(cf config.Config) *Manager {
	return &Manager{
		cfg: cf,
	}
}

func connect(keyspace string, host []string) (gocqlx.Session, error) {
	c := gocql.NewCluster(host...)
	c.Keyspace = keyspace
	return gocqlx.WrapSession(c.CreateSession())
}

func (m *Manager) Connect() (gocqlx.Session, error) {
	return connect(m.cfg.ScyllaKeySpace, m.cfg.ScyllaHost)
}

func (m *Manager) CreateKeySpace(keyspace string) error {
	sessions, err := connect("system", m.cfg.ScyllaHost)
	if err != nil {
		return err
	}
	defer sessions.Close()

	query := `CREATE KEYSPACE IF NOT EXISTS reporting WITH replication = {'class': 'SimpleStrategy', 'replicaation_factor':2}`
	return sessions.ExecStmt(query)

}
