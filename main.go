/*
NSQ demo za golang prezentaciju
*/
package main

import (
	"flag"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/minus5/svckit/log"
	nsq "github.com/minus5/svckit/nsq"
	"github.com/minus5/svckit/signal"
)

var (
	nsqtopic = "super.mario"
	listen   string
	prezDir  string
	upgrader = websocket.Upgrader{} // use default options
	producer *nsq.Producer
	consumer *nsq.Consumer
	ws       *websocket.Conn
)

func init() {
	flag.StringVar(&listen, "listen", "127.0.0.1:8080", "adresa i port na kome slusam")
	flag.StringVar(&prezDir, "prezDir", "./prez", "folder s prezentacijom")
}

// main je glavna ulazna tocka u program
func main() {
	log.Info("starting")
	defer log.Info("stopped")

	producer = nsq.MustNewProducer(nsqtopic)
	defer producer.Close()
	consumer = nsq.MustNewConsumer(nsqtopic, onNSQMsg)
	defer consumer.Close()

	// Default static web server
	http.Handle("/", http.FileServer(http.Dir(prezDir)))
	http.HandleFunc("/ws", serveWs)
	go http.ListenAndServe(listen, nil)

	// Ovdje cekamo izlaz iz aplikacije
	var wg sync.WaitGroup
	signal.WaitForInterupt()
	wg.Wait()
}

func onNSQMsg(msg *nsq.Message) error {
	log.S("poruka", string(msg.Body)).Info("NSQ Receive")
	if nil == ws {
		return nil
	}
	if err := ws.WriteMessage(websocket.TextMessage, msg.Body); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	var err error
	if ws, err = upgrader.Upgrade(w, r, nil); err != nil {
		log.Error(err)
		return
	}

	defer func() {
		ws.Close()
		ws = nil
	}()

	// Read and echo messsages
	for {
		// Read message
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Error(err)
			break
		}

		// Send NSQ
		log.S("poruka", string(msg)).Info("NSQ Send")
		producer.Publish(msg)
	}
}
