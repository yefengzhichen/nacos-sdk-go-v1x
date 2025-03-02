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

package cache

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yefengzhichen/nacos-sdk-go-v1x/common/constant"
)

func TestGetFileName(t *testing.T) {

	name := GetFileName("nacos@@providers:org.apache.dubbo.UserProvider:hangzhou", "tmp")

	if runtime.GOOS == constant.OS_WINDOWS {
		assert.Equal(t, name, "tmp\\nacos@@providers&&org.apache.dubbo.UserProvider&&hangzhou")
	} else {
		assert.Equal(t, name, "tmp/nacos@@providers:org.apache.dubbo.UserProvider:hangzhou")
	}
}
