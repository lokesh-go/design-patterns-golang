package adapter

import (
	"encoding/json"
	"log"
)

func jsonAdapter(data interface{}) (res *jsonData) {
	switch data.(type) {
	case string:
		{
			if data == nil {
				return nil
			}
			err := json.Unmarshal([]byte(data.(string)), &res)
			if err != nil {
				log.Fatal(err)
			}
			return res
		}
	}
	return nil
}
