package main

import (
	"crawler/engine"
	"crawler/persisit"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	var e = engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 500,
		ItemChan:    persisit.ItemSave(),

	}
	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun/beijing",
	//	ParserFunc:parser.ParseCity,
	//})
}
