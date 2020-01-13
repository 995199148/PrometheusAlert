package aegis

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

// DescribeAnalysisCurve invokes the aegis.DescribeAnalysisCurve API synchronously
// api document: https://help.aliyun.com/api/aegis/describeanalysiscurve.html
func (client *Client) DescribeAnalysisCurve(request *DescribeAnalysisCurveRequest) (response *DescribeAnalysisCurveResponse, err error) {
	response = CreateDescribeAnalysisCurveResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAnalysisCurveWithChan invokes the aegis.DescribeAnalysisCurve API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeanalysiscurve.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAnalysisCurveWithChan(request *DescribeAnalysisCurveRequest) (<-chan *DescribeAnalysisCurveResponse, <-chan error) {
	responseChan := make(chan *DescribeAnalysisCurveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAnalysisCurve(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeAnalysisCurveWithCallback invokes the aegis.DescribeAnalysisCurve API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeanalysiscurve.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAnalysisCurveWithCallback(request *DescribeAnalysisCurveRequest, callback func(response *DescribeAnalysisCurveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAnalysisCurveResponse
		var err error
		defer close(result)
		response, err = client.DescribeAnalysisCurve(request)
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

// DescribeAnalysisCurveRequest is the request struct for api DescribeAnalysisCurve
type DescribeAnalysisCurveRequest struct {
	*requests.RpcRequest
	SourceIp       string           `position:"Query" name:"SourceIp"`
	StartTimeStamp requests.Integer `position:"Query" name:"StartTimeStamp"`
	EndTimeStamp   requests.Integer `position:"Query" name:"EndTimeStamp"`
}

// DescribeAnalysisCurveResponse is the response struct for api DescribeAnalysisCurve
type DescribeAnalysisCurveResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Curves    Curves `json:"Curves" xml:"Curves"`
}

// CreateDescribeAnalysisCurveRequest creates a request to invoke DescribeAnalysisCurve API
func CreateDescribeAnalysisCurveRequest() (request *DescribeAnalysisCurveRequest) {
	request = &DescribeAnalysisCurveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeAnalysisCurve", "vipaegis", "openAPI")
	return
}

// CreateDescribeAnalysisCurveResponse creates a response to parse from DescribeAnalysisCurve response
func CreateDescribeAnalysisCurveResponse() (response *DescribeAnalysisCurveResponse) {
	response = &DescribeAnalysisCurveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}