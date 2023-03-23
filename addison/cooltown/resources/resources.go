package resources

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func cooltown(w http.ResponseWriter, r *http.Request) {

	snippet := map[string]interface{} {}

	if err := json.NewDecoder(r.Body).Decode(&snippet); err == nil {
		if audio, err := snippet["Audio"]; err || audio != "" {
			request := map[string]interface{} {"Audio" : audio}
			if marsh, err := json.Marshal(request); err == nil {
				data := bytes.NewBuffer(marsh)
				if search_responce, err := http.Post("http://127.0.0.1:3001/search", "application/json", data); err == nil {
					if search_responce.StatusCode == http.StatusOK {
						defer search_responce.Body.Close()
						search_body := map[string]interface{} {}
						if err := json.NewDecoder(search_responce.Body).Decode(&search_body); err == nil {
							if sid, err := search_body["Id"]; err {
								searchurl := "http://127.0.0.1:3000/tracks/" + strings.Replace(sid.(string), " ", "+", -1)
								url_responce, err := http.Get(searchurl)

								if err != nil {
									
								}

								if url_responce.StatusCode == http.StatusOK {
									defer url_responce.Body.Close()

									url_body := map[string]interface{} {}
									if err := json.NewDecoder(url_responce.Body).Decode(&url_body); err == nil {
										if url_audio, err := url_body["Audio"]; err {
											res := map[string]interface{} {"Audio" : url_audio}
											if err := json.NewEncoder(w).Encode(res); err == nil {
												w.WriteHeader(200)
											} else {
												w.WriteHeader(200)
											}

										} else {
											w.WriteHeader(500)
										}
									} else {
										w.WriteHeader(500)
									}

									

								} else {
									w.WriteHeader(url_responce.StatusCode)
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
    r.HandleFunc("/cooltown", cooltown).Methods("POST")
	
	return r
} 