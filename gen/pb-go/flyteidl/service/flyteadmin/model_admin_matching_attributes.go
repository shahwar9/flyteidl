/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Generic container for encapsulating all types of the above attributes messages.
type AdminMatchingAttributes struct {
	TaskResourceAttributes *AdminTaskResourceAttributes `json:"task_resource_attributes,omitempty"`
	ClusterResourceAttributes *AdminClusterResourceAttributes `json:"cluster_resource_attributes,omitempty"`
	ExecutionQueueAttributes *AdminExecutionQueueAttributes `json:"execution_queue_attributes,omitempty"`
	ExecutionClusterLabel *AdminExecutionClusterLabel `json:"execution_cluster_label,omitempty"`
	QualityOfService *AdminQualityOfService `json:"quality_of_service,omitempty"`
}
