package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Shopify/sarama"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Commands/V1"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Models"
	"github.com/gorilla/mux"
)

var (
	kafkaBrokers = flag.String("kafka_brokers", "172.24.96.1:9092", "The kafka broker address in the format of host:port")
)

var producer sarama.SyncProducer

func main() {
	flag.Parse()
	fmt.Printf("Broker address : %s\n", *kafkaBrokers)

	var err error
	producer, err = initProducer()
	if err != nil {
		fmt.Println("Error while creating producer : ", err.Error())
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.HandleFunc("/review", createReview).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func initProducer() (producer sarama.SyncProducer, err error) {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	config := sarama.NewConfig()
	config.ClientID = "Review.API"
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err = sarama.NewSyncProducer([]string{*kafkaBrokers}, config)

	return
}

func createReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	_ = json.NewDecoder(r.Body).Decode(&review)

	command := &commands.CreateReviewCommand{
		Review: review,
	}

	command.Review.Status = models.Created

	msg, err := json.Marshal(command)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	publish(string(msg), "create-review")
	w.WriteHeader(http.StatusCreated)
}

func publish(message, topicName string) {
	msg := &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.StringEncoder(message),
	}

	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	fmt.Printf("Delivered %s[part:%d] (@%d) - %s\n'", topicName, p, o, message)
}
