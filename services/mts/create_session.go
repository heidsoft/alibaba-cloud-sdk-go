package mts

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

func (client *Client) CreateSession(request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
	response = CreateCreateSessionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateSessionWithChan(request *CreateSessionRequest) (<-chan *CreateSessionResponse, <-chan error) {
	responseChan := make(chan *CreateSessionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSession(request)
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

func (client *Client) CreateSessionWithCallback(request *CreateSessionRequest, callback func(response *CreateSessionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSessionResponse
		var err error
		defer close(result)
		response, err = client.CreateSession(request)
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

type CreateSessionRequest struct {
	*requests.RpcRequest
	SessionTime          requests.Integer `position:"Query" name:"SessionTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	EndUserId            string           `position:"Query" name:"EndUserId"`
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
}

type CreateSessionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Session   string `json:"Session" xml:"Session"`
}

func CreateCreateSessionRequest() (request *CreateSessionRequest) {
	request = &CreateSessionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "CreateSession", "", "")
	return
}

func CreateCreateSessionResponse() (response *CreateSessionResponse) {
	response = &CreateSessionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}