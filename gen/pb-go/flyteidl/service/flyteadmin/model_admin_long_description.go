/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Full user description with formatting preserved. This can be rendered by clients, such as the console or command line tools with in-tact formatting.
type AdminLongDescription struct {
	Values string `json:"values,omitempty"`
	Uri string `json:"uri,omitempty"`
	LongFormat *LongDescriptionDescriptionFormat `json:"long_format,omitempty"`
	IconLink string `json:"icon_link,omitempty"`
}
