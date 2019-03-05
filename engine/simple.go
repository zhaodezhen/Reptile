package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		ParserResult,err := worker(r)
		if err !=nil {
			continue
		}
		requests = append(requests, ParserResult.Requests...)
		for _, item := range ParserResult.Items {
			log.Printf("Got item %s ", item)
		}
	}
}
func worker(r Request) (ParserResult, error) {
	log.Printf("Fething: %s ", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher error"+"fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}
