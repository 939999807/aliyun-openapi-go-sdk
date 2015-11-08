// Copyright 2015 Chen Xianren. All rights reserved.
// Code generated by openapi-meta; DO NOT EDIT

package ossadmin

import (
	"errors"
	"fmt"

	"git.oschina.net/cxr29/aliyun-openapi-go-sdk"
)

var (
	_ = errors.New("")
	_ = fmt.Sprint("")
	_ = make(openapi.M)
)

type API struct {
	openapi.Service
}

func New(accessKeyId, accessKeySecret string) API {
	return API{openapi.NewService(accessKeyId, accessKeySecret)}
}

func NewParams() openapi.Params {
	args := openapi.NewParams()
	args.Product = Product
	args.Style = Style
	args.Version = Version
	return args
}

const (
	Product = "OssAdmin"
	Style   = "RPC"
	Version = "2015-05-20"
)

// BindBucketVip version 2015-05-20
//
// required parameters:
//  name: BucketName, type: string
//  name: Region, type: string
//  name: Vip, type: string
//  name: VpcId, type: string
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) BindBucketVip(BucketName, Region, Vip, VpcId string, optional openapi.M) (*BindBucketVipResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "BindBucketVip")
	args.Query.Set("BucketName", BucketName)
	args.Query.Set("Region", Region)
	args.Query.Set("Vip", Vip)
	args.Query.Set("VpcId", VpcId)
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(BindBucketVipResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// BindBucketVipResponse represents the response of the api BindBucketVip.
type BindBucketVipResponse struct {
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateImgVpc version 2015-05-20
//
// required parameters:
//  name: Label, type: string
//  name: Region, type: string
//  name: VirtualSwitchId, type: string
//  name: VpcId, type: string
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) CreateImgVpc(Label, Region, VirtualSwitchId, VpcId string, optional openapi.M) (*CreateImgVpcResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "CreateImgVpc")
	args.Query.Set("Label", Label)
	args.Query.Set("Region", Region)
	args.Query.Set("VirtualSwitchId", VirtualSwitchId)
	args.Query.Set("VpcId", VpcId)
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(CreateImgVpcResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateImgVpcResponse represents the response of the api CreateImgVpc.
type CreateImgVpcResponse struct {
	RequestId string `json:"requestId" xml:"requestId"`
	Vip       string `json:"vip" xml:"vip"`
	VpcId     string `json:"vpcId" xml:"vpcId"`
}

// CreateOssVpc version 2015-05-20
//
// required parameters:
//  name: Label, type: string
//  name: Region, type: string
//  name: VirtualSwitchId, type: string
//  name: VpcId, type: string
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) CreateOssVpc(Label, Region, VirtualSwitchId, VpcId string, optional openapi.M) (*CreateOssVpcResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "CreateOssVpc")
	args.Query.Set("Label", Label)
	args.Query.Set("Region", Region)
	args.Query.Set("VirtualSwitchId", VirtualSwitchId)
	args.Query.Set("VpcId", VpcId)
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(CreateOssVpcResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOssVpcResponse represents the response of the api CreateOssVpc.
type CreateOssVpcResponse struct {
	RequestId string `json:"requestId" xml:"requestId"`
	Vip       string `json:"vip" xml:"vip"`
	VpcId     string `json:"vpcId" xml:"vpcId"`
}

// DeleteImgVpc version 2015-05-20
//
// required parameters:
//  name: Label, type: string
//  name: Region, type: string
//  name: VpcId, type: string
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DeleteImgVpc(Label, Region, VpcId string, optional openapi.M) (*DeleteImgVpcResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DeleteImgVpc")
	args.Query.Set("Label", Label)
	args.Query.Set("Region", Region)
	args.Query.Set("VpcId", VpcId)
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(DeleteImgVpcResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteImgVpcResponse represents the response of the api DeleteImgVpc.
type DeleteImgVpcResponse struct {
	RequestId string `json:"requestId" xml:"requestId"`
}

// DeleteOssVpc version 2015-05-20
//
// required parameters:
//  name: Label, type: string
//  name: Region, type: string
//  name: VpcId, type: string
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DeleteOssVpc(Label, Region, VpcId string, optional openapi.M) (*DeleteOssVpcResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DeleteOssVpc")
	args.Query.Set("Label", Label)
	args.Query.Set("Region", Region)
	args.Query.Set("VpcId", VpcId)
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(DeleteOssVpcResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteOssVpcResponse represents the response of the api DeleteOssVpc.
type DeleteOssVpcResponse struct {
	RequestId string `json:"requestId" xml:"requestId"`
}

// GetBucketVips version 2015-05-20
//
// required parameters:
//  name: BucketName, type: string
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) GetBucketVips(BucketName string, optional openapi.M) (*GetBucketVipsResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "GetBucketVips")
	args.Query.Set("BucketName", BucketName)
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(GetBucketVipsResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBucketVipsResponse represents the response of the api GetBucketVips.
type GetBucketVipsResponse struct {
	RequestId string `json:"requestId" xml:"requestId"`
}

// GetImgVpcInfo version 2015-05-20
//
// required parameters:
//  name: Label, type: string
//  name: Region, type: string
//  name: VpcId, type: string
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) GetImgVpcInfo(Label, Region, VpcId string, optional openapi.M) (*GetImgVpcInfoResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "GetImgVpcInfo")
	args.Query.Set("Label", Label)
	args.Query.Set("Region", Region)
	args.Query.Set("VpcId", VpcId)
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(GetImgVpcInfoResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// GetImgVpcInfoResponse represents the response of the api GetImgVpcInfo.
type GetImgVpcInfoResponse struct {
	RequestId string `json:"requestId" xml:"requestId"`
	Vip       string `json:"vip" xml:"vip"`
	VpcId     string `json:"vpcId" xml:"vpcId"`
}

// GetOssVpcInfo version 2015-05-20
//
// required parameters:
//  name: Label, type: string
//  name: Region, type: string
//  name: VpcId, type: string
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) GetOssVpcInfo(Label, Region, VpcId string, optional openapi.M) (*GetOssVpcInfoResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "GetOssVpcInfo")
	args.Query.Set("Label", Label)
	args.Query.Set("Region", Region)
	args.Query.Set("VpcId", VpcId)
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(GetOssVpcInfoResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// GetOssVpcInfoResponse represents the response of the api GetOssVpcInfo.
type GetOssVpcInfoResponse struct {
	RequestId string `json:"requestId" xml:"requestId"`
	Vip       string `json:"vip" xml:"vip"`
	VpcId     string `json:"vpcId" xml:"vpcId"`
}

// UnBindBucketVip version 2015-05-20
//
// required parameters:
//  name: BucketName, type: string
//  name: Region, type: string
//  name: Vip, type: string
//  name: VpcId, type: string
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) UnBindBucketVip(BucketName, Region, Vip, VpcId string, optional openapi.M) (*UnBindBucketVipResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "UnBindBucketVip")
	args.Query.Set("BucketName", BucketName)
	args.Query.Set("Region", Region)
	args.Query.Set("Vip", Vip)
	args.Query.Set("VpcId", VpcId)
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(UnBindBucketVipResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// UnBindBucketVipResponse represents the response of the api UnBindBucketVip.
type UnBindBucketVipResponse struct {
	RequestId string `json:"requestId" xml:"requestId"`
}