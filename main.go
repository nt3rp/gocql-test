package main

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("127.0.0.1", "127.0.0.2")
	cluster.Consistency = gocql.All
	cluster.Keyspace = "experiment"
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 6}
	cluster.Timeout = 1200 * time.Millisecond

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	session.SetPageSize(10000)
}

func main() {
	keys1 := []string{"simple", "complex"}
	keys2 := []string{"all", "some", "undefined"}
	keys3 := []string{"complete", "partial"}
	for _, key1 := range keys1 {
		for _, key2 := range keys2 {
			for _, key3 := range keys3 {
				for i := 0; i < 1000; i += 1 {
					gocql_case(key1, key2, key3)
				}
			}
		}
	}

	fmt.Println("Now, run `nodetool cfstats experiment`")
	fmt.Println("'Average tombstones' per slice should be 0.")
}
