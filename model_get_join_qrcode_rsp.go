/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wework
// GetJoinQrcodeRsp struct for GetJoinQrcodeRsp
type GetJoinQrcodeRsp struct {
	Errcode int32 `json:"errcode,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	JoinQrcode string `json:"join_qrcode,omitempty"`
}
