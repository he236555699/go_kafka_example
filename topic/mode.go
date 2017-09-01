package topic

// CreateMode value determines how the zookeeper node is created on ZooKeeper.
type Mode struct {
	Flag       int32
	Ephemeral  bool
	Sequential bool
}

const (
	PERSISTENT            = 0
	EPHEMERAL             = 1
	PERSISTENT_SEQUENTIAL = 2
	EPHEMERAL_SEQUENTIAL  = 3
)

var createMode = map[int]Mode{
	/**
	 * The zookeeper node will not be automatically deleted upon client's disconnect.
	 */
	PERSISTENT: Mode{0, false, false},

	/**
	 * The zookeeper node will be deleted upon the client's disconnect.
	 */
	EPHEMERAL: Mode{1, true, false},
	/**
	 * The zookeeper node will not be automatically deleted upon client's disconnect,
	 * and its name will be appended with a monotonically increasing number.
	 */
	PERSISTENT_SEQUENTIAL: Mode{2, false, false},

	/**
	 * The zookeeper will be deleted upon the client's disconnect, and its name
	 * will be appended with a monotonically increasing number.
	 */
	EPHEMERAL_SEQUENTIAL: Mode{3, true, true},
}

func CreateMode(mode int) Mode {
	bean, ok := createMode[mode]
	if !ok {
		// TODO log
	}

	return bean
}
