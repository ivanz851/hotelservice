package main

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
}

func sendMail(subject string, body string, to string) {
	auth := smtp.PlainAuth(
		"",
		"ilya.sokurwork@gmail.com",
		"mlionkvkzfkpgjbu",
		"smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"ilya.sokurwork@gmail.com",
		[]string{to},
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {

	reader := getKafkaReader("kafka:9092", "bookings", "booking-group")

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		fmt.Println(string(m.Value))
		to := strings.TrimSpace(string(m.Value))
		cleanValue := strings.ReplaceAll(to, "\n", "")
		cleanValue = strings.ReplaceAll(cleanValue, "\r", "")
		sendMail("LASTTEST", "lsattest", cleanValue)
	}
}
