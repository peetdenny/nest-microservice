package main
import (
	"fmt";
	"net/http";
	"os/exec";
	"os";
	"encoding/json";
)

	var bearer =""
	var device =""
	var port =""


func handler(w http.ResponseWriter, r *http.Request){
	// OK, so I'm using curl for this because I couldn't get the OAuth Bearer header to work with Golang directly, and I only had an hour, don't you judge me!
	line := "curl -v -L -H \"Authorization: Bearer"+bearer+" \" -X GET https://developer-api.nest.com/devices/smoke_co_alarms/"+device
	 out, err := exec.Command("bash", "-c", line).Output()
     	 if err != nil {
        	panic(err)
    	 }
	var m map[string]string
	json.Unmarshal(out, &m)
	message := fmt.Sprintf("The Nest in your %s reports %s for smoke, and %s for Carbon Monoxide. The battery reports %s. Overall the status is %s", 
		m["name"],
		m["smoke_alarm_state"],
		m["co_alarm_state"],
		m["battery_health"],
		m["ui_color_state"],
	)
	fmt.Fprintf(w, message)
}


func main(){
	port = os.Getenv("PORT")
	device = os.Getenv("DEVICE")
	bearer  = os.Getenv("BEARER")
	if (port=="" || device =="" || bearer==""){
		panic("System incorrectly configured. Ensure you have set all three environment variables: PORT, DEVICE and BEARER")
	}
	http.HandleFunc("/", handler)
	fmt.Println("Starting on port",port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}

}
