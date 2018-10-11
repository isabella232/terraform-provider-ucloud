//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UAccount GetProjectList

package uaccount

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// GetProjectListRequest is request schema for GetProjectList action
type GetProjectListRequest struct {
	request.CommonBase

	// 是否是财务账号(Yes: 是, No: 否)
	IsFinance *string `required:"false"`
}

// GetProjectListResponse is response schema for GetProjectList action
type GetProjectListResponse struct {
	response.CommonBase

	// 项目总数
	ProjectCount int

	// JSON格式的项目列表实例
	ProjectSet []ProjectListInfo
}

// NewGetProjectListRequest will create request of GetProjectList action.
func (c *UAccountClient) NewGetProjectListRequest() *GetProjectListRequest {
	req := &GetProjectListRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetProjectList - 获取项目列表
func (c *UAccountClient) GetProjectList(req *GetProjectListRequest) (*GetProjectListResponse, error) {
	var err error
	var res GetProjectListResponse

	err = c.client.InvokeAction("GetProjectList", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
