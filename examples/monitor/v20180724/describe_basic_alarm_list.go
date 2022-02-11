package main

import (
	"fmt"
	"os"

	"github.com/sesky4/tgo/tencentcloud/common"
	"github.com/sesky4/tgo/tencentcloud/common/errors"
	"github.com/sesky4/tgo/tencentcloud/common/profile"
	"github.com/sesky4/tgo/tencentcloud/common/regions"
	monitor "github.com/sesky4/tgo/tencentcloud/monitor/v20180724"
)

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	client, _ := monitor.NewClient(credential, regions.Guangzhou, profile.NewClientProfile())
	request := monitor.NewDescribeBasicAlarmListRequest()
	request.Module = common.StringPtr("monitor")
	request.ProjectIds = common.Int64Ptrs([]int64{0})
	response, err := client.DescribeBasicAlarmList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
