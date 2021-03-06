// Copyright 2015 Chen Xianren. All rights reserved.
// Code generated by openapi-meta; DO NOT EDIT

package drds // import "github.com/cxr29/aliyun-openapi-go-sdk/apis/drds/v20150413"

import (
	"errors"
	"fmt"

	"github.com/cxr29/aliyun-openapi-go-sdk"
)

var (
	_ = errors.New("")
	_ = fmt.Sprint("")
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
	Product = "Drds"
	Style   = "RPC"
	Version = "2015-04-13"
)

// AlterTable version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DdlSql, type: string
//  name: DrdsInstanceId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) AlterTable(DbName, DdlSql, DrdsInstanceId string, optional openapi.M) (*AlterTableResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "AlterTable")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DdlSql", DdlSql)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
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

	result := new(AlterTableResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// AlterTableResponse represents the response of the api AlterTable.
type AlterTableResponse struct {
	TaskId string
}

// CancelDDLTask version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DrdsInstanceId, type: string
//  name: TaskId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) CancelDDLTask(DbName, DrdsInstanceId, TaskId string, optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "CancelDDLTask")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
	args.Query.Set("TaskId", TaskId)
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

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDrdsDB version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DrdsInstanceId, type: string
//  name: Encode, type: string, optional values: gbk|latin1|utf8|utf8mb4
//  name: Password, type: string
//  name: RdsInstances, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) CreateDrdsDB(DbName, DrdsInstanceId, Encode, Password, RdsInstances string, optional openapi.M) (*CreateDrdsDBResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "CreateDrdsDB")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
	args.Query.Set("Password", Password)
	args.Query.Set("RdsInstances", RdsInstances)
	if !openapi.IsIn(Encode, "gbk|latin1|utf8|utf8mb4") {
		return nil, errors.New("Encode must be gbk|latin1|utf8|utf8mb4")
	}

	args.Query.Set("Encode", Encode)
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

	result := new(CreateDrdsDBResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDrdsDBResponse represents the response of the api CreateDrdsDB.
type CreateDrdsDBResponse struct {
	Success bool
}

// CreateDrdsInstance version 2015-04-13
//
// required parameters:
//  name: PayType, type: string
//  name: Quantity, type: int
//  name: Specification, type: string
//  name: Type, type: string, optional values: 0|1|PRIVATE|PUBLIC
//  name: ZoneId, type: string
//
// optional parameters:
//  name: Description, type: string
//  name: VpcId, type: string
//  name: VswitchId, type: string
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) CreateDrdsInstance(PayType string, Quantity int, Specification, Type, ZoneId string, optional openapi.M) (*CreateDrdsInstanceResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "CreateDrdsInstance")
	args.Query.Set("PayType", PayType)
	args.Query.Set("Quantity", fmt.Sprint(Quantity))
	args.Query.Set("Specification", Specification)
	args.Query.Set("ZoneId", ZoneId)
	if !openapi.IsIn(Type, "0|1|PRIVATE|PUBLIC") {
		return nil, errors.New("Type must be 0|1|PRIVATE|PUBLIC")
	}

	args.Query.Set("Type", Type)
	if v, ok := optional["Description"]; ok {
		if Description, ok := v.(string); ok {
			args.Query.Set("Description", Description)
		} else {
			return nil, errors.New("Description must be string")
		}
	}
	if v, ok := optional["VpcId"]; ok {
		if VpcId, ok := v.(string); ok {
			args.Query.Set("VpcId", VpcId)
		} else {
			return nil, errors.New("VpcId must be string")
		}
	}
	if v, ok := optional["VswitchId"]; ok {
		if VswitchId, ok := v.(string); ok {
			args.Query.Set("VswitchId", VswitchId)
		} else {
			return nil, errors.New("VswitchId must be string")
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

	result := new(CreateDrdsInstanceResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDrdsInstanceResponse represents the response of the api CreateDrdsInstance.
type CreateDrdsInstanceResponse struct {
	Data struct {
		OrderId int64
	}
	Success bool
}

// CreateIndex version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DdlSql, type: string
//  name: DrdsInstanceId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) CreateIndex(DbName, DdlSql, DrdsInstanceId string, optional openapi.M) (*CreateIndexResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "CreateIndex")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DdlSql", DdlSql)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
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

	result := new(CreateIndexResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIndexResponse represents the response of the api CreateIndex.
type CreateIndexResponse struct {
	TaskId string
}

// CreateTable version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DdlSql, type: string
//  name: DrdsInstanceId, type: string
//  name: ShardType, type: string
//
// optional parameters:
//  name: AllowFullTableScan, type: string
//  name: ShardKey, type: string
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) CreateTable(DbName, DdlSql, DrdsInstanceId, ShardType string, optional openapi.M) (*CreateTableResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "CreateTable")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DdlSql", DdlSql)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
	args.Query.Set("ShardType", ShardType)
	if v, ok := optional["AllowFullTableScan"]; ok {
		if AllowFullTableScan, ok := v.(string); ok {
			args.Query.Set("AllowFullTableScan", AllowFullTableScan)
		} else {
			return nil, errors.New("AllowFullTableScan must be string")
		}
	}
	if v, ok := optional["ShardKey"]; ok {
		if ShardKey, ok := v.(string); ok {
			args.Query.Set("ShardKey", ShardKey)
		} else {
			return nil, errors.New("ShardKey must be string")
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

	result := new(CreateTableResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTableResponse represents the response of the api CreateTable.
type CreateTableResponse struct {
	TaskId string
}

// DeleteDrdsDB version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DrdsInstanceId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DeleteDrdsDB(DbName, DrdsInstanceId string, optional openapi.M) (*DeleteDrdsDBResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DeleteDrdsDB")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
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

	result := new(DeleteDrdsDBResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteDrdsDBResponse represents the response of the api DeleteDrdsDB.
type DeleteDrdsDBResponse struct {
	Success bool
}

// DescribeCreateDrdsInstanceStatus version 2015-04-13
//
// required parameters:
//  name: DrdsInstanceId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DescribeCreateDrdsInstanceStatus(DrdsInstanceId string, optional openapi.M) (*DescribeCreateDrdsInstanceStatusResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DescribeCreateDrdsInstanceStatus")
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
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

	result := new(DescribeCreateDrdsInstanceStatusResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeCreateDrdsInstanceStatusResponse represents the response of the api DescribeCreateDrdsInstanceStatus.
type DescribeCreateDrdsInstanceStatusResponse struct {
	Data struct {
		Status string
	}
	Success bool
}

// DescribeDDLTask version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DrdsInstanceId, type: string
//  name: TaskId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DescribeDDLTask(DbName, DrdsInstanceId, TaskId string, optional openapi.M) (*DescribeDDLTaskResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DescribeDDLTask")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
	args.Query.Set("TaskId", TaskId)
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

	result := new(DescribeDDLTaskResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDDLTaskResponse represents the response of the api DescribeDDLTask.
type DescribeDDLTaskResponse struct {
	Data struct {
		AllowCancel string
		ErrMsg      string
		GmtCreate   int64
		RequestId   string
		TargetId    string
		TaskDetail  string
		TaskName    string
		TaskPhase   string
		TaskStatus  int
		TaskType    int
	}
}

// DescribeDrdsDB version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DrdsInstanceId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DescribeDrdsDB(DbName, DrdsInstanceId string, optional openapi.M) (*DescribeDrdsDBResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DescribeDrdsDB")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
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

	result := new(DescribeDrdsDBResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDrdsDBResponse represents the response of the api DescribeDrdsDB.
type DescribeDrdsDBResponse struct {
	Data struct {
		CreateTime string
		DbName     string
		Mode       string
		Msg        string
		Status     int
	}
	Success bool
}

// DescribeDrdsDBIpWhiteList version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DrdsInstanceId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DescribeDrdsDBIpWhiteList(DbName, DrdsInstanceId string, optional openapi.M) (*DescribeDrdsDBIpWhiteListResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DescribeDrdsDBIpWhiteList")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
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

	result := new(DescribeDrdsDBIpWhiteListResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDrdsDBIpWhiteListResponse represents the response of the api DescribeDrdsDBIpWhiteList.
type DescribeDrdsDBIpWhiteListResponse struct {
	Data struct {
	}
	Success bool
}

// DescribeDrdsDBs version 2015-04-13
//
// required parameters:
//  name: DrdsInstanceId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DescribeDrdsDBs(DrdsInstanceId string, optional openapi.M) (*DescribeDrdsDBsResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DescribeDrdsDBs")
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
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

	result := new(DescribeDrdsDBsResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDrdsDBsResponse represents the response of the api DescribeDrdsDBs.
type DescribeDrdsDBsResponse struct {
	Data struct {
		Db []struct {
			CreateTime string
			DbName     string
			Mode       string
			Msg        string
			Status     int
		}
	}
	Success bool
}

// DescribeDrdsInstance version 2015-04-13
//
// required parameters:
//  name: DrdsInstanceId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DescribeDrdsInstance(DrdsInstanceId string, optional openapi.M) (*DescribeDrdsInstanceResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DescribeDrdsInstance")
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
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

	result := new(DescribeDrdsInstanceResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDrdsInstanceResponse represents the response of the api DescribeDrdsInstance.
type DescribeDrdsInstanceResponse struct {
	Data struct {
		CreateTime     int64
		Description    string
		DrdsInstanceId string
		NetworkType    string
		RegionId       string
		Status         string
		Type           string
		Version        int64
		Vips           struct {
			Vip []struct {
				IP   string
				Port string
				Type string
			}
		}
		ZoneId string
	}
	Success bool
}

// DescribeDrdsInstances version 2015-04-13
//
// optional parameters:
//  name: Type, type: string, optional values: 0|1|PRIVATE|PUBLIC
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DescribeDrdsInstances(optional openapi.M) (*DescribeDrdsInstancesResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DescribeDrdsInstances")
	if v, ok := optional["Type"]; ok {
		if Type, ok := v.(string); ok {
			if !openapi.IsIn(Type, "0|1|PRIVATE|PUBLIC") {
				return nil, errors.New("Type must be 0|1|PRIVATE|PUBLIC")
			}

			args.Query.Set("Type", Type)
		} else {
			return nil, errors.New("Type must be string")
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

	result := new(DescribeDrdsInstancesResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDrdsInstancesResponse represents the response of the api DescribeDrdsInstances.
type DescribeDrdsInstancesResponse struct {
	Data struct {
		Instance []struct {
			CreateTime     int64
			Description    string
			DrdsInstanceId string
			NetworkType    string
			RegionId       string
			Status         string
			Type           string
			Version        int64
			Vips           struct {
				Vip []struct {
					IP   string
					Port string
					Type string
				}
			}
			ZoneId string
		}
	}
	Success bool
}

// DropIndexes version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DrdsInstanceId, type: string
//  name: Indexes, type: string
//  name: Table, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DropIndexes(DbName, DrdsInstanceId, Indexes, Table string, optional openapi.M) (*DropIndexesResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DropIndexes")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
	args.Query.Set("Indexes", Indexes)
	args.Query.Set("Table", Table)
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

	result := new(DropIndexesResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DropIndexesResponse represents the response of the api DropIndexes.
type DropIndexesResponse struct {
	TaskId string
}

// DropTables version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DrdsInstanceId, type: string
//  name: Tables, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DropTables(DbName, DrdsInstanceId, Tables string, optional openapi.M) (*DropTablesResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DropTables")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
	args.Query.Set("Tables", Tables)
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

	result := new(DropTablesResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DropTablesResponse represents the response of the api DropTables.
type DropTablesResponse struct {
	TaskId string
}

// ListUnCompleteTasks version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DrdsInstanceId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) ListUnCompleteTasks(DbName, DrdsInstanceId string, optional openapi.M) (*ListUnCompleteTasksResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "ListUnCompleteTasks")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
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

	result := new(ListUnCompleteTasksResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// ListUnCompleteTasksResponse represents the response of the api ListUnCompleteTasks.
type ListUnCompleteTasksResponse struct {
	Data struct {
		Task []struct {
			AllowCancel string
			ErrMsg      string
			GmtCreate   int64 `json:"gmtCreate" xml:"gmtCreate"`
			RequestId   string
			TargetId    string
			TaskDetail  string
			TaskName    string
			TaskPhase   string
			TaskStatus  int `json:"taskStatus" xml:"taskStatus"`
			TaskType    int
		} `json:"task" xml:"task"`
	}
}

// ModifyDrdsDBPasswd version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DrdsInstanceId, type: string
//  name: NewPasswd, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) ModifyDrdsDBPasswd(DbName, DrdsInstanceId, NewPasswd string, optional openapi.M) (*ModifyDrdsDBPasswdResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "ModifyDrdsDBPasswd")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
	args.Query.Set("NewPasswd", NewPasswd)
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

	result := new(ModifyDrdsDBPasswdResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyDrdsDBPasswdResponse represents the response of the api ModifyDrdsDBPasswd.
type ModifyDrdsDBPasswdResponse struct {
	Success bool
}

// ModifyDrdsInstanceDescription version 2015-04-13
//
// required parameters:
//  name: Description, type: string
//  name: DrdsInstanceId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) ModifyDrdsInstanceDescription(Description, DrdsInstanceId string, optional openapi.M) (*ModifyDrdsInstanceDescriptionResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "ModifyDrdsInstanceDescription")
	args.Query.Set("Description", Description)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
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

	result := new(ModifyDrdsInstanceDescriptionResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyDrdsInstanceDescriptionResponse represents the response of the api ModifyDrdsInstanceDescription.
type ModifyDrdsInstanceDescriptionResponse struct {
	Success bool
}

// ModifyDrdsIpWhiteList version 2015-04-13
//
// required parameters:
//  name: DbName, type: string
//  name: DrdsInstanceId, type: string
//  name: IpWhiteList, type: string
//
// optional parameters:
//  name: Mode, type: bool
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) ModifyDrdsIpWhiteList(DbName, DrdsInstanceId, IpWhiteList string, optional openapi.M) (*ModifyDrdsIpWhiteListResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "ModifyDrdsIpWhiteList")
	args.Query.Set("DbName", DbName)
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
	args.Query.Set("IpWhiteList", IpWhiteList)
	if v, ok := optional["Mode"]; ok {
		if Mode, ok := v.(bool); ok {
			args.Query.Set("Mode", fmt.Sprint(Mode))
		} else {
			return nil, errors.New("Mode must be bool")
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

	result := new(ModifyDrdsIpWhiteListResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyDrdsIpWhiteListResponse represents the response of the api ModifyDrdsIpWhiteList.
type ModifyDrdsIpWhiteListResponse struct {
	Success bool
}

// RemoveDrdsInstance version 2015-04-13
//
// required parameters:
//  name: DrdsInstanceId, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) RemoveDrdsInstance(DrdsInstanceId string, optional openapi.M) (*RemoveDrdsInstanceResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "RemoveDrdsInstance")
	args.Query.Set("DrdsInstanceId", DrdsInstanceId)
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

	result := new(RemoveDrdsInstanceResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// RemoveDrdsInstanceResponse represents the response of the api RemoveDrdsInstance.
type RemoveDrdsInstanceResponse struct {
	Success bool
}
