package main

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	// asyncModel := asyncmodel.AsyncModel{Version: asyncmodel.Version1, Msgid: asyncmodel.MsgidResp1112, Srcid: 1, Dstid: 1}
	// buf := asyncModel.Encode()
	// asyncModel.Decode(buf)
	// go udp.Listen()
	// udp.Dial()
	logger := log.New(os.Stdout, "// DEBUG: ", log.LstdFlags|log.Lshortfile)
	// logger.Printf("%s", "你好")
	session, err := mgo.Dial("192.168.10.91:10000")
	if err != nil {
		logger.Fatalln("连接MongoDB服务异常")
	}
	defer session.Close()
	session.SetMode(mgo.SecondaryPreferred, true)
	c := session.DB("sdnDB").C("stat_data")
	filter := bson.M{}
	count, err := c.Find(filter).Count()
	if err != nil {
		logger.Fatalln("查询异常:", err)
	}
	logger.Printf("stat_data集合中, 文档总数：%d", count)
	pageCount := 10000
	if count < pageCount {
		pageCount = count
	}
	iter := c.Find(filter).Limit(pageCount).Iter()
	result := &bson.M{}
	for iter.Next(result) {
		logger.Printf("%v", result)
	}
}
