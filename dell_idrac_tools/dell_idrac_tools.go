package dell_idrac_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bmc-toolbox/bmclib/v2"
	"github.com/george012/gtbox/gtbox_log"
)

func ConnectIDRACDevice(idrac_address string, username string, password string) {

	clientOpts := []bmclib.Option{}
	// init client
	client := bmclib.NewClient(idrac_address, username, password, clientOpts...)

	// open BMC session
	err := client.Open(context.Background())
	if err != nil {
		gtbox_log.LogErrorf("bmc login failed%s", err)
	}

	defer client.Close(context.Background())

	// retrieve inventory data
	inventory, err := client.Inventory(context.Background())
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	b, err := json.MarshalIndent(inventory, "", "  ")
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	fmt.Println(string(b))
}
