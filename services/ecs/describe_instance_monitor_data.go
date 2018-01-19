package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeInstanceMonitorData(request *DescribeInstanceMonitorDataRequest) (response *DescribeInstanceMonitorDataResponse, err error) {
	response = CreateDescribeInstanceMonitorDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeInstanceMonitorDataWithChan(request *DescribeInstanceMonitorDataRequest) (<-chan *DescribeInstanceMonitorDataResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceMonitorDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceMonitorData(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeInstanceMonitorDataWithCallback(request *DescribeInstanceMonitorDataRequest, callback func(response *DescribeInstanceMonitorDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceMonitorDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceMonitorData(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeInstanceMonitorDataRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Period               requests.Integer `position:"Query" name:"Period"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	StartTime            string           `position:"Query" name:"StartTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeInstanceMonitorDataResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	MonitorData struct {
		InstanceMonitorData []struct {
			InstanceId        string  `json:"InstanceId" xml:"InstanceId"`
			CPU               int     `json:"CPU" xml:"CPU"`
			IntranetRX        int     `json:"IntranetRX" xml:"IntranetRX"`
			IntranetTX        int     `json:"IntranetTX" xml:"IntranetTX"`
			IntranetBandwidth int     `json:"IntranetBandwidth" xml:"IntranetBandwidth"`
			InternetRX        int     `json:"InternetRX" xml:"InternetRX"`
			InternetTX        int     `json:"InternetTX" xml:"InternetTX"`
			InternetBandwidth int     `json:"InternetBandwidth" xml:"InternetBandwidth"`
			IOPSRead          int     `json:"IOPSRead" xml:"IOPSRead"`
			IOPSWrite         int     `json:"IOPSWrite" xml:"IOPSWrite"`
			BPSRead           int     `json:"BPSRead" xml:"BPSRead"`
			BPSWrite          int     `json:"BPSWrite" xml:"BPSWrite"`
			CPUCreditUsage    float64 `json:"CPUCreditUsage" xml:"CPUCreditUsage"`
			CPUCreditBalance  float64 `json:"CPUCreditBalance" xml:"CPUCreditBalance"`
			TimeStamp         string  `json:"TimeStamp" xml:"TimeStamp"`
		} `json:"InstanceMonitorData" xml:"InstanceMonitorData"`
	} `json:"MonitorData" xml:"MonitorData"`
}

func CreateDescribeInstanceMonitorDataRequest() (request *DescribeInstanceMonitorDataRequest) {
	request = &DescribeInstanceMonitorDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceMonitorData", "ecs", "openAPI")
	return
}

func CreateDescribeInstanceMonitorDataResponse() (response *DescribeInstanceMonitorDataResponse) {
	response = &DescribeInstanceMonitorDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
