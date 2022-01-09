package controller

import (
  "fmt"
  "io"
  "net/http"
)

func GetHeader(http.ResponseWriter, r *http.Request){
	for k, v := range r.Header {
		res := strconverter(v)
		w.Header().Set(k, res)
	}
	w.WriteHeader(http.StatusOK)
}

func strconverter(org []string) (res string) {
	for k, v := range org {
		res += v
		if k != len(slice)-1 {
			res += ", "
		}
	}
	return res
}
