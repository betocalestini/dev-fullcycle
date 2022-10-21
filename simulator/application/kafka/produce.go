package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	route2 "github.com/betocalestini/dev-fullcycle-simulator/application/route"
	"github.com/betocalestini/dev-fullcycle-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// Produce is responsible to publish the positions of each request
// kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal
// kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
// Example of a json request:
// {"clientId":"1","routeId":"1"}
// {"clientId":"2","routeId":"2"}
// {"clientId":"3","routeId":"3"}
func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
