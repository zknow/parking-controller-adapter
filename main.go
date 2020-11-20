package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	flock "github.com/theckman/go-flock"
	"github.com/zknow/parkingChargeAdapter/httpService"
	"github.com/zknow/parkingChargeAdapter/unixService"
)

func main() {
	l := lock("lock", false)

	if l == nil {
		log.Warning("ParkingChargeAdapter 正在執行")
	}
	defer func() { _ = l.Unlock() }()

	e := createEventPool()
	go httpService.Serve(e)
	go unixService.Serve(e)

	safeExitNotify()
}

// 建立事件池
func createEventPool() chan map[string]interface{} {
	return make(chan map[string]interface{}, 10)
}

// 監聽安全退出
func safeExitNotify() {
	fmt.Println("[Info] start listen quit signal ...")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	log.Info("Get signal : ", <-sigs)
	os.Exit(1)
}

// 程式鎖
func lock(path string, loop bool) *flock.Flock {
	lock := flock.New(path)
	for {
		locked, err := lock.TryLock()
		if err != nil {
			log.Error("flock Error Path:", path)
		}
		if locked {
			break
		}
		if !loop {
			return nil
		}
		time.Sleep(time.Millisecond * 500)
	}
	return lock
}
