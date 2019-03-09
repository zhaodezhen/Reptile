package persisit

import (
	"log"
	"golang.org/x/net/context"
	"src/gopkg.in/olivere/elastic.v6"
)

func ItemSave() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			id, err := save(item)
			if err != nil {
				log.Printf("Item saver err saving item %v: %v ", item, err)
			}
			log.Printf("Item Saver: got ID %#d", id)

		}
	}()
	return out
}
func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
