package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func SetUpKafka() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	var err error
	Producer, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating kafka producer: %v\n", err)
	}

	CreateTopics(config)

	consumerConfig := sarama.NewConfig()
    consumerConfig.Consumer.Return.Errors = true

    // Create a new consumer
    Consumer, err = sarama.NewConsumer([]string{"localhost:9092"}, consumerConfig)

	if err != nil {
		log.Fatalf("Error creating consumer: %v\n", err)
	}

	InventoryPartition, err = Consumer.ConsumePartition(GET_INVENTORY, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error consuming inventory partition: %v\n", err)
	}

	OrderPartition, err = Consumer.ConsumePartition(GET_ORDER_STATUS, 0, sarama.OffsetNewest)

	if err != nil {
		log.Fatalf("Error consuming order partition: %v\n", err)
	}
}

func CreateTopics(config *sarama.Config) {
	config.Version = sarama.V3_6_0_0

	admin, err := sarama.NewClusterAdmin([]string{"localhost:9092"}, config)

	if err != nil {
		log.Fatalf("Error creating kafka cluster admin: %v\n", err)
	}

	topics := []string{GET_INVENTORY, GET_ORDER_STATUS, SET_INVENTORY, CREATE_ORDER}

	for _, topic := range topics {
		exists, err := topicExists(admin, topic)
		if err != nil {
			log.Fatalf("Error checking if kafka topic exists: %v\n", err)
		}

		if !exists {
			topicDetail := &sarama.TopicDetail{
				NumPartitions:     1,
				ReplicationFactor: 1,
				ConfigEntries:     map[string]*string{}, // can include topic-specific configurations
			}
			err = admin.CreateTopic(topic, topicDetail, false)
			if err != nil {
				log.Fatalln("Failed to create topic:", err)
			}
		}
	}
}

func topicExists(admin sarama.ClusterAdmin, topic string) (bool, error) {
    topics, err := admin.ListTopics()

    if err != nil {
        return false, err
    }

    _, exists := topics[topic]
    return exists, nil
}