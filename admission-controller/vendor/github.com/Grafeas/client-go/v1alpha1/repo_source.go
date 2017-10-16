/* 
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * OpenAPI spec version: 0.1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

// RepoSource describes the location of the source in a Google Cloud Source Repository.
type RepoSource struct {

	// ID of the project that owns the repo.
	ProjectId string `json:"projectId,omitempty"`

	// Name of the repo.
	RepoName string `json:"repoName,omitempty"`

	// Name of the branch to build.
	BranchName string `json:"branchName,omitempty"`

	// Name of the tag to build.
	TagName string `json:"tagName,omitempty"`

	// Explicit commit SHA to build.
	CommitSha string `json:"commitSha,omitempty"`
}
