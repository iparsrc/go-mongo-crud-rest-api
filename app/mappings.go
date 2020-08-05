package app

import "github.com/parsaakbari1209/go-mongo-CRUD-web/controllers/ping"

func MapUrls() {
	router.GET("/ping", ping.Ping)
}
