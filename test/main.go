package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
)

func main() {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	q := elastic.NewMatchQuery("address", "street")
	result, err := client.Search().Index("user").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	total := result.Hits.TotalHits.Value
	fmt.Println(total)
	//result.Hits.Hits
}
