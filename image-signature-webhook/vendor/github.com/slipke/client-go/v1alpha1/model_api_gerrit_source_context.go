/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// A SourceContext referring to a Gerrit project.
type ApiGerritSourceContext struct {
	// The URI of a running Gerrit instance.
	HostUri string `json:"host_uri,omitempty"`
	// The full project name within the host. Projects may be nested, so \"project/subproject\" is a valid project name. The \"repo name\" is the hostURI/project.
	GerritProject string `json:"gerrit_project,omitempty"`
	// A revision (commit) ID.
	RevisionId string `json:"revision_id,omitempty"`
	// An alias, which may be a branch or tag.
	AliasContext *ApiAliasContext `json:"alias_context,omitempty"`
}
