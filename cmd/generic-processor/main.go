package main

import (
	"log"
	"os"

	"github.com/cisco-open/jalapeno/generic-processor/arango"
	"github.com/cisco-open/jalapeno/generic-processor/kafka"
	"github.com/cisco-open/jalapeno/generic-processor/processor"
)

func main() {
	// TODO: read from env variable
	kafkaConnection := []string{""}

	// TODO: create map from env variables
	topicToCollection := map[string]string{
		"jalapeno.telemetry.vrf":  "telemetry_vrf",
		"jalapeno.telemetry.lldp": "telemetry_lldp",
	}

	// create arango config
	arangoConfig := arango.ArangoConfig{
		URL:      os.Getenv("ARANGO_HOST"), // golang http error when port is specified "http: server gave HTTP response to HTTPS client"
		User:     os.Getenv("ARANGO_USER"),
		Password: os.Getenv("ARANGO_PASSWORD"),
		Database: os.Getenv("ARANGO_DB"),
	}

	// create arango connection
	arangoConn, err := arango.NewArangoConnection(arangoConfig)
	if err != nil {
		log.Fatalf("Failed to create ArangoConnection: %v", err)
	}

	// initialize Kafka Consumer
	kafkaConsumer := kafka.NewSaramaConsumer(kafkaConnection)

	// create processors
	// one processor holds the reference to consume from a topic and a reference to write to a collection
	var processors []processor.Processor
	for topic, collection := range topicToCollection {
		processors = append(processors, processor.NewProcessor(topic, collection, arangoConn, kafkaConsumer))
	}

	// start processing
	for _, processor := range processors {
		processor.StartProcessing()
	}

}
