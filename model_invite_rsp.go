/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wework
// InviteRsp struct for InviteRsp
type InviteRsp struct {
	Errcode int32 `json:"errcode,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	Invaliduser []string `json:"invaliduser,omitempty"`
	Invalidparty []string `json:"invalidparty,omitempty"`
	Invalidtag []string `json:"invalidtag,omitempty"`
}
