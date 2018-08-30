package gitlab_resp_body_parser

import "time"

type Groups []struct {
	ID                   int         `json:"id"`
	Name                 string      `json:"name"`
	Path                 string      `json:"path"`
	Description          string      `json:"description"`
	Visibility           string      `json:"visibility"`
	LfsEnabled           bool        `json:"lfs_enabled"`
	AvatarURL            interface{} `json:"avatar_url"`
	WebURL               string      `json:"web_url"`
	RequestAccessEnabled bool        `json:"request_access_enabled"`
	FullName             string      `json:"full_name"`
	FullPath             string      `json:"full_path"`
	ParentID             interface{} `json:"parent_id"`
}

type Projects []struct {
	ID                       int           `json:"id"`
	Description              string        `json:"description"`
	DefaultBranch            interface{}   `json:"default_branch"`
	TagList                  []interface{} `json:"tag_list"`
	Archived                 bool          `json:"archived"`
	Visibility               string        `json:"visibility"`
	SSHURLToRepo             string        `json:"ssh_url_to_repo"`
	HTTPURLToRepo            string        `json:"http_url_to_repo"`
	WebURL                   string        `json:"web_url"`
	Name                     string        `json:"name"`
	NameWithNamespace        string        `json:"name_with_namespace"`
	Path                     string        `json:"path"`
	PathWithNamespace        string        `json:"path_with_namespace"`
	ContainerRegistryEnabled bool          `json:"container_registry_enabled"`
	IssuesEnabled            bool          `json:"issues_enabled"`
	MergeRequestsEnabled     bool          `json:"merge_requests_enabled"`
	WikiEnabled              bool          `json:"wiki_enabled"`
	JobsEnabled              bool          `json:"jobs_enabled"`
	SnippetsEnabled          bool          `json:"snippets_enabled"`
	CreatedAt                time.Time     `json:"created_at"`
	LastActivityAt           time.Time     `json:"last_activity_at"`
	SharedRunnersEnabled     bool          `json:"shared_runners_enabled"`
	LfsEnabled               bool          `json:"lfs_enabled"`
	CreatorID                int           `json:"creator_id"`
	Namespace                struct {
		ID       int         `json:"id"`
		Name     string      `json:"name"`
		Path     string      `json:"path"`
		Kind     string      `json:"kind"`
		FullPath string      `json:"full_path"`
		ParentID interface{} `json:"parent_id"`
	} `json:"namespace"`
	ImportStatus                              string        `json:"import_status"`
	AvatarURL                                 interface{}   `json:"avatar_url"`
	StarCount                                 int           `json:"star_count"`
	ForksCount                                int           `json:"forks_count"`
	OpenIssuesCount                           int           `json:"open_issues_count,omitempty"`
	PublicJobs                                bool          `json:"public_jobs"`
	CiConfigPath                              interface{}   `json:"ci_config_path"`
	SharedWithGroups                          []interface{} `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds          bool          `json:"only_allow_merge_if_pipeline_succeeds"`
	RequestAccessEnabled                      bool          `json:"request_access_enabled"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool          `json:"only_allow_merge_if_all_discussions_are_resolved"`
	PrintingMergeRequestLinkEnabled           bool          `json:"printing_merge_request_link_enabled"`
}
