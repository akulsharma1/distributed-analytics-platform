package kafka

import "github.com/IBM/sarama"

var (
	Producer sarama.SyncProducer
	Consumer sarama.Consumer
)

const (
	GET_INVENTORY string = "get_inventory"
	INVENTORY string = "inventory"
	ORDER string = "order"
)

var (
	InventoryPartition sarama.PartitionConsumer
)