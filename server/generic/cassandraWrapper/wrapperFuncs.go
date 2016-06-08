package cassandraWrapper

import (
	"github.com/gocql/gocql"
)

var cluster *gocql.ClusterConfig

func init() {
	cluster = gocql.NewCluster("localhost")
	cluster.Keyspace = "datasheet"
	cluster.ProtoVersion = 4
	cluster.Consistency = gocql.Quorum
}

func CreateSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
