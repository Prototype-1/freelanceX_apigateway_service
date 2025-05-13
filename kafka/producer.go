package kafka

import (
    "github.com/Shopify/sarama"
    "log"
)

var Producer sarama.SyncProducer

func InitKafkaProducer(brokers []string) error {
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        return err
    }

    Producer = producer
    return nil
}

func SendMessage(topic string, key string, value []byte) error {
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Key:   sarama.StringEncoder(key),
        Value: sarama.ByteEncoder(value),
    }

    _, _, err := Producer.SendMessage(msg)
    if err != nil {
        log.Printf("Kafka send error: %v", err)
    }

    return err
}
