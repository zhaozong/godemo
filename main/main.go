package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	_ "net/http/pprof"
	"time"
	"unsafe"
)

type a struct {
	number int
}

func main() {
}

//func testPprof() {
//	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
//	flag.Parse()
//	if *cpuprofile != "" {
//		f, err := os.Create(*cpuprofile)
//		if err != nil {
//			log.Fatal(err)
//		}
//		pprof.StartCPUProfile(f)
//		defer pprof.StopCPUProfile()
//	}
//	for {
//		fmt.Println(Add("test pprof "))
//		time.Sleep(time.Microsecond * 100)
//	}
//}
//func test() {
//	go func() {
//		for {
//			fmt.Println(Add("test pprof "))
//			time.Sleep(1 * time.Second)
//		}
//	}()
//	http.ListenAndServe("localhost:6060", nil)
//}

func byte2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func redisClient() {
	var client = redis.NewFailoverClient(&redis.FailoverOptions{
		SentinelAddrs: []string{"10.10.108.75:26379", "10.10.108.76:26379", "10.10.108.78:26379"},
		MasterName:    "myredismaster",
		Password:      "qn7XrrXRDtR6QBv82xqMi",
		DB:            1,
	})

	value, err := client.Del("aaa", "bbb", "ccc").Result()
	fmt.Printf("a%#v\n", value)
	fmt.Printf("b%#v\n", err)
}

func insert(c *mgo.Collection) {
	err := c.Insert(&User{
		Id_:       bson.NewObjectId(),
		Name:      "AAA",
		Age:       3,
		JonedAt:   time.Now(),
		Interests: []string{"aaa", "bbb"},
	})
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id_       bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Age       int           `bson:"age"`
	JonedAt   time.Time     `bson:"joned_at"`
	Interests []string      `bson:"interests"`
}
