package Redis

import(
	// "github.com/garyburd/redigo/redis"
	"github.com/go-redis/redis"
)

var conn *redis.Client

func newPool() *redis.Client {
	// return &redis.Pool{
	// 	// Maximum number of idle connections in the pool.
	// 	MaxIdle: 80,
	// 	// max number of connections
	// 	MaxActive: 12000,
	// 	// Dial is an application supplied function for creating and
	// 	// configuring a connection.
	// 	Dial: func() (redis.Conn, error) {
	// 		c, err := redis.Dial("tcp", ":6379")
	// 		if err != nil {
	// 			panic(err.Error())
	// 		}
	// 		return c, err
	// 	},
	// }
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func init(){
	// pool:=newPool()
	// conn=pool.Get()
	pool:=newPool()
	conn=pool
}

func getRedisDB()(*redis.Client){
	return conn
}