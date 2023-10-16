package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ESXi 或 vCenter 的 URL、用户名和密码
	esxiURL := "https://192.168.99.4/sdk"
	username := "root"
	password := "qwer@1234"

	// 解析 URL 并创建连接
	u, err := url.Parse(esxiURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	u.User = url.UserPassword(username, password)

	// 连接到 ESXi 或 vCenter
	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		fmt.Printf("Error connecting to ESXi: %v\n", err)
		return
	}

	// 创建一个视图管理器
	m := view.NewManager(c.Client)

	// 创建一个容器视图，查看所有虚拟机对象
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		fmt.Printf("Error creating view: %v\n", err)
		return
	}

	defer v.Destroy(ctx)

	// 使用属性收集器检索虚拟机的摘要属性
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		fmt.Printf("Error retrieving VM information: %v\n", err)
		return
	}

	// 列出找到的虚拟机
	for _, vm := range vms {
		fmt.Printf("Found VM: %s, state: %s\n", vm.Summary.Config.Name, vm.Summary.Runtime.PowerState)
	}
}
