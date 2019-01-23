package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Shopify/sarama"
	"github.com/farukterzioglu/micGo-services/Review.API/Api"
	"github.com/gorilla/mux"
)

var (
	kafkaBrokers = flag.String("kafka_brokers", "localhost:9092", "The kafka broker address in the format of host:port")
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

	router := initRouter()
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

func initRouter() (router *mux.Router) {
	router = mux.NewRouter()

	v1 := router.PathPrefix("/v1").Subrouter()

	// TODO : Pass kafka publisher
	reviewRoutes := api.NewReviewRoutes()
	reviewRoutes.RegisterReviewRoutes(v1, "/review")
	return
}
