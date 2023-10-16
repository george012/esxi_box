package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 创建 vSphere API 的 URL。通常包括用户名、密码和 vSphere 主机地址。
	urlString := "https://root:qwer@1234@192.168.99.4/sdk"
	u, err := url.Parse(urlString)
	if err != nil {
		log.Fatalf("解析 URL 出错: %v", err)
	}

	// 使用 URL 创建一个新的 govmomi 客户端实例。
	client, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		log.Fatalf("连接到 vSphere 出错: %v", err)
	}

	// 创建一个新的 view manager
	m := view.NewManager(client.Client)

	// 创建一个新的容器视图，查看 VirtualMachine 对象
	v, err := m.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		log.Fatalf("创建容器视图出错: %v", err)
	}

	// 检索我们关心的对象的摘要
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, nil, &vms)
	if err != nil {
		log.Fatalf("检索虚拟机出错: %v", err)
	}

	// 处理结果，例如打印虚拟机的名称
	for _, vm := range vms {
		fmt.Printf("虚拟机: %s\n", vm.Name)
	}
}
