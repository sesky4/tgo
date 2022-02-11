// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20201002

import (
    "context"
    "github.com/sesky4/tgo/tencentcloud/common"
    tchttp "github.com/sesky4/tgo/tencentcloud/common/http"
    "github.com/sesky4/tgo/tencentcloud/common/profile"
)

const APIVersion = "2020-10-02"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewBatchSendEmailRequest() (request *BatchSendEmailRequest) {
    request = &BatchSendEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "BatchSendEmail")
    
    
    return
}

func NewBatchSendEmailResponse() (response *BatchSendEmailResponse) {
    response = &BatchSendEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchSendEmail
// 您可以通过此API批量发送TEXT或者HTML邮件，适用于营销类、通知类邮件。默认仅支持使用模板发送邮件，如需发送自定义内容，请单独联系商务开通此功能。批量发送之前，需先创建收件人列表，和收件人地址，并通过收件人列表id来进行发送。批量发送任务支持定时发送和周期重复发送，定时发送需传TimedParam，周期重复发送需传CycleParam
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BEGINTIMEBEFORENOW = "InvalidParameterValue.BeginTimeBeforeNow"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAERROR = "InvalidParameterValue.TemplateDataError"
//  MISSINGPARAMETER_CYCLEPARAMNECESSARY = "MissingParameter.CycleParamNecessary"
//  MISSINGPARAMETER_SENDPARAMNECESSARY = "MissingParameter.SendParamNecessary"
//  MISSINGPARAMETER_TIMEDPARAMNECESSARY = "MissingParameter.TimedParamNecessary"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
//  OPERATIONDENIED_RECEIVERSTATUSERROR = "OperationDenied.ReceiverStatusError"
//  OPERATIONDENIED_SENDADDRESSSTATUSERROR = "OperationDenied.SendAddressStatusError"
//  OPERATIONDENIED_TEMPLATESTATUSERROR = "OperationDenied.TemplateStatusError"
func (c *Client) BatchSendEmail(request *BatchSendEmailRequest) (response *BatchSendEmailResponse, err error) {
    if request == nil {
        request = NewBatchSendEmailRequest()
    }
    
    response = NewBatchSendEmailResponse()
    err = c.Send(request, response)
    return
}

// BatchSendEmail
// 您可以通过此API批量发送TEXT或者HTML邮件，适用于营销类、通知类邮件。默认仅支持使用模板发送邮件，如需发送自定义内容，请单独联系商务开通此功能。批量发送之前，需先创建收件人列表，和收件人地址，并通过收件人列表id来进行发送。批量发送任务支持定时发送和周期重复发送，定时发送需传TimedParam，周期重复发送需传CycleParam
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BEGINTIMEBEFORENOW = "InvalidParameterValue.BeginTimeBeforeNow"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_TEMPLATEDATAERROR = "InvalidParameterValue.TemplateDataError"
//  MISSINGPARAMETER_CYCLEPARAMNECESSARY = "MissingParameter.CycleParamNecessary"
//  MISSINGPARAMETER_SENDPARAMNECESSARY = "MissingParameter.SendParamNecessary"
//  MISSINGPARAMETER_TIMEDPARAMNECESSARY = "MissingParameter.TimedParamNecessary"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
//  OPERATIONDENIED_RECEIVERSTATUSERROR = "OperationDenied.ReceiverStatusError"
//  OPERATIONDENIED_SENDADDRESSSTATUSERROR = "OperationDenied.SendAddressStatusError"
//  OPERATIONDENIED_TEMPLATESTATUSERROR = "OperationDenied.TemplateStatusError"
func (c *Client) BatchSendEmailWithContext(ctx context.Context, request *BatchSendEmailRequest) (response *BatchSendEmailResponse, err error) {
    if request == nil {
        request = NewBatchSendEmailRequest()
    }
    request.SetContext(ctx)
    
    response = NewBatchSendEmailResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmailAddressRequest() (request *CreateEmailAddressRequest) {
    request = &CreateEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "CreateEmailAddress")
    
    
    return
}

func NewCreateEmailAddressResponse() (response *CreateEmailAddressResponse) {
    response = &CreateEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEmailAddress
// 在验证了发信域名之后，您需要一个发信地址来发送邮件。例如发信域名是mail.qcloud.com，那么发信地址可以为 service@mail.qcloud.com。如果您想要收件人在收件箱列表中显示您的别名，例如"腾讯云邮件通知"。那么发信地址为： 别名 空格 尖括号 邮箱地址。请注意中间需要有空格
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALEMAILADDRESS = "InvalidParameterValue.IllegalEmailAddress"
//  INVALIDPARAMETERVALUE_REPEATEMAILADDRESS = "InvalidParameterValue.RepeatEmailAddress"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOTVERIFIED = "OperationDenied.DomainNotVerified"
//  OPERATIONDENIED_EXCEEDSENDERLIMIT = "OperationDenied.ExceedSenderLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailAddress(request *CreateEmailAddressRequest) (response *CreateEmailAddressResponse, err error) {
    if request == nil {
        request = NewCreateEmailAddressRequest()
    }
    
    response = NewCreateEmailAddressResponse()
    err = c.Send(request, response)
    return
}

// CreateEmailAddress
// 在验证了发信域名之后，您需要一个发信地址来发送邮件。例如发信域名是mail.qcloud.com，那么发信地址可以为 service@mail.qcloud.com。如果您想要收件人在收件箱列表中显示您的别名，例如"腾讯云邮件通知"。那么发信地址为： 别名 空格 尖括号 邮箱地址。请注意中间需要有空格
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALEMAILADDRESS = "InvalidParameterValue.IllegalEmailAddress"
//  INVALIDPARAMETERVALUE_REPEATEMAILADDRESS = "InvalidParameterValue.RepeatEmailAddress"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOTVERIFIED = "OperationDenied.DomainNotVerified"
//  OPERATIONDENIED_EXCEEDSENDERLIMIT = "OperationDenied.ExceedSenderLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailAddressWithContext(ctx context.Context, request *CreateEmailAddressRequest) (response *CreateEmailAddressResponse, err error) {
    if request == nil {
        request = NewCreateEmailAddressRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmailIdentityRequest() (request *CreateEmailIdentityRequest) {
    request = &CreateEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "CreateEmailIdentity")
    
    
    return
}

func NewCreateEmailIdentityResponse() (response *CreateEmailIdentityResponse) {
    response = &CreateEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEmailIdentity
// 在使用身份发送电子邮件之前，您需要有一个电子邮件域名，该域名可以是您的网站或者移动应用的域名。您首先必须进行验证，证明自己是该域名的所有者，并且授权给腾讯云SES发送许可，才可以从该域名发送电子邮件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CREATEDBYOTHER = "InvalidParameterValue.CreatedByOther"
//  INVALIDPARAMETERVALUE_INVALIDEMAILIDENTITY = "InvalidParameterValue.InvalidEmailIdentity"
//  INVALIDPARAMETERVALUE_REPEATCREATION = "InvalidParameterValue.RepeatCreation"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_EXCEEDDOMAINLIMIT = "OperationDenied.ExceedDomainLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailIdentity(request *CreateEmailIdentityRequest) (response *CreateEmailIdentityResponse, err error) {
    if request == nil {
        request = NewCreateEmailIdentityRequest()
    }
    
    response = NewCreateEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

// CreateEmailIdentity
// 在使用身份发送电子邮件之前，您需要有一个电子邮件域名，该域名可以是您的网站或者移动应用的域名。您首先必须进行验证，证明自己是该域名的所有者，并且授权给腾讯云SES发送许可，才可以从该域名发送电子邮件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CREATEDBYOTHER = "InvalidParameterValue.CreatedByOther"
//  INVALIDPARAMETERVALUE_INVALIDEMAILIDENTITY = "InvalidParameterValue.InvalidEmailIdentity"
//  INVALIDPARAMETERVALUE_REPEATCREATION = "InvalidParameterValue.RepeatCreation"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_EXCEEDDOMAINLIMIT = "OperationDenied.ExceedDomainLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailIdentityWithContext(ctx context.Context, request *CreateEmailIdentityRequest) (response *CreateEmailIdentityResponse, err error) {
    if request == nil {
        request = NewCreateEmailIdentityRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmailTemplateRequest() (request *CreateEmailTemplateRequest) {
    request = &CreateEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "CreateEmailTemplate")
    
    
    return
}

func NewCreateEmailTemplateResponse() (response *CreateEmailTemplateResponse) {
    response = &CreateEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEmailTemplate
// 创建模板，该模板可以是TXT或者HTML，请注意如果HTML不要包含外部文件的CSS。模板中的变量使用 {{变量名}} 表示。
//
// 注意：模版需要审核通过才可以使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXCEEDTEMPLATELIMIT = "FailedOperation.ExceedTemplateLimit"
//  FAILEDOPERATION_TEMPLATECONTENTTOOLARGE = "FailedOperation.TemplateContentToolarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISNULL = "InvalidParameterValue.TemplateContentIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISWRONG = "InvalidParameterValue.TemplateContentIsWrong"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEILLEGAL = "InvalidParameterValue.TemplateNameIllegal"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailTemplate(request *CreateEmailTemplateRequest) (response *CreateEmailTemplateResponse, err error) {
    if request == nil {
        request = NewCreateEmailTemplateRequest()
    }
    
    response = NewCreateEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

// CreateEmailTemplate
// 创建模板，该模板可以是TXT或者HTML，请注意如果HTML不要包含外部文件的CSS。模板中的变量使用 {{变量名}} 表示。
//
// 注意：模版需要审核通过才可以使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXCEEDTEMPLATELIMIT = "FailedOperation.ExceedTemplateLimit"
//  FAILEDOPERATION_TEMPLATECONTENTTOOLARGE = "FailedOperation.TemplateContentToolarge"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISNULL = "InvalidParameterValue.TemplateContentIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISWRONG = "InvalidParameterValue.TemplateContentIsWrong"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEILLEGAL = "InvalidParameterValue.TemplateNameIllegal"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateEmailTemplateWithContext(ctx context.Context, request *CreateEmailTemplateRequest) (response *CreateEmailTemplateResponse, err error) {
    if request == nil {
        request = NewCreateEmailTemplateRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReceiverRequest() (request *CreateReceiverRequest) {
    request = &CreateReceiverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "CreateReceiver")
    
    
    return
}

func NewCreateReceiverResponse() (response *CreateReceiverResponse) {
    response = &CreateReceiverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateReceiver
// 创建收件人列表，收件人列表是发送批量邮件的目标邮件地址列表。创建列表后，需要上传收件人邮箱地址。之后创建发送任务，关联列表，便可以实现批量发送邮件的功能
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_RECEIVERDESCILLEGAL = "InvalidParameterValue.ReceiverDescIllegal"
//  INVALIDPARAMETERVALUE_RECEIVERNAMEILLEGAL = "InvalidParameterValue.ReceiverNameIllegal"
//  INVALIDPARAMETERVALUE_REPEATRECEIVERNAME = "InvalidParameterValue.RepeatReceiverName"
//  LIMITEXCEEDED_EXCEEDRECEIVERLIMIT = "LimitExceeded.ExceedReceiverLimit"
func (c *Client) CreateReceiver(request *CreateReceiverRequest) (response *CreateReceiverResponse, err error) {
    if request == nil {
        request = NewCreateReceiverRequest()
    }
    
    response = NewCreateReceiverResponse()
    err = c.Send(request, response)
    return
}

// CreateReceiver
// 创建收件人列表，收件人列表是发送批量邮件的目标邮件地址列表。创建列表后，需要上传收件人邮箱地址。之后创建发送任务，关联列表，便可以实现批量发送邮件的功能
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_RECEIVERDESCILLEGAL = "InvalidParameterValue.ReceiverDescIllegal"
//  INVALIDPARAMETERVALUE_RECEIVERNAMEILLEGAL = "InvalidParameterValue.ReceiverNameIllegal"
//  INVALIDPARAMETERVALUE_REPEATRECEIVERNAME = "InvalidParameterValue.RepeatReceiverName"
//  LIMITEXCEEDED_EXCEEDRECEIVERLIMIT = "LimitExceeded.ExceedReceiverLimit"
func (c *Client) CreateReceiverWithContext(ctx context.Context, request *CreateReceiverRequest) (response *CreateReceiverResponse, err error) {
    if request == nil {
        request = NewCreateReceiverRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateReceiverResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReceiverDetailRequest() (request *CreateReceiverDetailRequest) {
    request = &CreateReceiverDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "CreateReceiverDetail")
    
    
    return
}

func NewCreateReceiverDetailResponse() (response *CreateReceiverDetailResponse) {
    response = &CreateReceiverDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateReceiverDetail
// 在创建完收件人列表后，向这个收件人列表中批量增加收件人邮箱地址，一次最大支持10W，异步完成处理。收件人列表只可以上传一次，不可追加上传。数据量比较大的时候，上传可能需要一点时间，可以通过查询收件人列表了解上传状态和上传数量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_EXCEEDRECEIVERDETAILLIMIT = "LimitExceeded.ExceedReceiverDetailLimit"
//  MISSINGPARAMETER_EMAILSNECESSARY = "MissingParameter.EmailsNecessary"
//  MISSINGPARAMETER_RECEIVERIDNECESSARY = "MissingParameter.ReceiverIdNecessary"
//  OPERATIONDENIED_RECEIVERISOPERATING = "OperationDenied.ReceiverIsOperating"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
func (c *Client) CreateReceiverDetail(request *CreateReceiverDetailRequest) (response *CreateReceiverDetailResponse, err error) {
    if request == nil {
        request = NewCreateReceiverDetailRequest()
    }
    
    response = NewCreateReceiverDetailResponse()
    err = c.Send(request, response)
    return
}

// CreateReceiverDetail
// 在创建完收件人列表后，向这个收件人列表中批量增加收件人邮箱地址，一次最大支持10W，异步完成处理。收件人列表只可以上传一次，不可追加上传。数据量比较大的时候，上传可能需要一点时间，可以通过查询收件人列表了解上传状态和上传数量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_EXCEEDRECEIVERDETAILLIMIT = "LimitExceeded.ExceedReceiverDetailLimit"
//  MISSINGPARAMETER_EMAILSNECESSARY = "MissingParameter.EmailsNecessary"
//  MISSINGPARAMETER_RECEIVERIDNECESSARY = "MissingParameter.ReceiverIdNecessary"
//  OPERATIONDENIED_RECEIVERISOPERATING = "OperationDenied.ReceiverIsOperating"
//  OPERATIONDENIED_RECEIVERNOTEXIST = "OperationDenied.ReceiverNotExist"
func (c *Client) CreateReceiverDetailWithContext(ctx context.Context, request *CreateReceiverDetailRequest) (response *CreateReceiverDetailResponse, err error) {
    if request == nil {
        request = NewCreateReceiverDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateReceiverDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlackListRequest() (request *DeleteBlackListRequest) {
    request = &DeleteBlackListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "DeleteBlackList")
    
    
    return
}

func NewDeleteBlackListResponse() (response *DeleteBlackListResponse) {
    response = &DeleteBlackListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBlackList
// 邮箱被拉黑之后，用户如果确认收件邮箱有效或者已经处于激活状态，可以从腾讯云地址库中删除该黑名单之后继续投递。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteBlackList(request *DeleteBlackListRequest) (response *DeleteBlackListResponse, err error) {
    if request == nil {
        request = NewDeleteBlackListRequest()
    }
    
    response = NewDeleteBlackListResponse()
    err = c.Send(request, response)
    return
}

// DeleteBlackList
// 邮箱被拉黑之后，用户如果确认收件邮箱有效或者已经处于激活状态，可以从腾讯云地址库中删除该黑名单之后继续投递。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteBlackListWithContext(ctx context.Context, request *DeleteBlackListRequest) (response *DeleteBlackListResponse, err error) {
    if request == nil {
        request = NewDeleteBlackListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteBlackListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEmailAddressRequest() (request *DeleteEmailAddressRequest) {
    request = &DeleteEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "DeleteEmailAddress")
    
    
    return
}

func NewDeleteEmailAddressResponse() (response *DeleteEmailAddressResponse) {
    response = &DeleteEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEmailAddress
// 删除发信人地址
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOSUCHSENDER = "InvalidParameterValue.NoSuchSender"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailAddress(request *DeleteEmailAddressRequest) (response *DeleteEmailAddressResponse, err error) {
    if request == nil {
        request = NewDeleteEmailAddressRequest()
    }
    
    response = NewDeleteEmailAddressResponse()
    err = c.Send(request, response)
    return
}

// DeleteEmailAddress
// 删除发信人地址
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOSUCHSENDER = "InvalidParameterValue.NoSuchSender"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailAddressWithContext(ctx context.Context, request *DeleteEmailAddressRequest) (response *DeleteEmailAddressResponse, err error) {
    if request == nil {
        request = NewDeleteEmailAddressRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEmailIdentityRequest() (request *DeleteEmailIdentityRequest) {
    request = &DeleteEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "DeleteEmailIdentity")
    
    
    return
}

func NewDeleteEmailIdentityResponse() (response *DeleteEmailIdentityResponse) {
    response = &DeleteEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEmailIdentity
// 删除发信域名，删除后，将不可再使用该域名进行发信
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailIdentity(request *DeleteEmailIdentityRequest) (response *DeleteEmailIdentityResponse, err error) {
    if request == nil {
        request = NewDeleteEmailIdentityRequest()
    }
    
    response = NewDeleteEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

// DeleteEmailIdentity
// 删除发信域名，删除后，将不可再使用该域名进行发信
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailIdentityWithContext(ctx context.Context, request *DeleteEmailIdentityRequest) (response *DeleteEmailIdentityResponse, err error) {
    if request == nil {
        request = NewDeleteEmailIdentityRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEmailTemplateRequest() (request *DeleteEmailTemplateRequest) {
    request = &DeleteEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "DeleteEmailTemplate")
    
    
    return
}

func NewDeleteEmailTemplateResponse() (response *DeleteEmailTemplateResponse) {
    response = &DeleteEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEmailTemplate
// 删除发信模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailTemplate(request *DeleteEmailTemplateRequest) (response *DeleteEmailTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteEmailTemplateRequest()
    }
    
    response = NewDeleteEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

// DeleteEmailTemplate
// 删除发信模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteEmailTemplateWithContext(ctx context.Context, request *DeleteEmailTemplateRequest) (response *DeleteEmailTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteEmailTemplateRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewGetEmailIdentityRequest() (request *GetEmailIdentityRequest) {
    request = &GetEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "GetEmailIdentity")
    
    
    return
}

func NewGetEmailIdentityResponse() (response *GetEmailIdentityResponse) {
    response = &GetEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetEmailIdentity
// 获取某个发信域名的配置详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTEXISTDOMAIN = "InvalidParameterValue.NotExistDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) GetEmailIdentity(request *GetEmailIdentityRequest) (response *GetEmailIdentityResponse, err error) {
    if request == nil {
        request = NewGetEmailIdentityRequest()
    }
    
    response = NewGetEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

// GetEmailIdentity
// 获取某个发信域名的配置详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTEXISTDOMAIN = "InvalidParameterValue.NotExistDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) GetEmailIdentityWithContext(ctx context.Context, request *GetEmailIdentityRequest) (response *GetEmailIdentityResponse, err error) {
    if request == nil {
        request = NewGetEmailIdentityRequest()
    }
    request.SetContext(ctx)
    
    response = NewGetEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewGetEmailTemplateRequest() (request *GetEmailTemplateRequest) {
    request = &GetEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "GetEmailTemplate")
    
    
    return
}

func NewGetEmailTemplateResponse() (response *GetEmailTemplateResponse) {
    response = &GetEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetEmailTemplate
// 根据模板ID获取模板详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATENOTEXIST = "InvalidParameterValue.TemplateNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetEmailTemplate(request *GetEmailTemplateRequest) (response *GetEmailTemplateResponse, err error) {
    if request == nil {
        request = NewGetEmailTemplateRequest()
    }
    
    response = NewGetEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

// GetEmailTemplate
// 根据模板ID获取模板详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATENOTEXIST = "InvalidParameterValue.TemplateNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetEmailTemplateWithContext(ctx context.Context, request *GetEmailTemplateRequest) (response *GetEmailTemplateResponse, err error) {
    if request == nil {
        request = NewGetEmailTemplateRequest()
    }
    request.SetContext(ctx)
    
    response = NewGetEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewGetSendEmailStatusRequest() (request *GetSendEmailStatusRequest) {
    request = &GetSendEmailStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "GetSendEmailStatus")
    
    
    return
}

func NewGetSendEmailStatusResponse() (response *GetSendEmailStatusResponse) {
    response = &GetSendEmailStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetSendEmailStatus
// 获取邮件发送状态。仅支持查询30天之内的数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMAILADDRINBLACKLIST = "FailedOperation.EmailAddrInBlacklist"
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  FAILEDOPERATION_EXCEEDSENDLIMIT = "FailedOperation.ExceedSendLimit"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
//  FAILEDOPERATION_HIGHREJECTIONRATE = "FailedOperation.HighRejectionRate"
//  FAILEDOPERATION_INCORRECTEMAIL = "FailedOperation.IncorrectEmail"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_INSUFFICIENTQUOTA = "FailedOperation.InsufficientQuota"
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  FAILEDOPERATION_INVALIDTEMPLATEID = "FailedOperation.InvalidTemplateID"
//  FAILEDOPERATION_NOTAUTHENTICATEDSENDER = "FailedOperation.NotAuthenticatedSender"
//  FAILEDOPERATION_NOTSUPPORTDATE = "FailedOperation.NotSupportDate"
//  FAILEDOPERATION_PROTOCOLCHECKERR = "FailedOperation.ProtocolCheckErr"
//  FAILEDOPERATION_TEMPORARYBLOCKED = "FailedOperation.TemporaryBlocked"
//  FAILEDOPERATION_TOOMANYRECIPIENTS = "FailedOperation.TooManyRecipients"
//  FAILEDOPERATION_UNSUPPORTMAILTYPE = "FailedOperation.UnsupportMailType"
//  FAILEDOPERATION_WITHOUTPERMISSION = "FailedOperation.WithOutPermission"
//  FAILEDOPERATION_WRONGCONTENTJSON = "FailedOperation.WrongContentJson"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMAILADDRESSISNULL = "InvalidParameterValue.EmailAddressIsNULL"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSendEmailStatus(request *GetSendEmailStatusRequest) (response *GetSendEmailStatusResponse, err error) {
    if request == nil {
        request = NewGetSendEmailStatusRequest()
    }
    
    response = NewGetSendEmailStatusResponse()
    err = c.Send(request, response)
    return
}

// GetSendEmailStatus
// 获取邮件发送状态。仅支持查询30天之内的数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMAILADDRINBLACKLIST = "FailedOperation.EmailAddrInBlacklist"
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  FAILEDOPERATION_EXCEEDSENDLIMIT = "FailedOperation.ExceedSendLimit"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
//  FAILEDOPERATION_HIGHREJECTIONRATE = "FailedOperation.HighRejectionRate"
//  FAILEDOPERATION_INCORRECTEMAIL = "FailedOperation.IncorrectEmail"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_INSUFFICIENTQUOTA = "FailedOperation.InsufficientQuota"
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  FAILEDOPERATION_INVALIDTEMPLATEID = "FailedOperation.InvalidTemplateID"
//  FAILEDOPERATION_NOTAUTHENTICATEDSENDER = "FailedOperation.NotAuthenticatedSender"
//  FAILEDOPERATION_NOTSUPPORTDATE = "FailedOperation.NotSupportDate"
//  FAILEDOPERATION_PROTOCOLCHECKERR = "FailedOperation.ProtocolCheckErr"
//  FAILEDOPERATION_TEMPORARYBLOCKED = "FailedOperation.TemporaryBlocked"
//  FAILEDOPERATION_TOOMANYRECIPIENTS = "FailedOperation.TooManyRecipients"
//  FAILEDOPERATION_UNSUPPORTMAILTYPE = "FailedOperation.UnsupportMailType"
//  FAILEDOPERATION_WITHOUTPERMISSION = "FailedOperation.WithOutPermission"
//  FAILEDOPERATION_WRONGCONTENTJSON = "FailedOperation.WrongContentJson"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMAILADDRESSISNULL = "InvalidParameterValue.EmailAddressIsNULL"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSendEmailStatusWithContext(ctx context.Context, request *GetSendEmailStatusRequest) (response *GetSendEmailStatusResponse, err error) {
    if request == nil {
        request = NewGetSendEmailStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewGetSendEmailStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetStatisticsReportRequest() (request *GetStatisticsReportRequest) {
    request = &GetStatisticsReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "GetStatisticsReport")
    
    
    return
}

func NewGetStatisticsReportResponse() (response *GetStatisticsReportResponse) {
    response = &GetStatisticsReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetStatisticsReport
// 获取近期发送的统计情况，包含发送量、送达率、打开率、退信率等一系列数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) GetStatisticsReport(request *GetStatisticsReportRequest) (response *GetStatisticsReportResponse, err error) {
    if request == nil {
        request = NewGetStatisticsReportRequest()
    }
    
    response = NewGetStatisticsReportResponse()
    err = c.Send(request, response)
    return
}

// GetStatisticsReport
// 获取近期发送的统计情况，包含发送量、送达率、打开率、退信率等一系列数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) GetStatisticsReportWithContext(ctx context.Context, request *GetStatisticsReportRequest) (response *GetStatisticsReportResponse, err error) {
    if request == nil {
        request = NewGetStatisticsReportRequest()
    }
    request.SetContext(ctx)
    
    response = NewGetStatisticsReportResponse()
    err = c.Send(request, response)
    return
}

func NewListBlackEmailAddressRequest() (request *ListBlackEmailAddressRequest) {
    request = &ListBlackEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "ListBlackEmailAddress")
    
    
    return
}

func NewListBlackEmailAddressResponse() (response *ListBlackEmailAddressResponse) {
    response = &ListBlackEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListBlackEmailAddress
// 腾讯云发送的邮件一旦被收件方判断为硬退(Hard Bounce)，腾讯云会拉黑该地址，并不允许所有用户向该地址发送邮件。成为邮箱黑名单。如果业务方确认是误判，可以从黑名单中删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListBlackEmailAddress(request *ListBlackEmailAddressRequest) (response *ListBlackEmailAddressResponse, err error) {
    if request == nil {
        request = NewListBlackEmailAddressRequest()
    }
    
    response = NewListBlackEmailAddressResponse()
    err = c.Send(request, response)
    return
}

// ListBlackEmailAddress
// 腾讯云发送的邮件一旦被收件方判断为硬退(Hard Bounce)，腾讯云会拉黑该地址，并不允许所有用户向该地址发送邮件。成为邮箱黑名单。如果业务方确认是误判，可以从黑名单中删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_WRONGDATE = "InvalidParameterValue.WrongDate"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListBlackEmailAddressWithContext(ctx context.Context, request *ListBlackEmailAddressRequest) (response *ListBlackEmailAddressResponse, err error) {
    if request == nil {
        request = NewListBlackEmailAddressRequest()
    }
    request.SetContext(ctx)
    
    response = NewListBlackEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewListEmailAddressRequest() (request *ListEmailAddressRequest) {
    request = &ListEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "ListEmailAddress")
    
    
    return
}

func NewListEmailAddressResponse() (response *ListEmailAddressResponse) {
    response = &ListEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListEmailAddress
// 获取发信地址列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailAddress(request *ListEmailAddressRequest) (response *ListEmailAddressResponse, err error) {
    if request == nil {
        request = NewListEmailAddressRequest()
    }
    
    response = NewListEmailAddressResponse()
    err = c.Send(request, response)
    return
}

// ListEmailAddress
// 获取发信地址列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailAddressWithContext(ctx context.Context, request *ListEmailAddressRequest) (response *ListEmailAddressResponse, err error) {
    if request == nil {
        request = NewListEmailAddressRequest()
    }
    request.SetContext(ctx)
    
    response = NewListEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewListEmailIdentitiesRequest() (request *ListEmailIdentitiesRequest) {
    request = &ListEmailIdentitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "ListEmailIdentities")
    
    
    return
}

func NewListEmailIdentitiesResponse() (response *ListEmailIdentitiesResponse) {
    response = &ListEmailIdentitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListEmailIdentities
// 获取当前发信域名列表，包含已验证通过与未验证的域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailIdentities(request *ListEmailIdentitiesRequest) (response *ListEmailIdentitiesResponse, err error) {
    if request == nil {
        request = NewListEmailIdentitiesRequest()
    }
    
    response = NewListEmailIdentitiesResponse()
    err = c.Send(request, response)
    return
}

// ListEmailIdentities
// 获取当前发信域名列表，包含已验证通过与未验证的域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailIdentitiesWithContext(ctx context.Context, request *ListEmailIdentitiesRequest) (response *ListEmailIdentitiesResponse, err error) {
    if request == nil {
        request = NewListEmailIdentitiesRequest()
    }
    request.SetContext(ctx)
    
    response = NewListEmailIdentitiesResponse()
    err = c.Send(request, response)
    return
}

func NewListEmailTemplatesRequest() (request *ListEmailTemplatesRequest) {
    request = &ListEmailTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "ListEmailTemplates")
    
    
    return
}

func NewListEmailTemplatesResponse() (response *ListEmailTemplatesResponse) {
    response = &ListEmailTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListEmailTemplates
// 获取当前邮件模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailTemplates(request *ListEmailTemplatesRequest) (response *ListEmailTemplatesResponse, err error) {
    if request == nil {
        request = NewListEmailTemplatesRequest()
    }
    
    response = NewListEmailTemplatesResponse()
    err = c.Send(request, response)
    return
}

// ListEmailTemplates
// 获取当前邮件模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ListEmailTemplatesWithContext(ctx context.Context, request *ListEmailTemplatesRequest) (response *ListEmailTemplatesResponse, err error) {
    if request == nil {
        request = NewListEmailTemplatesRequest()
    }
    request.SetContext(ctx)
    
    response = NewListEmailTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewListReceiversRequest() (request *ListReceiversRequest) {
    request = &ListReceiversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "ListReceivers")
    
    
    return
}

func NewListReceiversResponse() (response *ListReceiversResponse) {
    response = &ListReceiversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListReceivers
// 根据条件查询收件人列表，支持分页，模糊查询，状态查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INTERNALERROR = "InternalError"
func (c *Client) ListReceivers(request *ListReceiversRequest) (response *ListReceiversResponse, err error) {
    if request == nil {
        request = NewListReceiversRequest()
    }
    
    response = NewListReceiversResponse()
    err = c.Send(request, response)
    return
}

// ListReceivers
// 根据条件查询收件人列表，支持分页，模糊查询，状态查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
//  INTERNALERROR = "InternalError"
func (c *Client) ListReceiversWithContext(ctx context.Context, request *ListReceiversRequest) (response *ListReceiversResponse, err error) {
    if request == nil {
        request = NewListReceiversRequest()
    }
    request.SetContext(ctx)
    
    response = NewListReceiversResponse()
    err = c.Send(request, response)
    return
}

func NewListSendTasksRequest() (request *ListSendTasksRequest) {
    request = &ListSendTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "ListSendTasks")
    
    
    return
}

func NewListSendTasksResponse() (response *ListSendTasksResponse) {
    response = &ListSendTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListSendTasks
// 分页查询批量发送邮件任务，包含即时发送任务，定时发送任务，周期重复发送任务，查询发送情况，包括请求数量，已发数量，缓存数量，任务状态等信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
func (c *Client) ListSendTasks(request *ListSendTasksRequest) (response *ListSendTasksResponse, err error) {
    if request == nil {
        request = NewListSendTasksRequest()
    }
    
    response = NewListSendTasksResponse()
    err = c.Send(request, response)
    return
}

// ListSendTasks
// 分页查询批量发送邮件任务，包含即时发送任务，定时发送任务，周期重复发送任务，查询发送情况，包括请求数量，已发数量，缓存数量，任务状态等信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDLIMIT = "FailedOperation.InvalidLimit"
func (c *Client) ListSendTasksWithContext(ctx context.Context, request *ListSendTasksRequest) (response *ListSendTasksResponse, err error) {
    if request == nil {
        request = NewListSendTasksRequest()
    }
    request.SetContext(ctx)
    
    response = NewListSendTasksResponse()
    err = c.Send(request, response)
    return
}

func NewSendEmailRequest() (request *SendEmailRequest) {
    request = &SendEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "SendEmail")
    
    
    return
}

func NewSendEmailResponse() (response *SendEmailResponse) {
    response = &SendEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendEmail
// 您可以通过此API发送TEXT或者HTML邮件，适用于触发类邮件（验证码、交易类）。默认仅支持使用模板发送邮件，如需发送自定义内容，请单独联系商务开通此功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ATTACHCONTENTTOOLARGE = "FailedOperation.AttachContentToolarge"
//  FAILEDOPERATION_EMAILADDRINBLACKLIST = "FailedOperation.EmailAddrInBlacklist"
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  FAILEDOPERATION_EXCEEDSENDLIMIT = "FailedOperation.ExceedSendLimit"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
//  FAILEDOPERATION_HIGHREJECTIONRATE = "FailedOperation.HighRejectionRate"
//  FAILEDOPERATION_INCORRECTEMAIL = "FailedOperation.IncorrectEmail"
//  FAILEDOPERATION_INCORRECTSENDER = "FailedOperation.IncorrectSender"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_INSUFFICIENTQUOTA = "FailedOperation.InsufficientQuota"
//  FAILEDOPERATION_INVALIDATTACHNAME = "FailedOperation.InvalidAttachName"
//  FAILEDOPERATION_INVALIDTEMPLATEID = "FailedOperation.InvalidTemplateID"
//  FAILEDOPERATION_MISSINGEMAILCONTENT = "FailedOperation.MissingEmailContent"
//  FAILEDOPERATION_NOATTACHPERMISSION = "FailedOperation.NoAttachPermission"
//  FAILEDOPERATION_NOTAUTHENTICATEDSENDER = "FailedOperation.NotAuthenticatedSender"
//  FAILEDOPERATION_PROTOCOLCHECKERR = "FailedOperation.ProtocolCheckErr"
//  FAILEDOPERATION_SENDEMAILERR = "FailedOperation.SendEmailErr"
//  FAILEDOPERATION_TEMPORARYBLOCKED = "FailedOperation.TemporaryBlocked"
//  FAILEDOPERATION_TOOMANYATTACHMENTS = "FailedOperation.TooManyAttachments"
//  FAILEDOPERATION_TOOMANYRECIPIENTS = "FailedOperation.TooManyRecipients"
//  FAILEDOPERATION_UNSUPPORTMAILTYPE = "FailedOperation.UnsupportMailType"
//  FAILEDOPERATION_WITHOUTPERMISSION = "FailedOperation.WithOutPermission"
//  FAILEDOPERATION_WRONGCONTENTJSON = "FailedOperation.WrongContentJson"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ATTACHCONTENTISWRONG = "InvalidParameterValue.AttachContentIsWrong"
//  INVALIDPARAMETERVALUE_EMAILADDRESSISNULL = "InvalidParameterValue.EmailAddressIsNULL"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_INVALIDEMAILIDENTITY = "InvalidParameterValue.InvalidEmailIdentity"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendEmail(request *SendEmailRequest) (response *SendEmailResponse, err error) {
    if request == nil {
        request = NewSendEmailRequest()
    }
    
    response = NewSendEmailResponse()
    err = c.Send(request, response)
    return
}

// SendEmail
// 您可以通过此API发送TEXT或者HTML邮件，适用于触发类邮件（验证码、交易类）。默认仅支持使用模板发送邮件，如需发送自定义内容，请单独联系商务开通此功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ATTACHCONTENTTOOLARGE = "FailedOperation.AttachContentToolarge"
//  FAILEDOPERATION_EMAILADDRINBLACKLIST = "FailedOperation.EmailAddrInBlacklist"
//  FAILEDOPERATION_EMAILCONTENTTOOLARGE = "FailedOperation.EmailContentToolarge"
//  FAILEDOPERATION_EXCEEDSENDLIMIT = "FailedOperation.ExceedSendLimit"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
//  FAILEDOPERATION_HIGHREJECTIONRATE = "FailedOperation.HighRejectionRate"
//  FAILEDOPERATION_INCORRECTEMAIL = "FailedOperation.IncorrectEmail"
//  FAILEDOPERATION_INCORRECTSENDER = "FailedOperation.IncorrectSender"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_INSUFFICIENTQUOTA = "FailedOperation.InsufficientQuota"
//  FAILEDOPERATION_INVALIDATTACHNAME = "FailedOperation.InvalidAttachName"
//  FAILEDOPERATION_INVALIDTEMPLATEID = "FailedOperation.InvalidTemplateID"
//  FAILEDOPERATION_MISSINGEMAILCONTENT = "FailedOperation.MissingEmailContent"
//  FAILEDOPERATION_NOATTACHPERMISSION = "FailedOperation.NoAttachPermission"
//  FAILEDOPERATION_NOTAUTHENTICATEDSENDER = "FailedOperation.NotAuthenticatedSender"
//  FAILEDOPERATION_PROTOCOLCHECKERR = "FailedOperation.ProtocolCheckErr"
//  FAILEDOPERATION_SENDEMAILERR = "FailedOperation.SendEmailErr"
//  FAILEDOPERATION_TEMPORARYBLOCKED = "FailedOperation.TemporaryBlocked"
//  FAILEDOPERATION_TOOMANYATTACHMENTS = "FailedOperation.TooManyAttachments"
//  FAILEDOPERATION_TOOMANYRECIPIENTS = "FailedOperation.TooManyRecipients"
//  FAILEDOPERATION_UNSUPPORTMAILTYPE = "FailedOperation.UnsupportMailType"
//  FAILEDOPERATION_WITHOUTPERMISSION = "FailedOperation.WithOutPermission"
//  FAILEDOPERATION_WRONGCONTENTJSON = "FailedOperation.WrongContentJson"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ATTACHCONTENTISWRONG = "InvalidParameterValue.AttachContentIsWrong"
//  INVALIDPARAMETERVALUE_EMAILADDRESSISNULL = "InvalidParameterValue.EmailAddressIsNULL"
//  INVALIDPARAMETERVALUE_EMAILCONTENTISWRONG = "InvalidParameterValue.EmailContentIsWrong"
//  INVALIDPARAMETERVALUE_INVALIDEMAILIDENTITY = "InvalidParameterValue.InvalidEmailIdentity"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendEmailWithContext(ctx context.Context, request *SendEmailRequest) (response *SendEmailResponse, err error) {
    if request == nil {
        request = NewSendEmailRequest()
    }
    request.SetContext(ctx)
    
    response = NewSendEmailResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEmailIdentityRequest() (request *UpdateEmailIdentityRequest) {
    request = &UpdateEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "UpdateEmailIdentity")
    
    
    return
}

func NewUpdateEmailIdentityResponse() (response *UpdateEmailIdentityResponse) {
    response = &UpdateEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateEmailIdentity
// 您已经成功配置好了您的DNS，接下来请求腾讯云验证您的DNS配置是否正确
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICENOTAVAILABLE = "FailedOperation.ServiceNotAvailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTEXISTDOMAIN = "InvalidParameterValue.NotExistDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) UpdateEmailIdentity(request *UpdateEmailIdentityRequest) (response *UpdateEmailIdentityResponse, err error) {
    if request == nil {
        request = NewUpdateEmailIdentityRequest()
    }
    
    response = NewUpdateEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

// UpdateEmailIdentity
// 您已经成功配置好了您的DNS，接下来请求腾讯云验证您的DNS配置是否正确
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICENOTAVAILABLE = "FailedOperation.ServiceNotAvailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTEXISTDOMAIN = "InvalidParameterValue.NotExistDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) UpdateEmailIdentityWithContext(ctx context.Context, request *UpdateEmailIdentityRequest) (response *UpdateEmailIdentityResponse, err error) {
    if request == nil {
        request = NewUpdateEmailIdentityRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEmailTemplateRequest() (request *UpdateEmailTemplateRequest) {
    request = &UpdateEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "UpdateEmailTemplate")
    
    
    return
}

func NewUpdateEmailTemplateResponse() (response *UpdateEmailTemplateResponse) {
    response = &UpdateEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateEmailTemplate
// 更新邮件模板，更新后需再次审核
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISNULL = "InvalidParameterValue.TemplateContentIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISWRONG = "InvalidParameterValue.TemplateContentIsWrong"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEILLEGAL = "InvalidParameterValue.TemplateNameIllegal"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEISNULL = "InvalidParameterValue.TemplateNameIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATENOTEXIST = "InvalidParameterValue.TemplateNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) UpdateEmailTemplate(request *UpdateEmailTemplateRequest) (response *UpdateEmailTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateEmailTemplateRequest()
    }
    
    response = NewUpdateEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

// UpdateEmailTemplate
// 更新邮件模板，更新后需再次审核
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISNULL = "InvalidParameterValue.TemplateContentIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATECONTENTISWRONG = "InvalidParameterValue.TemplateContentIsWrong"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEILLEGAL = "InvalidParameterValue.TemplateNameIllegal"
//  INVALIDPARAMETERVALUE_TEMPLATENAMEISNULL = "InvalidParameterValue.TemplateNameIsNULL"
//  INVALIDPARAMETERVALUE_TEMPLATENOTEXIST = "InvalidParameterValue.TemplateNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) UpdateEmailTemplateWithContext(ctx context.Context, request *UpdateEmailTemplateRequest) (response *UpdateEmailTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateEmailTemplateRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateEmailTemplateResponse()
    err = c.Send(request, response)
    return
}
