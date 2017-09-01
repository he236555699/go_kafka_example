package topic

import (
	"github.com/samuel/go-zookeeper/zk"
)

type Topic interface {
	Create(name string, data []byte, crateMode Mode, aclList []zk.ACL) error
	List() error
	Update() error
	Delete() error
	Describe() error
	Exists(conn *zk.Conn) error
	Alter() error // Alter the configuration for the topic
}

type TopicInformation struct {
	TopicName   string
	TopicData   []byte
	Partitions  int
	Replication int
}

func (t *TopicInformation) Create(conn *zk.Conn, crateMode Mode, aclList []zk.ACL) (path string, err error) {
	defer conn.Close()

	path, err = conn.Create(t.TopicName, t.TopicData, crateMode.Flag, aclList)
	if err != nil {
		// TODO logger
		return
	}

	return
}

func (t *TopicInformation) Describe(conn *zk.Conn) (data []byte, stat *zk.Stat, err error) {
	data, stat, err = conn.Get(t.TopicName)
	if err != nil {
		// TODO logger
		return
	}

	return
}

func (t *TopicInformation) Exists(conn *zk.Conn) (exists bool, stat *zk.Stat, err error) {
	exists, stat, err = conn.Exists(t.TopicName)
	if err != nil {
		return
	}

	return
}
