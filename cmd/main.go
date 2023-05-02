package main

import (
	// "github.com/asgharalitaj/chatychat"
	"os"
	"github.com/asgharalitaj/chatychat/postgres"
	"github.com/asgharalitaj/chatychat/web"
	"net/http"
)

func main() {
	postgres.Init()

	h := postgres.NewThreadStore()
	handler := web.NewHandler(h)

	http.ListenAndServe(":3000", handler)

}
