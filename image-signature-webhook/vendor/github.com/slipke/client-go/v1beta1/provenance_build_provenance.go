/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

import (
	"time"
)

// Provenance of a build. Contains all information needed to verify the full details about the build from source to completion.
type ProvenanceBuildProvenance struct {

	// Required. Unique identifier of the build.
	Id string `json:"id,omitempty"`

	// ID of the project.
	ProjectId string `json:"project_id,omitempty"`

	// Commands requested by the build.
	Commands []ProvenanceCommand `json:"commands,omitempty"`

	// Output of the build.
	BuiltArtifacts []ProvenanceArtifact `json:"built_artifacts,omitempty"`

	// Time at which the build was created.
	CreateTime time.Time `json:"create_time,omitempty"`

	// Time at which execution of the build was started.
	StartTime time.Time `json:"start_time,omitempty"`

	// Time at which execution of the build was finished.
	EndTime time.Time `json:"end_time,omitempty"`

	// E-mail address of the user who initiated this build. Note that this was the user's e-mail address at the time the build was initiated; this address may not represent the same end-user for all time.
	Creator string `json:"creator,omitempty"`

	// URI where any logs for this provenance were written.
	LogsUri string `json:"logs_uri,omitempty"`

	// Details of the Source input to the build.
	SourceProvenance *ProvenanceSource `json:"source_provenance,omitempty"`

	// Trigger identifier if the build was triggered automatically; empty if not.
	TriggerId string `json:"trigger_id,omitempty"`

	// Special options applied to this build. This is a catch-all field where build providers can enter any desired additional details.
	BuildOptions map[string]string `json:"build_options,omitempty"`

	// Version string of the builder at the time this build was executed.
	BuilderVersion string `json:"builder_version,omitempty"`
}
