package kafka

import "github.com/IBM/sarama"

func SetUpKafka() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	var err error
	Producer, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	consumerConfig := sarama.NewConfig()
    consumerConfig.Consumer.Return.Errors = true

    // Create a new consumer
    Consumer, err = sarama.NewConsumer([]string{"localhost:9092"}, config)

	if err != nil {
		panic(err)
	}

	InventoryPartition, err = Consumer.ConsumePartition(INVENTORY, 0, sarama.OffsetNewest)

	if err != nil {
		panic(err)
	}
}
