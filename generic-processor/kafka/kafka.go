package kafka

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

const VRF_KAFKA_TOPIC = "jalapeno.telemetry.vrf"
const LLDP_KAFKA_TOPIC = "jalapeno.telemetry.lldp"

func NewSaramaConsumer(brokerAddresses []string) sarama.Consumer {
	sarama.Logger = log.New(os.Stdout, "[sarama]", log.LstdFlags)
	consumer, err := sarama.NewConsumer(brokerAddresses, sarama.NewConfig())
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	return consumer
}

func NewPartitionConsumer(consumer sarama.Consumer, topic string) sarama.PartitionConsumer {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating Sarama PartitionConsumer: %v", err)
	}
	return partitionConsumer
}

func CloseConsumer(consumer sarama.PartitionConsumer) {
	log.Printf("Closing sarama partition consumer\n")
	err := consumer.Close()
	if err != nil {
		log.Printf("Error closing sarama paritition consumer: %v", err)
	}
}
