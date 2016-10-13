package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"log"
)

var bearer = ""
var device = ""
var port = ""

type Health struct {
	Name          string `json:"name"`
	SmokeState    string `json:"smoke_alarm_state"`
	COState       string `json:"co_alarm_state"`
	BatteryHealth string `json:"battery_health"`
	UIState       string `json:"ui_color_state"`
	LastConn      string `json:"last_connection"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://developer-api.nest.com/devices/smoke_co_alarms/%s", device), nil)
	if err != nil {
		log.Fatalf("[NewRequest]: %v", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearer))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("[client.Do]: %v", err)
	}

	var health Health
	// Decode the incoming JSON
	err = json.NewDecoder(resp.Body).Decode(&health)
	if err != nil {
		log.Fatalf("[client.Do]: %v", err)
		panic(err)
	}

	message := fmt.Sprintf(
		"The Nest in your %s reports %s for smoke, and %s for Carbon Monoxide. The battery reports %s. Overall the status is %s. The last update was at %s",
		health.Name,
		health.SmokeState,
		health.COState,
		health.BatteryHealth,
		health.UIState,
		health.LastConn,
	)
	fmt.Fprintf(w, message)
}

func main() {
	port = os.Getenv("PORT")
	device = os.Getenv("DEVICE")
	bearer = os.Getenv("BEARER")
	if port == "" || device == "" || bearer == "" {
		panic("System incorrectly configured. Ensure you have set all three environment variables: PORT, DEVICE and BEARER")
	}
	http.HandleFunc("/", handler)
	fmt.Println("Starting on port", port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		panic(err)
	}

}
