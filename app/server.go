package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var server *OVHTestServer

type OVHTestServer struct {
	DbType      string
	Database    *gorm.DB
	TodosStore  TodosStore
	Router      *mux.Router
	AllowOrigin string
}

func (s OVHTestServer) ListenAndServe(port string) error {
	return http.ListenAndServe(":"+port, s.Router)
}

func SetServer(s *OVHTestServer) {
	server = s
}

func GetServer() *OVHTestServer {
	return server
}
