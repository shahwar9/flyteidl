/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type AdminQualityOfService struct {
	Tier *QualityOfServiceTier `json:"tier,omitempty"`
	Spec *AdminQualityOfServiceSpec `json:"spec,omitempty"`
}
