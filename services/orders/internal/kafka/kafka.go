package kafka

import "github.com/IBM/sarama"

var (
	Producer sarama.SyncProducer
	Consumer sarama.Consumer
)

const (
	GET_INVENTORY string = "get_inventory"
	SET_INVENTORY string = "inventory"

	CREATE_ORDER string = "create_order"
	GET_ORDER_STATUS string = "order_status"
)

var (
	InventoryPartition sarama.PartitionConsumer
	OrderPartition sarama.PartitionConsumer
)