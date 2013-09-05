package goredis

import (
	"./storage"
)

// GoRedisServer
type GoRedisServer struct {
	RedisServer
	Storages *storage.RedisStorages
}

func NewGoRedisServer() (server *GoRedisServer) {
	server = &GoRedisServer{}
	server.Init()
	server.Storages = &storage.RedisStorages{}
	server.Storages.StringStorage = storage.NewMemoryStringStorage()
	return
}

func (server *GoRedisServer) Init() {
	server.RedisServer.Init()
	server.initForKeys()
	server.initForStrings()
}