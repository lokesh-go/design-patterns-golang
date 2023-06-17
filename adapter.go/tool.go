package adapter

import (
	"log"
)

type jsonData struct {
	App  string
	Host string
	IP   string
}

func dataAnalyticsTool(data *jsonData) {
	log.Println(data)
}
