package processor

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/arangodb/go-driver"
	"github.com/cisco-open/jalapeno/generic-processor/arango"
	"github.com/cisco-open/jalapeno/generic-processor/kafka"
)

type Processor struct {
	TopicConsumer    sarama.PartitionConsumer //consumer of a specific topic
	ArangoCollection driver.Collection        // specific collection in arango
}

func NewProcessor(topic string, collectionName string, arangoConn *arango.ArangoConn, kafkaConsumer sarama.Consumer) Processor {
	return Processor{
		TopicConsumer:    kafka.NewPartitionConsumer(kafkaConsumer, topic),
		ArangoCollection: arangoConn.GetCollection(collectionName),
	}
}

func (p *Processor) ProcessOne() {
	go func() {

		defer func() {
			kafka.CloseConsumer(p.TopicConsumer)
		}()

		msg := <-p.TopicConsumer.Messages()
		// TODO: create document from message and add it to collection
		fmt.Println(string(msg.Value))
	}()
}

func (p *Processor) StartProcessing() {
	go func() {
		defer func() {
			kafka.CloseConsumer(p.TopicConsumer)
		}()

		for {
			msg := <-p.TopicConsumer.Messages()
			// TODO: create document from message and add it to collection

			// create hash of message -> check if in arango, if yes -> ignore - nothing changed
			// if hash not in arango - need to updae old document or update - how to figure out which document?
			fmt.Println(string(msg.Value))
		}
	}()
}
