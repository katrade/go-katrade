package server

import "go-katrade/db"

func Init() {
	r := NewRouter()
	db.ConnectMongoDB()
	r.Run()
}