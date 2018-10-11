//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api VPC DescribeVPCIntercom

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeVPCIntercomRequest is request schema for DescribeVPCIntercom action
type DescribeVPCIntercomRequest struct {
	request.CommonBase

	// VPC短ID
	VPCId *string `required:"true"`

	// 目的地域
	DstRegion *string `required:"false"`

	// 目的项目ID
	DstProjectId *string `required:"false"`
}

// DescribeVPCIntercomResponse is response schema for DescribeVPCIntercom action
type DescribeVPCIntercomResponse struct {
	response.CommonBase

	// 联通VPC信息数组
	DataSet []VPCIntercomInfo
}

// NewDescribeVPCIntercomRequest will create request of DescribeVPCIntercom action.
func (c *VPCClient) NewDescribeVPCIntercomRequest() *DescribeVPCIntercomRequest {
	req := &DescribeVPCIntercomRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeVPCIntercom - 获取VPC互通信息
func (c *VPCClient) DescribeVPCIntercom(req *DescribeVPCIntercomRequest) (*DescribeVPCIntercomResponse, error) {
	var err error
	var res DescribeVPCIntercomResponse

	err = c.client.InvokeAction("DescribeVPCIntercom", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
