package entity

type Shelf struct {
	Games []Game`json:"games"`
	Name string `json:"name"`
	LEDOn bool `json:"led_on"`
}
