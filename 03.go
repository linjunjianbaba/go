package main

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

func main() {
	client, err := slb.NewClientWithAccessKey("cn-shenzhen", "BsBpRuF5sUPLxysV", "Ihvc21Pt9ljOleMfJrIuBBzCN5fQNe")

	request := slb.CreateSetBackendServersRequest()

	request.BackendServers = `[{"Type":"ecs","ServerID":"i-wz9ccq9xttl0jouyczje","Weight":"50" }]`
	//  request.BackendServers = "[{"Type":"ecs","ServerId":"i-wz9ccq9xttl0jouyczje","Weight":"50"}]"
	request.LoadBalancerId = "lb-wz9lhl0u7dmsivupl91y3"

	response, err := client.SetBackendServers(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
