package main

import "net/http"


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

    if ok {
        w.WriteHeader(http.StatusOK)
    } else {
        http.Error(w, errMsg, http.StatusServiceUnavailable)
    }
}
