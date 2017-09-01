package main

import (
	"time"

	"github.com/he236555699/go_kafka_example/topic"
	"github.com/he236555699/go_kafka_example/zookeeper"
	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	topicInfo := topic.TopicInformation{}
	topicInfo.TopicName = "/test"
	topicInfo.TopicData = []byte("test message")
	topicInfo.Partitions = 2
	topicInfo.Replication = 2

	conn, err := zookeeper.CrateConnection("localhost:2181", time.Second*5, true)
	if err != nil {
		println(err.Error())
		return
	}

	data, stat, err := topicInfo.Describe(conn)
	if err != nil {
		println(err.Error())
		return
	}

	println(string(data[:]))
	println(stat.Ctime)

	acl := zk.WorldACL(zk.PermAll)

	flag := topic.CreateMode(topic.PERSISTENT)
	println(flag.Flag)

	path, err := topicInfo.Create(conn, flag, acl)
	if err != nil {
		println(err.Error())
		return
	}

	println(path)

}
