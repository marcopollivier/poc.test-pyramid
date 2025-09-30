package messaging

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/marcopollivier/poc.test-pyramid/model"
)

type KafkaPublisher struct {
	producer sarama.SyncProducer
}

func NewKafkaPublisher(brokers []string) (*KafkaPublisher, error) {
	log.Printf("KafkaPublisher: Initializing with brokers: %v", brokers)
	
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Printf("KafkaPublisher: Failed to create producer - %v", err)
		return nil, err
	}
	
	log.Println("KafkaPublisher: Successfully connected to Kafka")
	return &KafkaPublisher{producer: producer}, nil
}

func (k *KafkaPublisher) PublishDiscount(discount *model.Discount) error {
	log.Printf("KafkaPublisher: Publishing discount ID: %d to topic 'discount-calculated'", discount.ID)
	
	data, err := json.Marshal(discount)
	if err != nil {
		log.Printf("KafkaPublisher: JSON marshal error - %v", err)
		return err
	}
	
	log.Printf("KafkaPublisher: Message payload: %s", string(data))
	
	msg := &sarama.ProducerMessage{
		Topic: "discount-calculated",
		Value: sarama.StringEncoder(data),
	}
	
	partition, offset, err := k.producer.SendMessage(msg)
	if err != nil {
		log.Printf("KafkaPublisher: Send message error - %v", err)
		return err
	}
	
	log.Printf("KafkaPublisher: Message sent successfully - Partition: %d, Offset: %d", partition, offset)
	return nil
}

func (k *KafkaPublisher) Close() error {
	log.Println("KafkaPublisher: Closing producer")
	return k.producer.Close()
}
