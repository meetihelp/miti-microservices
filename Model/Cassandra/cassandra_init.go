package Cassandra

import(
	"github.com/gocql/gocql"
)

var session *gocql.Session

func init(){
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "miti"
	cluster.Consistency = gocql.Any
	session, _ = cluster.CreateSession()
}

func GetDB() *gocql.Session{
	return session
}