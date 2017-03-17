package liar

import (
	"net/http"
)

func New(humanHandler, botHandler http.Handler) http.Handler {

	agentsMap := make(map[string]bool)

	for _, item := range agentSlice {
		agentsMap[item] = true
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := agentsMap[r.UserAgent()]; ok {
			botHandler.ServeHTTP(w, r)
		} else {
			humanHandler.ServeHTTP(w, r)
		}
	})
}
