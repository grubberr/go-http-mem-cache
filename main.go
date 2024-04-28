package main

import (
	"encoding/json"
	"github.com/grubberr/go-http-mem-cache/lrucache"
	"io"
	"net/http"
)

type CacheService struct {
	cache *lrucache.LRUCache
}

func (s *CacheService) setCacheHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) // 10Mb
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key := r.FormValue("key")
	if key == "" {
		http.Error(w, "key field is required", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	s.cache.Set(key, data)

	w.WriteHeader(http.StatusCreated)
}

func (s *CacheService) getCacheHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	if key == "" {
		http.Error(w, "key field is required", http.StatusBadRequest)
		return
	}

	data, ok := s.cache.Get(key)
	if !ok {
		http.Error(w, "key not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)
}

func (s *CacheService) getStatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.MarshalIndent(s.cache.GetTopKeys(5), "", " ")
	w.Write(res)
}

func main() {
	service := CacheService{cache: lrucache.NewLRUCache(10)}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /cache/", service.setCacheHandler)
	mux.HandleFunc("GET /cache/", service.getCacheHandler)
	mux.HandleFunc("GET /stat/", service.getStatHandler)
	http.ListenAndServe("localhost:8000", mux)
}
