// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteNATGWRequest is request schema for DeleteNATGW action
type DeleteNATGWRequest struct {
	request.CommonBase

	// [公共参数] 项目Id。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// NAT网关Id
	NATGWId *string `required:"true"`

	// 是否释放绑定的EIP。true：解绑并释放；false：只解绑不释放。默认为false
	ReleaseEip *bool `required:"false"`
}

// DeleteNATGWResponse is response schema for DeleteNATGW action
type DeleteNATGWResponse struct {
	response.CommonBase
}

// NewDeleteNATGWRequest will create request of DeleteNATGW action.
func (c *VPCClient) NewDeleteNATGWRequest() *DeleteNATGWRequest {
	req := &DeleteNATGWRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteNATGW - 删除NAT网关
func (c *VPCClient) DeleteNATGW(req *DeleteNATGWRequest) (*DeleteNATGWResponse, error) {
	var err error
	var res DeleteNATGWResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteNATGW", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
