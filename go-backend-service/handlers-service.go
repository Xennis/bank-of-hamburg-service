package main

import (
	"net/http"
	"fmt"
	"github.com/go-redis/redis"
)


func LivenessProbe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ReadinessProbe(w http.ResponseWriter, r *http.Request) {
    ok := true
    errMsg := ""

    // Check database
    if err := len(accounts) < 1; err {
        ok = false
        errMsg += "Database not ok."
	}

	// Check redis
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-master:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		ok = false
		errMsg += "Redis not ok"
	}

    if ok {
        w.WriteHeader(http.StatusOK)
    } else {
        http.Error(w, errMsg, http.StatusServiceUnavailable)
    }
}
