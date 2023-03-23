package resources

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tracks/repository"

	"github.com/gorilla/mux"
)

func addCell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var c repository.Cell
	if err := json.NewDecoder(r.Body).Decode(&c); err == nil {
		if id == c.Id {
			if n := repository.Update(c); n > 0 {
				w.WriteHeader(204)
			} else if n := repository.Insert(c); n > 0 {
				w.WriteHeader(201)
			} else {
				w.WriteHeader(500)
			}
		} else {
			w.WriteHeader(400)
		}
	} else {
		w.WriteHeader(400)
	}
}

func readCell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if c, n := repository.Read(id); n > 0 {
		d := repository.Cell{Id: c.Id, Audio: c.Audio}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(d)
	} else if n == 0 {
		w.WriteHeader(404)
	} else {
		w.WriteHeader(500)
	}
}

func deleteCell(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)
	id := data["id"]
	if n := repository.Delete(id); n > 0 {
		w.WriteHeader(204)
	} else if n == 0 {
		w.WriteHeader(404)
	} else {
		w.WriteHeader(500)
	}
}

func listCell(w http.ResponseWriter, r *http.Request) {
	if c, n := repository.ReadAll(); n > 0 {
		names := []string{}
		for _, cell := range c {
			names = append(names, cell.Id)
			fmt.Println(names)
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(names)
	} else {
		w.WriteHeader(500)
	}
}

func Router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/tracks/{id}", addCell).Methods("PUT")
	r.HandleFunc("/tracks", listCell).Methods("GET")
	r.HandleFunc("/tracks/{id}", readCell).Methods("GET")
	r.HandleFunc("/tracks/{id}", deleteCell).Methods("DELETE")

	return r

}