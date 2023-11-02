package vo

import "github.com/yefengzhichen/nacos-sdk-go-v1x/common/constant"

type NacosClientParam struct {
	ClientConfig  *constant.ClientConfig  // optional
	ServerConfigs []constant.ServerConfig // optional
}
