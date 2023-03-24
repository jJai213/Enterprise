package resources

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// add API key here
const api_key = "56cb44ed64fee7c9dfe03a74ccee5db3"


func search(w http.ResponseWriter, r *http.Request) {


	http_request := map[string]interface{} {}

	if err := json.NewDecoder(r.Body).Decode(&http_request); err == nil {
		if audio, err := http_request["Audio"]; err {
			request := map[string]interface{} {"api_token" : api_key, "audio" : audio}

			if marsh, err := json.Marshal(request); err == nil {
				data := bytes.NewBuffer(marsh)
				if api_responce, err := http.Post("https://api.audd.io/recognize", "application/json", data); err == nil {
					if api_responce.StatusCode == http.StatusOK {
						defer api_responce.Body.Close()
						if marshbody, err := io.ReadAll(api_responce.Body); err == nil {
							apibody := Rep{}
							if err = json.Unmarshal(marshbody, &apibody); err == nil {
								if apibody.Status == "success" {
									result := map[string]interface{}{"Id": apibody.Result.Title}
									if err = json.NewEncoder(w).Encode(result); err == nil {
										w.WriteHeader(200)
									} else {
										w.WriteHeader(500)
									}
								} else {
									w.WriteHeader(500)
								}
							} else {
								w.WriteHeader(500)
							}

						} else {
							w.WriteHeader(500)
						}
					} else {
						w.WriteHeader(500)
					}
				} else {
					w.WriteHeader(500)
				} 
			} else {
				w.WriteHeader(500)
			}
		} else {
			w.WriteHeader(400)
		}
	} else {
		w.WriteHeader(500)
	}
}

func Router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/search", search).Methods("POST")

	return r

}