/*
   Copyright (c) 2020 XiaochengTech
   gitee.com/xiaochengtech/docker is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
       http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package docker

import (
	"errors"
	"fmt"
	"strings"
)

// 路径(Short Syntax)
type VolumeMap struct {
	Host      string // 外部主机的路径
	Container string // 内部容器的路径
	Mode      string // 权限
}

// 新建一个路径A到路径B的映射
func NewVolumeMap(hostVolumeMap string, containerVolumeMap string) VolumeMap {
	return VolumeMap{
		Host:      hostVolumeMap,
		Container: containerVolumeMap,
	}
}

// 新建一个相同的路径映射
func NewVolumeMapSame(volumeMap string) VolumeMap {
	vm := NewVolumeMap(volumeMap, volumeMap)
	vm.Mode = VolumeReadOnly
	return vm
}

func (m VolumeMap) MarshalYAML() (result interface{}, err error) {
	if len(m.Host) == 0 {
		err = errors.New("docker: simple-volume-map host can not be empty")
		return
	}
	tmp := m.Host
	if len(m.Container) > 0 {
		tmp += fmt.Sprintf(":%s", m.Container)
		if len(m.Mode) > 0 {
			tmp += fmt.Sprintf(":%s", m.Mode)
		}
	}
	result = tmp
	return
}

func (m *VolumeMap) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var origin string
	if err = unmarshal(&origin); err != nil {
		return
	}
	// 拆分
	parts := strings.Split(origin, ":")
	if len(parts) > 3 {
		err = errors.New("docker: simple-volume-map format error")
		return
	}
	m.Host = parts[0]
	if len(parts) > 1 {
		m.Container = parts[1]
	}
	if len(parts) > 2 {
		m.Mode = parts[2]
	}
	// 校验
	if len(m.Host) == 0 {
		err = errors.New("docker: simple-volume-map format error")
		return
	}
	if len(m.Container) == 0 && len(m.Mode) > 0 {
		err = errors.New("docker: simple-volume-map format error")
		return
	}
	return
}
