package urlbase

import (
	"fmt"
	"strconv"
	
	"github.com/garyburd/redigo/redis"
)

const (
	Version = 0.01
)

var (	
	redisServ redis.Conn
	err error
	keyform = "golib:url:"
)

// Init initialize the Redis DB connection
// Returns error=false for success and error=true for failure
func Initialize(server string) bool {
	if redisServ == nil {
		redisServ, err = redis.Dial("tcp", server)
		if err != nil {
			fmt.Println("Failed to connect to Redis DB.")
			return true
		}
		return false
	} else {
		fmt.Println("Redis DB already connected.")
		return true
	}
}

// Finalize closes Redis DB connection
// Returns true on success and false on failure
func Finalize() bool {
	if redisConnected() {
		redisServ.Close()
		return true
	} else {
		fmt.Println("Redis DB already disconnected.")
		return false
	}		
}

// Store stores "url" into the Redis DB
// against key "(keyform + id)"
// Returns true for success and false for failure
func Store(id int, url string) bool {
	if redisConnected() {
		key := keyform + strconv.Itoa(id)
		redisServ.Do("SET", key, url)
		return true
	} 
	return false
}

// Fetch gets "url" from the Redis DB
// against key "(keyform + id)"
// Returns true, value for success and false, nil for failure
func Fetch(id int) (value string, error bool){
	if redisConnected() {	
		key := keyform + strconv.Itoa(id)
		value, err := redis.String(redisServ.Do("GET", key))
		if err != nil {
			fmt.Printf("Failed to GET key ", key)
			return "", true
		}
		return value, false
	}
	return "", true
}

// redisConnected helper function to check
// if Redis DB is connected or not
func redisConnected() bool {
	if redisServ != nil {
		return true
	} else {
		fmt.Println("Redis DB connection is not open.")
		return false
	}
}
	
