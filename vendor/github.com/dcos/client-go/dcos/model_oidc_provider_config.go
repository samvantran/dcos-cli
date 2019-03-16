/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type OidcProviderConfig struct {
	Description             string `json:"description"`
	Issuer                  string `json:"issuer"`
	BaseUrl                 string `json:"base_url"`
	ClientSecret            string `json:"client_secret"`
	ClientId                string `json:"client_id"`
	VerifyServerCertificate bool   `json:"verify_server_certificate,omitempty"`
	CaCerts                 string `json:"ca_certs,omitempty"`
}
