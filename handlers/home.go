package handlers

import (
	"LRProject3/utils"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetNoCacheHeaders(w)
	http.ServeFile(w, r, "static/home.html")
}
