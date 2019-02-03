// Review API
//
// Web API for review project
//
//     Schemes: http
//     BasePath: /v1
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Faruk Terzioğlu <faruk.terzioglu@hotmail.com>
//     Host:
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Shopify/sarama"
	"github.com/farukterzioglu/micGo-services/Review.API/api"
	_ "github.com/farukterzioglu/micGo-services/Review.API/swagger" // Required for Swagger to explore models
	pb "github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/reviewservice"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	kafkaBrokers = flag.String("kafka_brokers", "localhost:9092", "The kafka broker address in the format of host:port")
	serverAddr   = flag.String("server_addr", "localhost:10000", "The rpc server address in the format of host:port")
)

func main() {
	flag.Parse()
	fmt.Printf("Broker address : %s\n", *kafkaBrokers)

	// Init Kafka producer
	// TODO : Retry + fail over
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error while creating producer : ", err.Error())
		os.Exit(1)
	}

	// TODO : Retry + fail over
	rpcServer, conn, err := initRPCServer()
	if err != nil {
		fmt.Println("Error while dialing rpc server : ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	router := initRouter(producer, rpcServer)

	// Host Swagger UI
	fs := http.FileServer(http.Dir("./swaggerui/"))
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))

	// Start to listen
	log.Fatal(http.ListenAndServe(":8000", router))
}

func initProducer() (*sarama.SyncProducer, error) {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	config := sarama.NewConfig()
	config.ClientID = "Review.API"
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{*kafkaBrokers}, config)
	return &producer, err
}

func initRPCServer() (*pb.ReviewServiceClient, *grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		return nil, nil, err
	}

	rpcConnState := conn.GetState()
	fmt.Printf("Rpc cpnnection state : %s\n", rpcConnState)

	client := pb.NewReviewServiceClient(conn)
	return &client, conn, nil
}

func initRouter(producer *sarama.SyncProducer, client *pb.ReviewServiceClient) (router *mux.Router) {
	router = mux.NewRouter()
	// TODO : Causes Swagger UI to be parsed as JSON instead of html
	// router.Use(commonMiddleware)

	v1 := router.PathPrefix("/v1").Subrouter()

	reviewRoutes := api.NewReviewRoutes(producer)
	reviewRoutes.RegisterReviewRoutes(v1, "/review")

	queryController := api.NewQueryController(client)
	queryController.RegisterRoutes(v1, "/review")

	return
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
