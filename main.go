package main

import (
	"encoding/json"
	"net/http"

	spinhttp "github.com/rajatjindal/wasip2-string/sdk/http"
	"github.com/rajatjindal/wasip2-string/sdk/kv"
	"github.com/rajatjindal/wasip2-string/sdk/variables"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		// kv test
		store, err := kv.OpenDefault()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = store.Set("testme", []byte("oksure"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = store.Set("testme-2", []byte("oksure"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = store.Set("testme-3", []byte("oksure"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		keys, err := store.GetKeys()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		val, err := variables.Get("password")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		output := map[string]interface{}{
			"variables": val,
			"keys":      keys,
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(output)
	})
}

func main() {}
