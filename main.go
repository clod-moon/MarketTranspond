package main

import (

	"flag"
	"log"
	"os"
	"os/signal"
	"testWebsocket/websocket"
	"sync"
	//"golang.org/x/net/websocket"
)



var wg sync.WaitGroup


func main() {

	flag.Parse()

	log.SetFlags(0)



	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt)


	wg.Add(1)
	go websocket.GetMarketHandler()

	wg.Add(1) //
	go websocket.GetHuobiMarket(wg)  //无需本地IP地址，直接运行

	wg.Add(1)
	go websocket.GetOKExMarket(wg)

	wg.Add(1)
	go websocket.GetBianMarket(wg)


	wg.Wait()
}