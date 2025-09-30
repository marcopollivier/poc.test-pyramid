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
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	
	return &KafkaPublisher{producer: producer}, nil
}

func (k *KafkaPublisher) PublishDiscount(discount *model.Discount) error {
	data, err := json.Marshal(discount)
	if err != nil {
		return err
	}
	
	msg := &sarama.ProducerMessage{
		Topic: "discount-calculated",
		Value: sarama.StringEncoder(data),
	}
	
	partition, offset, err := k.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	
	log.Printf("Message sent to partition %d at offset %d", partition, offset)
	return nil
}

func (k *KafkaPublisher) Close() error {
	return k.producer.Close()
}
