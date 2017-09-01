package zookeeper

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func CrateConnection(zkUrl string, sessionTimeOut time.Duration, isZkSecurityEnabled bool) (conn *zk.Conn, err error) {
	// TODO validation zookeeper url

	urls := []string{zkUrl}

	conn, _, err = zk.Connect(urls, sessionTimeOut)
	if err != nil {
		// TODO logger
		return
	}

	return
}
