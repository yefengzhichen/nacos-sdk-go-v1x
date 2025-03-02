/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package naming_client

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/yefengzhichen/nacos-sdk-go-v1x/clients/nacos_client"
	"github.com/yefengzhichen/nacos-sdk-go-v1x/common/constant"
	"github.com/yefengzhichen/nacos-sdk-go-v1x/common/http_agent"
	"github.com/yefengzhichen/nacos-sdk-go-v1x/mock"
	"github.com/yefengzhichen/nacos-sdk-go-v1x/model"
	"github.com/yefengzhichen/nacos-sdk-go-v1x/vo"
)

var clientConfigTest = *constant.NewClientConfig(
	constant.WithTimeoutMs(10*1000),
	constant.WithBeatInterval(5*1000),
	constant.WithNotLoadCacheAtStart(true),
)

var serverConfigTest = *constant.NewServerConfig("console.nacos.io", 80, constant.WithContextPath("/nacos"))

func Test_RegisterServiceInstance_withoutGroupName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Request(gomock.Eq("POST"),
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"namespaceId": "",
			"serviceName": "DEFAULT_GROUP@@DEMO",
			"groupName":   "DEFAULT_GROUP",
			"app":         "",
			"clusterName": "",
			"ip":          "10.0.0.10",
			"port":        "80",
			"weight":      "0",
			"enable":      "false",
			"healthy":     "false",
			"metadata":    "{}",
			"ephemeral":   "false",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)
	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	success, err := client.RegisterInstance(vo.RegisterInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Ephemeral:   false,
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func Test_RegisterServiceInstance_withGroupName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	mockIHttpAgent.EXPECT().Request(gomock.Eq("POST"),
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"namespaceId": "",
			"serviceName": "test_group@@DEMO2",
			"groupName":   "test_group",
			"app":         "",
			"clusterName": "",
			"ip":          "10.0.0.10",
			"port":        "80",
			"weight":      "0",
			"enable":      "false",
			"healthy":     "false",
			"metadata":    "{}",
			"ephemeral":   "false",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)

	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	success, err := client.RegisterInstance(vo.RegisterInstanceParam{
		ServiceName: "DEMO2",
		Ip:          "10.0.0.10",
		Port:        80,
		GroupName:   "test_group",
		Ephemeral:   false,
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func Test_RegisterServiceInstance_withCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	mockIHttpAgent.EXPECT().Request(gomock.Eq("POST"),
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"namespaceId": "",
			"serviceName": "test_group@@DEMO3",
			"groupName":   "test_group",
			"app":         "",
			"clusterName": "test",
			"ip":          "10.0.0.10",
			"port":        "80",
			"weight":      "0",
			"enable":      "false",
			"healthy":     "false",
			"metadata":    "{}",
			"ephemeral":   "false",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)

	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	success, err := client.RegisterInstance(vo.RegisterInstanceParam{
		ServiceName: "DEMO3",
		Ip:          "10.0.0.10",
		Port:        80,
		GroupName:   "test_group",
		ClusterName: "test",
		Ephemeral:   false,
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func Test_RegisterServiceInstance_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	mockIHttpAgent.EXPECT().Request(gomock.Eq("POST"),
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"namespaceId": "",
			"serviceName": "test_group@@DEMO4",
			"groupName":   "test_group",
			"app":         "",
			"clusterName": "",
			"ip":          "10.0.0.10",
			"port":        "80",
			"weight":      "0",
			"enable":      "false",
			"healthy":     "false",
			"metadata":    "{}",
			"ephemeral":   "false",
		})).Times(3).
		Return(http_agent.FakeHttpResponse(401, `no security`), nil)

	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	result, err := client.RegisterInstance(vo.RegisterInstanceParam{
		ServiceName: "DEMO4",
		Ip:          "10.0.0.10",
		Port:        80,
		GroupName:   "test_group",
		Ephemeral:   false,
	})
	assert.Equal(t, false, result)
	assert.NotNil(t, err)
}

func TestNamingProxy_DeregisterService_WithoutGroupName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	mockIHttpAgent.EXPECT().Request(gomock.Eq("DELETE"),
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"namespaceId": "",
			"serviceName": "DEFAULT_GROUP@@DEMO5",
			"clusterName": "",
			"ip":          "10.0.0.10",
			"port":        "80",
			"ephemeral":   "true",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)
	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	_, _ = client.DeregisterInstance(vo.DeregisterInstanceParam{
		ServiceName: "DEMO5",
		Ip:          "10.0.0.10",
		Port:        80,
		Ephemeral:   true,
	})
}

func TestNamingProxy_DeregisterService_WithGroupName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	mockIHttpAgent.EXPECT().Request(gomock.Eq("DELETE"),
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"namespaceId": "",
			"serviceName": "test_group@@DEMO6",
			"clusterName": "",
			"ip":          "10.0.0.10",
			"port":        "80",
			"ephemeral":   "true",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)
	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	_, _ = client.DeregisterInstance(vo.DeregisterInstanceParam{
		ServiceName: "DEMO6",
		Ip:          "10.0.0.10",
		Port:        80,
		GroupName:   "test_group",
		Ephemeral:   true,
	})
}

func Test_UpdateServiceInstance_withoutGroupName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Request(gomock.Eq("PUT"),
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"namespaceId": "",
			"serviceName": "DEFAULT_GROUP@@DEMO",
			"clusterName": "",
			"ip":          "10.0.0.10",
			"port":        "80",
			"weight":      "0",
			"enable":      "false",
			"metadata":    "{}",
			"ephemeral":   "false",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)
	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	success, err := client.UpdateInstance(vo.UpdateInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Ephemeral:   false,
		Metadata:    map[string]string{},
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func TestNamingProxy_DeregisterService_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	mockIHttpAgent.EXPECT().Request(gomock.Eq("DELETE"),
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"namespaceId": "",
			"serviceName": "test_group@@DEMO7",
			"clusterName": "",
			"ip":          "10.0.0.10",
			"port":        "80",
			"ephemeral":   "true",
		})).Times(3).
		Return(http_agent.FakeHttpResponse(401, `no security`), nil)
	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	_, _ = client.DeregisterInstance(vo.DeregisterInstanceParam{
		ServiceName: "DEMO7",
		Ip:          "10.0.0.10",
		Port:        80,
		GroupName:   "test_group",
		Ephemeral:   true,
	})
}

func TestNamingClient_SelectOneHealthyInstance_SameWeight(t *testing.T) {
	services := model.Service(model.Service{
		Name:            "DEFAULT_GROUP@@DEMO",
		CacheMillis:     1000,
		UseSpecifiedURL: false,
		Hosts: []model.Instance{
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.10-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.10",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO1",
				Enable:      true,
				Healthy:     true,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.11-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.11",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      true,
				Healthy:     true,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.12-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.12",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      true,
				Healthy:     false,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.13-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.13",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      false,
				Healthy:     true,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.14-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.14",
				Weight:      0,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      true,
				Healthy:     true,
			},
		},
		Checksum:    "3bbcf6dd1175203a8afdade0e77a27cd1528787794594",
		LastRefTime: 1528787794594, Env: "", Clusters: "a",
		Metadata: map[string]string(nil)})
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	instance1, err := client.selectOneHealthyInstances(services)
	assert.Nil(t, err)
	assert.NotNil(t, instance1)
	instance2, err := client.selectOneHealthyInstances(services)
	assert.Nil(t, err)
	assert.NotNil(t, instance2)
}

func TestNamingClient_SelectOneHealthyInstance_Empty(t *testing.T) {
	services := model.Service(model.Service{
		Name:            "DEFAULT_GROUP@@DEMO",
		CacheMillis:     1000,
		UseSpecifiedURL: false,
		Hosts:           []model.Instance{},
		Checksum:        "3bbcf6dd1175203a8afdade0e77a27cd1528787794594",
		LastRefTime:     1528787794594, Env: "", Clusters: "a",
		Metadata: map[string]string(nil)})
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	instance, err := client.selectOneHealthyInstances(services)
	assert.NotNil(t, err)
	assert.Nil(t, instance)
}

func TestNamingClient_SelectInstances_Healthy(t *testing.T) {
	services := model.Service(model.Service{
		Name:            "DEFAULT_GROUP@@DEMO",
		CacheMillis:     1000,
		UseSpecifiedURL: false,
		Hosts: []model.Instance{
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.10-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.10",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      true,
				Healthy:     true,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.11-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.11",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      true,
				Healthy:     true,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.12-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.12",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      true,
				Healthy:     false,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.13-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.13",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      false,
				Healthy:     true,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.14-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.14",
				Weight:      0,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      true,
				Healthy:     true,
			},
		},
		Checksum:    "3bbcf6dd1175203a8afdade0e77a27cd1528787794594",
		LastRefTime: 1528787794594, Env: "", Clusters: "a",
		Metadata: map[string]string(nil)})
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	instances, err := client.selectInstances(services, true)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(instances))
}

func TestNamingClient_SelectInstances_Unhealthy(t *testing.T) {
	services := model.Service(model.Service{
		Name:            "DEFAULT_GROUP@@DEMO",
		CacheMillis:     1000,
		UseSpecifiedURL: false,
		Hosts: []model.Instance{
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.10-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.10",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      true,
				Healthy:     true,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.11-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.11",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      true,
				Healthy:     true,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.12-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.12",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      true,
				Healthy:     false,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.13-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.13",
				Weight:      1,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      false,
				Healthy:     true,
			},
			{
				Valid:       true,
				Marked:      false,
				InstanceId:  "10.10.10.14-80-a-DEMO",
				Port:        80,
				Ip:          "10.10.10.14",
				Weight:      0,
				Metadata:    map[string]string{},
				ClusterName: "a",
				ServiceName: "DEMO",
				Enable:      true,
				Healthy:     true,
			},
		},
		Checksum:    "3bbcf6dd1175203a8afdade0e77a27cd1528787794594",
		LastRefTime: 1528787794594, Env: "", Clusters: "a",
		Metadata: map[string]string(nil)})
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	instances, err := client.selectInstances(services, false)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(instances))
}

func TestNamingClient_SelectInstances_Empty(t *testing.T) {
	services := model.Service(model.Service{
		Name:            "DEFAULT_GROUP@@DEMO",
		CacheMillis:     1000,
		UseSpecifiedURL: false,
		Hosts:           []model.Instance{},
		Checksum:        "3bbcf6dd1175203a8afdade0e77a27cd1528787794594",
		LastRefTime:     1528787794594, Env: "", Clusters: "a",
		Metadata: map[string]string(nil)})
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(mockIHttpAgent)
	client, _ := NewNamingClient(&nc)
	instances, err := client.selectInstances(services, false)
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(instances))
}

func TestNamingClient_GetAllServicesInfo(t *testing.T) {
	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	_ = nc.SetHttpAgent(&http_agent.HttpAgent{})
	client, _ := NewNamingClient(&nc)
	reslut, err := client.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		GroupName: "DEFAULT_GROUP",
		PageNo:    1,
		PageSize:  20,
	})

	assert.NotNil(t, reslut.Doms)
	assert.Nil(t, err)
}

func TestNamingClient_selectOneHealthyInstanceResult(t *testing.T) {
	services := model.Service(model.Service{
		Name: "DEFAULT_GROUP@@DEMO",
		Hosts: []model.Instance{
			{
				Ip:      "127.0.0.1",
				Weight:  1,
				Enable:  true,
				Healthy: true,
			},
			{
				Ip:      "127.0.0.2",
				Weight:  9,
				Enable:  true,
				Healthy: true,
			},
		}})
	nc := nacos_client.NacosClient{}
	_ = nc.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = nc.SetClientConfig(clientConfigTest)
	client, _ := NewNamingClient(&nc)
	for i := 0; i < 10; i++ {
		_, _ = client.selectOneHealthyInstances(services)
	}
}
