package main

// DeploymentsListResponse https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentsListResponse
type DeploymentsListResponse struct {
	Deployments []DeploymentsListingData `json:"deployments"`
}

// DeploymentsListingData https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentsListingData
type DeploymentsListingData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// DeploymentCreateRequest https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentCreateRequest
type DeploymentCreateRequest struct {
	Name      string                    `json:"name"`
	Resources DeploymentCreateResources `json:"resources"`
}

// DeploymentCreateResources https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentCreateResources
type DeploymentCreateResources struct {
	Elasticsearch []ElasticsearchPayload `json:"elasticsearch"`
	Kibana        []KibanaPayload        `json:"kibana"`
}

// ApmPayload https://www.elastic.co/guide/en/cloud/current/definitions.html#ApmPayload
type ApmPayload struct {
	DisplayName               string      `json:"display_name"`
	ElasticsearchClusterRefID string      `json:"elasticsearch_cluster_ref_id"`
	Plan                      ApmPlan     `json:"plan"`
	RefID                     string      `json:"ref_id"`
	Region                    string      `json:"region"`
	Settings                  ApmSettings `json:"settings"`
}

// ApmPlan https://www.elastic.co/guide/en/cloud/current/definitions.html#ApmPlan
type ApmPlan struct {
	Apm             ApmConfiguration              `json:"apm"`
	ClusterTopology []ApmTopologyElement          `json:"cluster_topology"`
	Transient       TransientApmPlanConfiguration `json:"transient"`
}

// ApmConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#ApmConfiguration
type ApmConfiguration struct {
	DockerImage              string            `json:"docker_image"`
	SystemSettings           ApmSystemSettings `json:"system_settings"`
	UserSettingsJSON         string            `json:"user_settings_json"`
	UserSettingsOverrideJSON string            `json:"user_settings_override_json"`
	UserSettingsOverrideYAML string            `json:"user_settings_override_yaml"`
	UserSettingsYAML         string            `json:"user_settings_yaml"`
	Version                  string            `json:"version"`
}

// ApmSystemSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#ApmSystemSettings
type ApmSystemSettings struct {
	DebugEnabled          bool   `json:"debug_enabled"`
	ElasticsearchPassword string `json:"elasticsearch_password"`
	ElasticsearchURL      string `json:"elasticsearch_url"`
	ElasticsearchUsername string `json:"elasticsearch_username"`
	KibanaURL             string `json:"kibana_url"`
	SecretToken           string `json:"secret_token"`
}

// ApmTopologyElement https://www.elastic.co/guide/en/cloud/current/definitions.html#ApmTopologyElement
type ApmTopologyElement struct {
	Apm                     ApmConfiguration `json:"apm"`
	InstanceConfigurationID string           `json:"instance_configuration_id"`
	Size                    TopologySize     `json:"size"`
	ZoneCount               int              `json:"zone_count"`
}

// TopologySize https://www.elastic.co/guide/en/cloud/current/definitions.html#TopologySize
type TopologySize struct {
	Resource string `json:"resource"`
	Value    int    `json:"value"`
}

// TransientApmPlanConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#TransientApmPlanConfiguration
type TransientApmPlanConfiguration struct {
	PlanConfiguration ApmPlanControlConfiguration `json:"plan_configuration"`
	Strategy          PlanStrategy                `json:"strategy"`
}

// ApmPlanControlConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#ApmPlanControlConfiguration
type ApmPlanControlConfiguration struct {
	CalmWaitTime        int    `json:"calm_wait_time"`
	ClusterReboot       string `json:"cluster_reboot"`
	ExtendedMaintenance bool   `json:"extended_maintenance"`
	Timeout             int    `json:"timeout"`
}

// PlanStrategy https://www.elastic.co/guide/en/cloud/current/definitions.html#PlanStrategy
type PlanStrategy struct {
	Autodetect           string                `json:"autodetect"`
	GrowAndShrink        string                `json:"grow_and_shrink"`
	Rolling              RollingStrategyConfig `json:"rolling"`
	RollingGrowAndShrink string                `json:"rolling_grow_and_shrink"`
}

// RollingStrategyConfig https://www.elastic.co/guide/en/cloud/current/definitions.html#RollingStrategyConfig
type RollingStrategyConfig struct {
	AllowInlineResize bool   `json:"allow_inline_resize"`
	GroupBy           string `json:"group_by"`
	ShardInitWaitTime int    `json:"shard_init_wait_time"`
	SkipSyncedFlush   bool   `json:"skip_synced_flush"`
}

// ApmSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#ApmSettings
type ApmSettings struct {
	Metadata ClusterMetadataSettings `json:"metadata"`
}

// ClusterMetadataSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#ClusterMetadataSettings
type ClusterMetadataSettings struct {
	Name string `json:"name"`
}

// AppSearchPayload https://www.elastic.co/guide/en/cloud/current/definitions.html#AppSearchPayload
type AppSearchPayload struct {
	DisplayName               string            `json:"display_name"`
	ElasticsearchClusterRefID string            `json:"elasticsearch_cluster_ref_id"`
	Plan                      AppSearchPlan     `json:"plan"`
	RefID                     string            `json:"ref_id"`
	Region                    string            `json:"region"`
	Settings                  AppSearchSettings `json:"settings"`
}

// AppSearchPlan https://www.elastic.co/guide/en/cloud/current/definitions.html#AppSearchPlan
type AppSearchPlan struct {
	Appsearch       AppSearchConfiguration              `json:"appsearch"`
	ClusterTopology []AppSearchTopologyElement          `json:"cluster_topology"`
	Transient       TransientAppSearchPlanConfiguration `json:"transient"`
}

// AppSearchConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#AppSearchConfiguration
type AppSearchConfiguration struct {
	DockerImage              string                  `json:"docker_image"`
	SystemSettings           AppSearchSystemSettings `json:"system_settings"`
	UserSettingsJSON         string                  `json:"user_settings_json"`
	UserSettingsOverrideJSON string                  `json:"user_settings_override_json"`
	UserSettingsOverrideYAML string                  `json:"user_settings_override_yaml"`
	UserSettingsYAML         string                  `json:"user_settings_yaml"`
	Version                  string                  `json:"version"`
}

// AppSearchSystemSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#AppSearchSystemSettings
type AppSearchSystemSettings struct {
	ElasticsearchPassword string `json:"elasticsearch_password"`
	ElasticsearchURL      string `json:"elasticsearch_url"`
	ElasticsearchUsername string `json:"elasticsearch_username"`
	SecretSessionKey      string `json:"secret_session_key"`
}

// AppSearchTopologyElement https://www.elastic.co/guide/en/cloud/current/definitions.html#AppSearchTopologyElement
type AppSearchTopologyElement struct {
	Appsearch               AppSearchConfiguration `json:"appsearch"`
	InstanceConfigurationID string                 `json:"instance_configuration_id"`
	NodeType                AppSearchNodeTypes     `json:"node_type"`
	Size                    TopologySize           `json:"size"`
	ZoneCount               int                    `json:"zone_count"`
}

// AppSearchNodeTypes https://www.elastic.co/guide/en/cloud/current/definitions.html#AppSearchNodeTypes
type AppSearchNodeTypes struct {
	Appserver bool `json:"appserver"`
	Worker    bool `json:"worker"`
}

// TransientAppSearchPlanConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#TransientAppSearchPlanConfiguration
type TransientAppSearchPlanConfiguration struct {
	PlanConfiguration AppSearchPlanControlConfiguration `json:"plan_configuration"`
	Strategy          PlanStrategy                      `json:"strategy"`
}

// AppSearchPlanControlConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#AppSearchPlanControlConfiguration
type AppSearchPlanControlConfiguration struct {
	CalmWaitTime        int                    `json:"calm_wait_time"`
	ClusterReboot       string                 `json:"cluster_reboot"`
	ExtendedMaintenance bool                   `json:"extended_maintenance"`
	MoveAllocators      []AllocatorMoveRequest `json:"move_allocators"`
	MoveInstances       []InstanceMoveRequest  `json:"move_instances"`
	PreferredAllocators []string               `json:"preferred_allocators"`
	ReallocateInstances []string               `json:"reallocate_instances"`
	Timeout             int                    `json:"timeout"`
}

// AllocatorMoveRequest https://www.elastic.co/guide/en/cloud/current/definitions.html#AllocatorMoveRequest
type AllocatorMoveRequest struct {
	AllocatorDown bool     `json:"allocator_down"`
	From          string   `json:"from"`
	To            []string `json:"to"`
}

// InstanceMoveRequest https://www.elastic.co/guide/en/cloud/current/definitions.html#InstanceMoveRequest
type InstanceMoveRequest struct {
	InstanceDown bool     `json:"instance_down"`
	From         string   `json:"from"`
	To           []string `json:"to"`
}

// AppSearchSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#AppSearchSettings
type AppSearchSettings struct {
	Metadata ClusterMetadataSettings `json:"metadata"`
}

// ElasticsearchPayload https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchPayload
type ElasticsearchPayload struct {
	Plan   ElasticsearchClusterPlan `json:"plan"`
	RefID  string                   `json:"ref_id"`
	Region string                   `json:"region"`
}

// ElasticsearchClusterPlan https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchClusterPlan
type ElasticsearchClusterPlan struct {
	ClusterTopology    []ElasticsearchClusterTopologyElement `json:"cluster_topology"`
	DeploymentTemplate DeploymentTemplateReference           `json:"deployment_template"`
	Elasticsearch      ElasticsearchConfiguration            `json:"elasticsearch"`
}

// ElasticsearchClusterTopologyElement https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchClusterTopologyElement
type ElasticsearchClusterTopologyElement struct {
	InstanceConfigurationID string                `json:"instance_configuration_id"`
	NodeType                ElasticsearchNodeType `json:"node_type"`
	Size                    TopologySize          `json:"size"`
	ZoneCount               int                   `json:"zone_count"`
}

// ElasticsearchConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchConfiguration
type ElasticsearchConfiguration struct {
	Version string `json:"version"`
}

// ElasticsearchCuration https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchCuration
type ElasticsearchCuration struct {
	FromInstanceConfigurationID string `json:"from_instance_configuration_id"`
	ToInstanceConfigurationID   string `json:"to_instance_configuration_id"`
}

// ElasticsearchSystemSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchSystemSettings
type ElasticsearchSystemSettings struct {
	AutoCreateIndex              bool                               `json:"auto_create_index"`
	DefaultShardsPerIndex        int                                `json:"default_shards_per_index"`
	DestructiveRequiresName      bool                               `json:"destructive_requires_name"`
	EnableCloseIndex             bool                               `json:"enable_close_index"`
	MonitoringCollectionInterval int                                `json:"monitoring_collection_interval"`
	MonitoringHistoryDuration    string                             `json:"monitoring_history_duration"`
	ReindexWhitelist             []string                           `json:"reindex_whitelist"`
	Scripting                    ElasticsearchScriptingUserSettings `json:"scripting"`
	UseDiskThreshold             bool                               `json:"use_disk_threshold"`
	WatcherTriggerEngine         string                             `json:"watcher_trigger_engine"`
}

// ElasticsearchScriptingUserSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchScriptingUserSettings
type ElasticsearchScriptingUserSettings struct {
	ExpressionsEnabled bool                            `json:"expressions_enabled"`
	File               ElasticsearchScriptTypeSettings `json:"file"`
	Inline             ElasticsearchScriptTypeSettings `json:"inline"`
	MustacheEnabled    bool                            `json:"mustache_enabled"`
	PainlessEnabled    bool                            `json:"painless_enabled"`
	Stored             ElasticsearchScriptTypeSettings `json:"stored"`
}

// ElasticsearchScriptTypeSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchScriptTypeSettings
type ElasticsearchScriptTypeSettings struct {
	Enabled     bool `json:"enabled"`
	SandboxMode bool `json:"sandbox_mode"`
}

// ElasticsearchUserBundle https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchUserBundle
type ElasticsearchUserBundle struct {
	ElasticsearchVersion string `json:"elasticsearch_version"`
	Name                 string `json:"name"`
	URL                  string `json:"url"`
}

// ElasticsearchUserPlugin https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchUserPlugin
type ElasticsearchUserPlugin struct {
	ElasticsearchVersion string `json:"elasticsearch_version"`
	Name                 string `json:"name"`
	URL                  string `json:"url"`
}

// ElasticsearchNodeType https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchNodeType
type ElasticsearchNodeType struct {
	Data   bool `json:"data"`
	Ingest bool `json:"ingest"`
	Master bool `json:"master"`
	ML     bool `json:"ml"`
}

// DeploymentTemplateReference https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentTemplateReference
type DeploymentTemplateReference struct {
	ID string `json:"id"`
}

// TransientElasticsearchPlanConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#TransientElasticsearchPlanConfiguration
type TransientElasticsearchPlanConfiguration struct {
	ClusterSettingsJSON string                                `json:"cluster_settings_json"`
	PlanConfiguration   ElasticsearchPlanControlConfiguration `json:"plan_configuration"`
	RestoreSnapshot     RestoreSnapshotConfiguration          `json:"restore_snapshot"`
	Strategy            PlanStrategy                          `json:"strategy"`
}

// ElasticsearchPlanControlConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchPlanControlConfiguration
type ElasticsearchPlanControlConfiguration struct {
	CalmWaitTime        int    `json:"calm_wait_time"`
	ClusterReboot       string `json:"cluster_reboot"`
	ExtendedMaintenance bool   `json:"extended_maintenance"`
	MaxSnapshotAge      int    `json:"max_snapshot_age"`
	MaxSnapshotAttempts int    `json:"max_snapshot_attempts"`
	SkipSnapshot        bool   `json:"skip_snapshot"`
	Timeout             int    `json:"timeout"`
}

// RestoreSnapshotConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#RestoreSnapshotConfiguration
type RestoreSnapshotConfiguration struct {
	RepositoryConfig RestoreSnapshotRepoConfiguration `json:"repository_config"`
	RepositoryName   string                           `json:"repository_name"`
	RestorePayload   RestoreSnapshotAPIConfiguration  `json:"restore_payload"`
	SnapshotName     string                           `json:"snapshot_name"`
	SourceClusterID  string                           `json:"source_cluster_id"`
	Strategy         string                           `json:"strategy"`
}

// RestoreSnapshotRepoConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#RestoreSnapshotRepoConfiguration
type RestoreSnapshotRepoConfiguration struct {
	RawSettings string `json:"raw_settings"`
}

// RestoreSnapshotAPIConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#RestoreSnapshotApiConfiguration
type RestoreSnapshotAPIConfiguration struct {
	Indices     []string `json:"indices"`
	RawSettings string   `json:"raw_settings"`
}

// ElasticsearchClusterSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchClusterSettings
type ElasticsearchClusterSettings struct {
	Ccs                       CrossClusterSearchSettings `json:"ccs"`
	Curation                  ClusterCurationSettings    `json:"curation"`
	DedicatedMastersThreshold int                        `json:"dedicated_masters_threshold"`
	IPFiltering               IPFilteringSettings        `json:"ip_filtering"`
	Metadata                  ClusterMetadataSettings    `json:"metadata"`
	Monitoring                ManagedMonitoringSettings  `json:"monitoring"`
	Snapshot                  ClusterSnapshotSettings    `json:"snapshot"`
	TrafficFilter             TrafficFilterSettings      `json:"traffic_filter"`
}

// CrossClusterSearchSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#CrossClusterSearchSettings
type CrossClusterSearchSettings struct {
	RemoteClusters map[string]string `json:"remote_clusters"`
}

// ClusterCurationSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#ClusterCurationSettings
type ClusterCurationSettings struct {
	Specs []string `json:"specs"`
}

// IPFilteringSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#IpFilteringSettings
type IPFilteringSettings struct {
	Rulesets []string `json:"rulesets"`
}

// ManagedMonitoringSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#ManagedMonitoringSettings
type ManagedMonitoringSettings struct {
	TargetClusterID string `json:"target_cluster_id"`
}

// ClusterSnapshotSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#ClusterSnapshotSettings
type ClusterSnapshotSettings struct {
	CronExpression string                   `json:"cron_expression"`
	Interval       string                   `json:"interval"`
	Retention      ClusterSnapshotRetention `json:"retention"`
	Slm            bool                     `json:"slm"`
}

// ClusterSnapshotRetention https://www.elastic.co/guide/en/cloud/current/definitions.html#ClusterSnapshotRetention
type ClusterSnapshotRetention struct {
	MaxAge    string `json:"max_age"`
	Snapshots int    `json:"snapshots"`
}

// TrafficFilterSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#TrafficFilterSettings
type TrafficFilterSettings struct {
	Rulesets []string `json:"rulesets"`
}

// KibanaPayload https://www.elastic.co/guide/en/cloud/current/definitions.html#KibanaPayload
type KibanaPayload struct {
	ElasticsearchClusterRefID string            `json:"elasticsearch_cluster_ref_id"`
	Plan                      KibanaClusterPlan `json:"plan"`
	RefID                     string            `json:"ref_id"`
	Region                    string            `json:"region"`
}

// KibanaClusterPlan https://www.elastic.co/guide/en/cloud/current/definitions.html#KibanaClusterPlan
type KibanaClusterPlan struct {
	ClusterTopology []KibanaClusterTopologyElement `json:"cluster_topology"`
	Kibana          KibanaConfiguration            `json:"kibana"`
}

// KibanaClusterTopologyElement https://www.elastic.co/guide/en/cloud/current/definitions.html#KibanaClusterTopologyElement
type KibanaClusterTopologyElement struct {
	InstanceConfigurationID string `json:"instance_configuration_id"`
	Size      TopologySize `json:"size"`
	ZoneCount int          `json:"zone_count"`
}

// KibanaConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#KibanaConfiguration
type KibanaConfiguration struct {
	Version string `json:"version"`
}

// KibanaSystemSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#KibanaSystemSettings
type KibanaSystemSettings struct {
	ElasticsearchPassword string `json:"elasticsearch_password"`
	ElasticsearchURL      string `json:"elasticsearch_url"`
	ElasticsearchUsername string `json:"elasticsearch_username"`
}

// TransientKibanaPlanConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#TransientKibanaPlanConfiguration
type TransientKibanaPlanConfiguration struct {
	PlanConfiguration KibanaPlanControlConfiguration `json:"plan_configuration"`
	Strategy          PlanStrategy                   `json:"strategy"`
}

// KibanaPlanControlConfiguration https://www.elastic.co/guide/en/cloud/current/definitions.html#KibanaPlanControlConfiguration
type KibanaPlanControlConfiguration struct {
	CalmWaitTime        int    `json:"calm_wait_time"`
	ClusterReboot       string `json:"cluster_reboot"`
	ExtendedMaintenance bool   `json:"extended_maintenance"`
	Timeout             int    `json:"timeout"`
}

// KibanaClusterSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#KibanaClusterSettings
type KibanaClusterSettings struct {
	Metadata ClusterMetadataSettings `json:"metadata"`
}

// DeploymentCreateSettings https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentCreateSettings
type DeploymentCreateSettings struct {
	IPFilteringSettings   IPFilteringSettings   `json:"ip_filtering_settings"`
	TrafficFilterSettings TrafficFilterSettings `json:"traffic_filter_settings"`
}

// DeploymentCreateResponse https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentCreateResponse
type DeploymentCreateResponse struct {
	Created     bool                  `json:"created"`
	Diagnostics DeploymentDiagnostics `json:"diagnostics"`
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Resources   []DeploymentResource  `json:"resources"`
}

// DeploymentDiagnostics https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentDiagnostics
type DeploymentDiagnostics struct {
	Creates Creates `json:"creates"`
	Updates Updates `json:"updates"`
}

// Creates https://www.elastic.co/guide/en/cloud/current/definitions.html#Creates
type Creates struct {
	Apm              []Apm              `json:"apm"`
	Appsearch        []Appsearch        `json:"appsearch"`
	Elasticsearch    []Elasticsearch    `json:"elasticsearch"`
	EnterpriseSearch []EnterpriseSearch `json:"enterprise_search"`
	Kibana           []Kibana           `json:"kibana"`
}

// Apm https://www.elastic.co/guide/en/cloud/current/definitions.html#Apm
type Apm struct {
	BackendPlan               string `json:"backend_plan"`
	DisplayName               string `json:"display_name"`
	ElasticsearchClusterRefID string `json:"elasticsearch_cluster_ref_id"`
	RefID                     string `json:"ref_id"`
}

// Appsearch https://www.elastic.co/guide/en/cloud/current/definitions.html#AppSearch
type Appsearch struct {
	BackendPlan               string `json:"backend_plan"`
	DisplayName               string `json:"display_name"`
	ElasticsearchClusterRefID string `json:"elasticsearch_cluster_ref_id"`
	RefID                     string `json:"ref_id"`
}

// Elasticsearch https://www.elastic.co/guide/en/cloud/current/definitions.html#Elasticsearch
type Elasticsearch struct {
	BackendPlan string `json:"backend_plan"`
	DisplayName string `json:"display_name"`
	RefID       string `json:"ref_id"`
}

// EnterpriseSearch https://www.elastic.co/guide/en/cloud/current/definitions.html#EnterpriseSearch
type EnterpriseSearch struct {
	BackendPlan               string `json:"backend_plan"`
	DisplayName               string `json:"display_name"`
	ElasticsearchClusterRefID string `json:"elasticsearch_cluster_ref_id"`
	RefID                     string `json:"ref_id"`
}

// Kibana https://www.elastic.co/guide/en/cloud/current/definitions.html#Kibana
type Kibana struct {
	BackendPlan               string `json:"backend_plan"`
	DisplayName               string `json:"display_name"`
	ElasticsearchClusterRefID string `json:"elasticsearch_cluster_ref_id"`
	RefID                     string `json:"ref_id"`
}

// Updates https://www.elastic.co/guide/en/cloud/current/definitions.html#Updates
type Updates struct {
	Apm              []Apm              `json:"apm"`
	Appsearch        []Appsearch        `json:"appsearch"`
	Elasticsearch    []Elasticsearch    `json:"elasticsearch"`
	EnterpriseSearch []EnterpriseSearch `json:"enterprise_search"`
	Kibana           []Kibana           `json:"kibana"`
}

// DeploymentResource https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentResource
type DeploymentResource struct {
	CloudID                   string             `json:"cloud_id"`
	Credentials               ClusterCredentials `json:"credentials"`
	ElasticsearchClusterRefID string             `json:"elasticsearch_cluster_ref_id"`
	ID                        string             `json:"id"`
	Kind                      string             `json:"kind"`
	RefID                     string             `json:"ref_id"`
	Region                    string             `json:"region"`
	SecretToken               string             `json:"secret_token"`
	Warnings                  []ReplyWarning     `json:"warnings"`
}

// ClusterCredentials https://www.elastic.co/guide/en/cloud/current/definitions.html#ClusterCredentials
type ClusterCredentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// ReplyWarning https://www.elastic.co/guide/en/cloud/current/definitions.html#ReplyWarning
type ReplyWarning struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// DeploymentGetResponse https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentGetResponse
type DeploymentGetResponse struct {
	Healthy   bool                `json:"healthy"`
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	Resources DeploymentResources `json:"resources"`
}

// DeploymentResources https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentResources
type DeploymentResources struct {
	Apm              []ApmResourceInfo              `json:"apm"`
	Appsearch        []AppSearchResourceInfo        `json:"appsearch"`
	Elasticsearch    []ElasticsearchResourceInfo    `json:"elasticsearch"`
	EnterpriseSearch []EnterpriseSearchResourceInfo `json:"enterprise_search"`
	Kibana           []KibanaResourceInfo           `json:"kibana"`
}

// ApmResourceInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#ApmResourceInfo
type ApmResourceInfo struct {
	ElasticsearchClusterRefID string  `json:"elasticsearch_cluster_ref_id"`
	ID                        string  `json:"id"`
	Info                      ApmInfo `json:"info"`
	RefID                     string  `json:"ref_id"`
	Region                    string  `json:"region"`
}

// ApmInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#ApmInfo
type ApmInfo struct{}

// AppSearchResourceInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#AppSearchResourceInfo
type AppSearchResourceInfo struct{}

// EnterpriseSearchResourceInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#EnterpriseSearchResourceInfo
type EnterpriseSearchResourceInfo struct{}

// ElasticsearchResourceInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchResourceInfo
type ElasticsearchResourceInfo struct {
	ID     string                   `json:"id"`
	Info   ElasticsearchClusterInfo `json:"info"`
	RefID  string                   `json:"ref_id"`
	Region string                   `json:"region"`
}

// ElasticsearchClusterInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchClusterInfo
type ElasticsearchClusterInfo struct {
	ClusterID    string                        `json:"cluster_id"`
	ClusterName  string                        `json:"cluster_name"`
	DeploymentID string                        `json:"deployment_id"`
	Healthy      bool                          `json:"healthy"`
	PlanInfo     ElasticsearchClusterPlansInfo `json:"plan_info"`
	Status       string                        `json:"status"`
}

// ElasticsearchClusterPlansInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchClusterPlansInfo
type ElasticsearchClusterPlansInfo struct {
	Current ElasticsearchClusterPlanInfo `json:"current"`
}

// ElasticsearchClusterPlanInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#ElasticsearchClusterPlanInfo
type ElasticsearchClusterPlanInfo struct {
	Healthy bool                     `json:"healthy"`
	Plan    ElasticsearchClusterPlan `json:"plan"`
}

// KibanaResourceInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#KibanaResourceInfo
type KibanaResourceInfo struct {
	ElasticsearchClusterRefID string            `json:"elasticsearch_cluster_ref_id"`
	ID                        string            `json:"id"`
	RefID                     string            `json:"ref_id"`
	Region                    string            `json:"region"`
	Info                      KibanaClusterInfo `json:"info"`
}

// KibanaClusterInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#KibanaClusterInfo
type KibanaClusterInfo struct {
	PlanInfo KibanaClusterPlansInfo `json:"plan_info"`
}

// KibanaClusterPlansInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#KibanaClusterPlansInfo
type KibanaClusterPlansInfo struct {
	Current KibanaClusterPlanInfo `json:"current"`
}

// KibanaClusterPlanInfo https://www.elastic.co/guide/en/cloud/current/definitions.html#KibanaClusterPlanInfo
type KibanaClusterPlanInfo struct {
	Plan KibanaClusterPlan `json:"plan"`
}

// DeploymentUpdateRequest https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentUpdateRequest
type DeploymentUpdateRequest struct {
	Name         string                    `json:"name"`
	PruneOrphans bool                      `json:"prune_orphans"`
	Resources    DeploymentUpdateResources `json:"resources"`
}

// DeploymentUpdateResources https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentUpdateResources
type DeploymentUpdateResources struct {
	Elasticsearch []ElasticsearchPayload `json:"elasticsearch"`
}

// DeploymentUpdateResponse https://www.elastic.co/guide/en/cloud/current/definitions.html#DeploymentUpdateResponse
type DeploymentUpdateResponse struct {
	ID        string               `json:"id"`
	Resources []DeploymentResource `json:"resources"`
}

