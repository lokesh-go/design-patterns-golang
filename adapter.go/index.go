package adapter

// AdapterExample ...
func AdapterExample() {
	stringData := "{\"app\": \"test-app\", \"host\": \"localhost\", \"ip\": \"172.31.100.12\"}"
	dataAnalyticsTool(jsonAdapter(stringData))
}
