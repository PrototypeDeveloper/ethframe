package main

import (
	"encoding/hex"
	"ethframe"
	"log"
	"time"
)

func main() {

	driver := ethframe.NewDriver()
	if err := driver.Open("eth0"); err != nil {
		log.Fatal(err)
	}
	defer driver.Close()

	packet := "sample"
	frame, _ := hex.DecodeString(packet)
	driver.Send(frame)

	for {
		f, err := driver.Receive()
		if err != nil {
			log.Println("recv err:", err)
			time.Sleep(time.Second)
			continue
		}
		log.Println("recv:", hex.EncodeToString(f))
	}
}
