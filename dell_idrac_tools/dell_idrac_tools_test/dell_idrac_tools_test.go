package dell_idrac_tools_test

import (
	"github.com/george012/esxi_box/dell_idrac_tools"
	"testing"
)

func TestConnect(t *testing.T) {
	// iDRAC 的 IP 地址和登录凭据
	idracIP := "192.168.99.5"
	username := "root"
	password := "qwe@123"

	dell_idrac_tools.ConnectIDRACDevice(idracIP, username, password)
}
