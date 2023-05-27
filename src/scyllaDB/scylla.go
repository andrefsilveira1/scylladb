package scyllaDB

import (
	"src/config"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type Manager struct {
	cfg config.Config
}

func CreateManager(cfg config.Config) *Manager {
	return &Manager{
		cfg: cfg,
	}
}

func (m *Manager) connect(keyspace string, host []string) (gocqlx.Session, error) {
	c := gocql.NewCluster(host...)
	c.Keyspace = keyspace
	return gocqlx.WrapSession(c.CreateSession())
}

func (m *Manager) Connect() (gocqlx.Session, error) {
	return m.connect(m.cfg.ScyllaKeySpace, m.cfg.ScyllaHost)
}

func (m *Manager) CreateKeySpace(keyspace string) error {
	sessions, err := m.connect("system", m.cfg.ScyllaHost)
	if err != nil {
		return err
	}
	defer sessions.Close()

	query := `CREATE KEYSPACE IF NOT EXISTS reporting WITH replication = {'class': 'SimpleStrategy', 'replication_factor':1}`
	return sessions.ExecStmt(query)

}
