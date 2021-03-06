# Go API client for wework

企业微信服务端API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./wework"
```

## Documentation for API Endpoints

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**GetJoinQrcode**](docs/DefaultApi.md#getjoinqrcode) | **Get** /corp/get_join_qrcode | 获取加入企业二维码
*MessageApi* | [**Send**](docs/MessageApi.md#send) | **Post** /message/send | 发送消息
*TokenApi* | [**GetApiDomainIp**](docs/TokenApi.md#getapidomainip) | **Get** /get_api_domain_ip | 获取企业微信API域名IP段
*TokenApi* | [**GetToken**](docs/TokenApi.md#gettoken) | **Get** /gettoken | 每个应用有独立的secret，获取到的access_token只能本应用使用，所以每个应用的access_token应该分开来获取
*UserApi* | [**Authsucc**](docs/UserApi.md#authsucc) | **Get** /user/authsucc | 二次验证
*UserApi* | [**Batchdelete**](docs/UserApi.md#batchdelete) | **Post** /user/batchdelete | 批量删除成员
*UserApi* | [**ConvertToOpenid**](docs/UserApi.md#converttoopenid) | **Post** /user/convert_to_openid | userid与openid互换
*UserApi* | [**Create**](docs/UserApi.md#create) | **Post** /user/create | 创建成员
*UserApi* | [**Delete**](docs/UserApi.md#delete) | **Get** /user/delete | 删除成员
*UserApi* | [**Get**](docs/UserApi.md#get) | **Get** /user/get | 读取成员
*UserApi* | [**Invite**](docs/UserApi.md#invite) | **Post** /batch/invite | 邀请成员
*UserApi* | [**List**](docs/UserApi.md#list) | **Get** /user/list | 获取部门成员详情
*UserApi* | [**Simplelist**](docs/UserApi.md#simplelist) | **Get** /user/simplelist | 获取部门成员
*UserApi* | [**Update**](docs/UserApi.md#update) | **Post** /user/update | 更新成员


## Documentation For Models

 - [ArticleMsg](docs/ArticleMsg.md)
 - [BaseResponse](docs/BaseResponse.md)
 - [BatchDeleteUserReq](docs/BatchDeleteUserReq.md)
 - [Btn](docs/Btn.md)
 - [ConvertToOpenidReq](docs/ConvertToOpenidReq.md)
 - [ConvertToOpenidRsp](docs/ConvertToOpenidRsp.md)
 - [ConvertToOpenidRspAllOf](docs/ConvertToOpenidRspAllOf.md)
 - [CreateUserReq](docs/CreateUserReq.md)
 - [ExtAttr](docs/ExtAttr.md)
 - [ExtAttrs](docs/ExtAttrs.md)
 - [ExternalProfile](docs/ExternalProfile.md)
 - [FileMsg](docs/FileMsg.md)
 - [GetApiDomainIpRsp](docs/GetApiDomainIpRsp.md)
 - [GetJoinQrcodeRsp](docs/GetJoinQrcodeRsp.md)
 - [GetJoinQrcodeRspAllOf](docs/GetJoinQrcodeRspAllOf.md)
 - [GetUserRsp](docs/GetUserRsp.md)
 - [ImageMsg](docs/ImageMsg.md)
 - [InviteReq](docs/InviteReq.md)
 - [InviteRsp](docs/InviteRsp.md)
 - [InviteRspAllOf](docs/InviteRspAllOf.md)
 - [ListUserRsp](docs/ListUserRsp.md)
 - [ListUserRspAllOf](docs/ListUserRspAllOf.md)
 - [MarkdownMsg](docs/MarkdownMsg.md)
 - [MiniProgramMsg](docs/MiniProgramMsg.md)
 - [MiniProgramMsgContentItem](docs/MiniProgramMsgContentItem.md)
 - [MpArticleMsg](docs/MpArticleMsg.md)
 - [MpNewsMsg](docs/MpNewsMsg.md)
 - [NewsMsg](docs/NewsMsg.md)
 - [SendMessageReq](docs/SendMessageReq.md)
 - [SendMessageRsp](docs/SendMessageRsp.md)
 - [SendMessageRspAllOf](docs/SendMessageRspAllOf.md)
 - [SimplelistRsp](docs/SimplelistRsp.md)
 - [TaskcardMsg](docs/TaskcardMsg.md)
 - [TextMsg](docs/TextMsg.md)
 - [TextcardMsg](docs/TextcardMsg.md)
 - [TokenRsp](docs/TokenRsp.md)
 - [TokenRspAllOf](docs/TokenRspAllOf.md)
 - [UpdateUserReq](docs/UpdateUserReq.md)
 - [User](docs/User.md)
 - [VideoMsg](docs/VideoMsg.md)
 - [VoiceMsg](docs/VoiceMsg.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



