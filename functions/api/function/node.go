package function

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/my-ermes-labs/api-go/api"
	"github.com/my-ermes-labs/api-go/infrastructure"
	log "github.com/my-ermes-labs/log"
	rc "github.com/my-ermes-labs/storage-redis/packages/go"
	"github.com/redis/go-redis/v9"
)

// The node that the function is running on.
var node *api.Node

// The Redis client.
var redisClient *redis.Client

func init() {

	log.MyLog("-------- Initializing API --------")

	// Get the node from the environment variable.
	encodedJsonNode := envOrPanic("ERMES_NODE")
	decodedJsonNode, err := base64.StdEncoding.DecodeString(encodedJsonNode)
	if err != nil {
		log.MyLog("Error30")
		fmt.Println("Errore nella decodifica:", err)
		return
	}

	// Unmarshal the environment variable to get the node.
	infraNode, err := infrastructure.UnmarshalNode([]byte(decodedJsonNode))
	// Check if there was an error unmarshalling the node.
	if err != nil {
		log.MyLog("Error38")
		panic(err)
	}
	// Get the Redis connection details from the environment variables.
	redisHost := envOrDefault("REDIS_HOST", "10.62.0.1")
	redisPort := envOrDefault("REDIS_PORT", "6379")
	redisPassword := envOrDefault("REDIS_PASSWORD", "")

	// Create a new Redis client.
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	if err := checkRedisConnection(redisClient); err != nil {
		log.MyLog(fmt.Sprintf("Error in Redis Connection: %v", err))
	} else {
		log.MyLog("Redis Connection Done!")
	}

	// The Redis commands.
	var RedisCommands = rc.NewRedisCommands(redisClient)

	// Create a new node with the Redis commands.
	node = api.NewNode(*infraNode, RedisCommands)

	log.MyNodeLog(node.AreaName, "init() for API completed!\n")
	log.MyNodeLog(node.AreaName, "-------------------------\n")
}

// Get the value of an environment variable or return a default value.
func envOrDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

// Get the value of an environment variable or panic if it is not set.
func envOrPanic(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(key + " env variable is not set")
	}
	return value
}

func checkRedisConnection(client *redis.Client) error {
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	return err
}
