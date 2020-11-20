package unixService

import (
	"encoding/json"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

var sockPath = "/tmp/unix.sock"

//Serve 對內unix server服務
func Serve(eventPool chan map[string]interface{}) {
	os.Remove(sockPath)
	listen, err := net.Listen("unix", sockPath)
	if err != nil {
		panic(err)
	}
	log.Info("Unix Server Listening...")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info("Connected From:", conn.RemoteAddr().String())
		go func() {
			parseAndAdapteData(conn, eventPool)
		}()
	}
}

func parseAndAdapteData(conn net.Conn, eventPool chan map[string]interface{}) {
	for {
		m := <-eventPool
		data, err := json.Marshal(m)
		if err != nil {
			log.Error(err)
			continue
		}
		_, err = conn.Write(data)
		if err != nil {
			log.Error(err)
		}
	}
}
