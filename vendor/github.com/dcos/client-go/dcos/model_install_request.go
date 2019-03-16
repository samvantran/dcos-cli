/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type InstallRequest struct {
	AppId          string                            `json:"appId,omitempty"`
	Options        map[string]map[string]interface{} `json:"options,omitempty"`
	PackageName    string                            `json:"packageName"`
	PackageVersion string                            `json:"packageVersion,omitempty"`
}
