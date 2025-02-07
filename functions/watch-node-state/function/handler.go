package function

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	log "github.com/my-ermes-labs/log"
)

var ctx = context.Background()
var knownKeys = make(map[string]bool)

func checkRedis() {
	for {
		log.MyNodeLog(node.AreaName, "NEW Check")
		keys, err := redisClient.Keys(ctx, "*").Result()
		if err != nil {
			fmt.Println("Error in retrieving Redis Key: ", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, key := range keys {
			if !knownKeys[key] {
				value, err := redisClient.Get(ctx, key).Result()
				if err != nil {
					fmt.Println("Error in retrieving value: ", err)
					continue
				}

				log.MyNodeLog(node.AreaName, "New Session! "+key+": "+value+"\n")
				// invoke migration

				sendMigrationRequest("192.168.64.28")

				knownKeys[key] = true
			}
		}

		time.Sleep(5 * time.Second)
	}

}

func Handle(w http.ResponseWriter, r *http.Request) {
	log.MyNodeLog(node.AreaName, "START WATCHER")
	go checkRedis()
}

func sendMigrationRequest(url string) {

	// prendo le subareas

	// per ogni subareas mando il messaggio
	completeURL := "http://" + url + ":8080/function/migrate"

	log.MyNodeLog(node.AreaName, "Is sending a migration request to "+completeURL)
	// empty data
	data := []byte(`{}`)

	req, err := http.NewRequest("POST", completeURL, bytes.NewBuffer(data))
	if err != nil {
		log.MyNodeLog(node.AreaName, fmt.Sprintf("Error creating request: %s", err))
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.MyNodeLog(node.AreaName, fmt.Sprintf("Error sending request to watcher function: %s", err))
		return
	}
	defer resp.Body.Close()

	log.MyNodeLog(node.AreaName, fmt.Sprintf("Watcher function responded with status code: %d", resp.StatusCode))

}
