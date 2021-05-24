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

// 服务配置
type Service struct {
	Build        Build    `yaml:"build,omitempty"`         // 构建时的配置项
	CapAdd       []string `yaml:"cap_add,omitempty"`       // 添加的容器功能
	CapDrop      []string `yaml:"cap_drop,omitempty"`      // 添加的容器功能
	CgroupParent string   `yaml:"cgroup_parent,omitempty"` // 可选的父控制组
	Command      string `yaml:"command,omitempty"`       // 默认命令
	// TODO configs
	ContainerName string `yaml:"container_name,omitempty"` // 容器名称
	// TODO credential_spec
	DependsOn []string `yaml:"depends_on,omitempty"` // 服务之间的依赖关系
	// TODO deploy
	// TODO devices
	// TODO dns
	// TODO dns_search
	// TODO entrypoint
	// TODO env_file
	Environment map[string]interface{} `yaml:"environment,omitempty"` // 环境变量
	// TODO expose
	// TODO external_links
	// TODO extra_hosts
	// TODO healthcheck
	Image Image `yaml:"image"` // 容器启动的镜像
	// TODO init
	// TODO isolation
	// TODO labels
	// TODO links
	// TODO logging
	NetworkMode string                `yaml:"network_mode,omitempty"` // 网络模式，可参见常量
	Networks    map[string]NetworkMap `yaml:"networks,omitempty"`     // 加入的网络
	// TODO pid
	Ports   []Port   `yaml:"ports,omitempty"`   // 暴露的端口号
	Restart string   `yaml:"restart,omitempty"` // 重启策略
	Secrets []Secret `yaml:"secrets,omitempty"` // 密钥
	// TODO security_opt
	// TODO stop_grace_period
	// TODO stop_signal
	// TODO sysctls
	// TODO tmpfs
	// TODO ulimits
	// TODO userns_mode
	Volumes []VolumeMap `yaml:"volumes,omitempty"` // 挂载卷
}
