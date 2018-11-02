/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Metadata for any related URL information.
type V1beta1RelatedUrl struct {

	// Specific URL associated with the resource.
	Url string `json:"url,omitempty"`

	// Label to describe usage of the URL.
	Label string `json:"label,omitempty"`
}
