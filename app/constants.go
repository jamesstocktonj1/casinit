package app

import "github.com/spf13/viper"

const (
	// CassandraHosts is the environment variable name for the Cassandra hosts
	CassandraHost = "CASSANDRA_HOST"

	// CassandraUser is the environment variable name for the Cassandra user
	CassandraUser = "CASSANDRA_USER"

	// CassandraPass is the environment variable name for the Cassandra password
	CassandraPass = "CASSANDRA_PASS"

	// CassandraTimeout is the environment variable name for the Cassandra timeout
	CassandraTimeout = "STARTUP_TIMEOUT"
)

func BindEnv(cfg *viper.Viper) {
	cfg.SetDefault(CassandraHost, "localhost")
	cfg.SetDefault(CassandraTimeout, 60)
	cfg.AutomaticEnv()
}
