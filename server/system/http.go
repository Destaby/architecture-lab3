package system

import (
	"encoding/json"
	"github.com/Destaby/architecture-lab3/server/tools"
	"log"
	"net/http"
)

// System HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of system HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
	/*	if (r.Method == "GET" && r.URL.Path == "/balancers") {
			handleListBalancers(store, rw)
		} else*/ if (r.Method == "PUT" && r.URL.Path == "/machines") {
			handleUpdateWorkingStatus(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleUpdateWorkingStatus(r *http.Request, rw http.ResponseWriter, store *Store) {
	var c Machine
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Printf("Error decoding machine input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.UpdateMachineStatus(c.Id, c.Working)
	if err == nil {
		tools.WriteJsonOk(rw, &c)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

// func handleListBalancers(store *Store, rw http.ResponseWriter) {
// 	res, err := store.ListBalancers()
// 	if err != nil {
// 		log.Printf("Error making query to the db: %s", err)
// 		tools.WriteJsonInternalError(rw)
// 		return
// 	}
// 	tools.WriteJsonOk(rw, res)
// }
