package test

import (
	"errors"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

// CheckServerWorking 连接服务器,确保API服务能够正常访问.
func CheckServerWorking(max_num int) error {
	for i := 0; i < max_num; i++ {
		// 连接API服务器的"/health"GET请求接口
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// 请求失败后,等待1秒，然后进行下一次的重新连接访问
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Can not connect to the router.")
}
