package datastore

import (
	"github.com/gilcrest/errors"
	"github.com/gomodule/redigo/redis"
)

// newCacheDb returns an pool of redis connections from
// which an application can get a new connection
func newCacheDb() *redis.Pool {
	const op errors.Op = "env/datastore/newCacheDb"
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// RedisConn gets a connection from ds.cacheDB redis cache
func (ds Datastore) RedisConn() (redis.Conn, error) {
	const op errors.Op = "env/datastore/Datastore.RedisConn"

	conn := ds.cacheDB.Get()

	err := conn.Err()
	if err != nil {
		return nil, errors.E(op, err)
	}
	return conn, nil
}
