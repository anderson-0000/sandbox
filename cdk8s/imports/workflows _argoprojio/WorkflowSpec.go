package workflows _argoprojio


type WorkflowSpec struct {
	ActiveDeadlineSeconds *float64 `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	Affinity *WorkflowSpecAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	Arguments *WorkflowSpecArguments `field:"optional" json:"arguments" yaml:"arguments"`
	ArtifactGc *WorkflowSpecArtifactGc `field:"optional" json:"artifactGc" yaml:"artifactGc"`
	ArtifactRepositoryRef *WorkflowSpecArtifactRepositoryRef `field:"optional" json:"artifactRepositoryRef" yaml:"artifactRepositoryRef"`
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	DnsConfig *WorkflowSpecDnsConfig `field:"optional" json:"dnsConfig" yaml:"dnsConfig"`
	DnsPolicy *string `field:"optional" json:"dnsPolicy" yaml:"dnsPolicy"`
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	Executor *WorkflowSpecExecutor `field:"optional" json:"executor" yaml:"executor"`
	Hooks *map[string]*WorkflowSpecHooks `field:"optional" json:"hooks" yaml:"hooks"`
	HostAliases *[]*WorkflowSpecHostAliases `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	HostNetwork *bool `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	ImagePullSecrets *[]*WorkflowSpecImagePullSecrets `field:"optional" json:"imagePullSecrets" yaml:"imagePullSecrets"`
	Metrics *WorkflowSpecMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	OnExit *string `field:"optional" json:"onExit" yaml:"onExit"`
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	PodDisruptionBudget *WorkflowSpecPodDisruptionBudget `field:"optional" json:"podDisruptionBudget" yaml:"podDisruptionBudget"`
	PodGc *WorkflowSpecPodGc `field:"optional" json:"podGc" yaml:"podGc"`
	PodMetadata *WorkflowSpecPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	PodPriority *float64 `field:"optional" json:"podPriority" yaml:"podPriority"`
	PodPriorityClassName *string `field:"optional" json:"podPriorityClassName" yaml:"podPriorityClassName"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	RetryStrategy *WorkflowSpecRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	SchedulerName *string `field:"optional" json:"schedulerName" yaml:"schedulerName"`
	SecurityContext *WorkflowSpecSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Shutdown *string `field:"optional" json:"shutdown" yaml:"shutdown"`
	Suspend *bool `field:"optional" json:"suspend" yaml:"suspend"`
	Synchronization *WorkflowSpecSynchronization `field:"optional" json:"synchronization" yaml:"synchronization"`
	TemplateDefaults *WorkflowSpecTemplateDefaults `field:"optional" json:"templateDefaults" yaml:"templateDefaults"`
	Templates *[]*WorkflowSpecTemplates `field:"optional" json:"templates" yaml:"templates"`
	Tolerations *[]*WorkflowSpecTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
	TtlStrategy *WorkflowSpecTtlStrategy `field:"optional" json:"ttlStrategy" yaml:"ttlStrategy"`
	VolumeClaimGc *WorkflowSpecVolumeClaimGc `field:"optional" json:"volumeClaimGc" yaml:"volumeClaimGc"`
	VolumeClaimTemplates *[]*WorkflowSpecVolumeClaimTemplates `field:"optional" json:"volumeClaimTemplates" yaml:"volumeClaimTemplates"`
	Volumes *[]*WorkflowSpecVolumes `field:"optional" json:"volumes" yaml:"volumes"`
	WorkflowMetadata *WorkflowSpecWorkflowMetadata `field:"optional" json:"workflowMetadata" yaml:"workflowMetadata"`
	WorkflowTemplateRef *WorkflowSpecWorkflowTemplateRef `field:"optional" json:"workflowTemplateRef" yaml:"workflowTemplateRef"`
}

