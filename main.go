package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	serverConfigs := []constant.ServerConfig{
		{
			Scheme:      "http",
			ContextPath: "/nacos",
			IpAddr:      "127.0.0.1",
			Port:        8848,
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         "",
		TimeoutMs:           10 * 1000,
		NotLoadCacheAtStart: true,
		LogDir:              "./log",
		CacheDir:            "./cache",
		LogLevel:            "debug",
	}
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	serviceName := "judge"
	ip := "127.0.0.1"
	port := 8888
	// 注册一个实例
	instance := vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        uint64(port),
		ServiceName: serviceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    map[string]string{},
		Ephemeral:   true,
	}
	success, err := namingClient.RegisterInstance(instance)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(success)
}
