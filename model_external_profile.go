/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wework
// ExternalProfile ExternalProfile
type ExternalProfile struct {
	ExternalCorpName string `json:"external_corp_name,omitempty"`
	ExternalAttr []ExtAttr `json:"external_attr,omitempty"`
}
