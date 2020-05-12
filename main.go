package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"os/user"
	"strconv"
	"syscall"
	"time"

	hue "github.com/collinux/gohue"
)

type config struct {
	Api string `json:"ip"`
	Key string `json:"key"`
}

var c config

// First parameter is a lamp ID
func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("No Params found")
		return
	}

	lampID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("no ID set")
		return
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	jsonFile, err := os.Open(usr.HomeDir + "/.hue/config_go")
	data, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	bridge, _ := hue.NewBridge(c.Api)
	bridge.Login(c.Key)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGUSR1)

	go toggle(bridge, sigs, lampID)

	for true {
		light, _ := bridge.GetLightByIndex(lampID)

		status := "OFF"
		if light.State.On {
			status = "ON"
		}
		fmt.Println(status)

		time.Sleep(2 * time.Second)
	}

}

func toggle(bridge *hue.Bridge, sigs chan os.Signal, lampID int) {
	<-sigs // if SIGUSR1 was fired
	light, _ := bridge.GetLightByIndex(lampID)
	light.Toggle()
	toggle(bridge, sigs, lampID)
}
