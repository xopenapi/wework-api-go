/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wework
// TaskcardMsg struct for TaskcardMsg
type TaskcardMsg struct {
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
	TaskId string `json:"task_id,omitempty"`
	Btn []Btn `json:"btn,omitempty"`
}
