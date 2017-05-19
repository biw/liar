package liar

import (
	"net/http"
)

type LiarRouter struct {
	agentMap   map[string]bool
	botHandler http.Handler
}

func New(botHandler http.Handler) LiarRouter {

	newLiar := LiarRouter{}

	newLiar.agentMap = make(map[string]bool)

	for _, item := range agentSlice {
		newLiar.agentMap[item] = true
	}

	newLiar.botHandler = botHandler

	return newLiar
}

func NewFunc(f func(http.ResponseWriter, *http.Request)) LiarRouter {
	return New(http.HandlerFunc(f))
}

func (LR *LiarRouter) Human(humanHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := LR.agentMap[r.UserAgent()]; ok {
			LR.botHandler.ServeHTTP(w, r)
		} else {
			humanHandler.ServeHTTP(w, r)
		}
	})
}

func (LR *LiarRouter) HumanFunc(f func(http.ResponseWriter, *http.Request)) http.Handler {
	return LR.Human(http.HandlerFunc(f))
}
