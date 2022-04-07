package database

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func GetESClient() {

	cfg := elasticsearch.Config{
		Username: "elastic",
		Password: "dqQt9gdW9nUArPqbD5VN",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	log.Println(elasticsearch.Version)

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
}
