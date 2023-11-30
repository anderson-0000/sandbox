package workflowtemplates_argoprojio


type WorkflowTemplateSpec struct {
	ActiveDeadlineSeconds *float64 `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	Affinity *WorkflowTemplateSpecAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	Arguments *WorkflowTemplateSpecArguments `field:"optional" json:"arguments" yaml:"arguments"`
	ArtifactGc *WorkflowTemplateSpecArtifactGc `field:"optional" json:"artifactGc" yaml:"artifactGc"`
	ArtifactRepositoryRef *WorkflowTemplateSpecArtifactRepositoryRef `field:"optional" json:"artifactRepositoryRef" yaml:"artifactRepositoryRef"`
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	DnsConfig *WorkflowTemplateSpecDnsConfig `field:"optional" json:"dnsConfig" yaml:"dnsConfig"`
	DnsPolicy *string `field:"optional" json:"dnsPolicy" yaml:"dnsPolicy"`
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	Executor *WorkflowTemplateSpecExecutor `field:"optional" json:"executor" yaml:"executor"`
	Hooks *map[string]*WorkflowTemplateSpecHooks `field:"optional" json:"hooks" yaml:"hooks"`
	HostAliases *[]*WorkflowTemplateSpecHostAliases `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	HostNetwork *bool `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	ImagePullSecrets *[]*WorkflowTemplateSpecImagePullSecrets `field:"optional" json:"imagePullSecrets" yaml:"imagePullSecrets"`
	Metrics *WorkflowTemplateSpecMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	OnExit *string `field:"optional" json:"onExit" yaml:"onExit"`
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	PodDisruptionBudget *WorkflowTemplateSpecPodDisruptionBudget `field:"optional" json:"podDisruptionBudget" yaml:"podDisruptionBudget"`
	PodGc *WorkflowTemplateSpecPodGc `field:"optional" json:"podGc" yaml:"podGc"`
	PodMetadata *WorkflowTemplateSpecPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	PodPriority *float64 `field:"optional" json:"podPriority" yaml:"podPriority"`
	PodPriorityClassName *string `field:"optional" json:"podPriorityClassName" yaml:"podPriorityClassName"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	RetryStrategy *WorkflowTemplateSpecRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	SchedulerName *string `field:"optional" json:"schedulerName" yaml:"schedulerName"`
	SecurityContext *WorkflowTemplateSpecSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Shutdown *string `field:"optional" json:"shutdown" yaml:"shutdown"`
	Suspend *bool `field:"optional" json:"suspend" yaml:"suspend"`
	Synchronization *WorkflowTemplateSpecSynchronization `field:"optional" json:"synchronization" yaml:"synchronization"`
	TemplateDefaults *WorkflowTemplateSpecTemplateDefaults `field:"optional" json:"templateDefaults" yaml:"templateDefaults"`
	Templates *[]*WorkflowTemplateSpecTemplates `field:"optional" json:"templates" yaml:"templates"`
	Tolerations *[]*WorkflowTemplateSpecTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
	TtlStrategy *WorkflowTemplateSpecTtlStrategy `field:"optional" json:"ttlStrategy" yaml:"ttlStrategy"`
	VolumeClaimGc *WorkflowTemplateSpecVolumeClaimGc `field:"optional" json:"volumeClaimGc" yaml:"volumeClaimGc"`
	VolumeClaimTemplates *[]*WorkflowTemplateSpecVolumeClaimTemplates `field:"optional" json:"volumeClaimTemplates" yaml:"volumeClaimTemplates"`
	Volumes *[]*WorkflowTemplateSpecVolumes `field:"optional" json:"volumes" yaml:"volumes"`
	WorkflowMetadata *WorkflowTemplateSpecWorkflowMetadata `field:"optional" json:"workflowMetadata" yaml:"workflowMetadata"`
	WorkflowTemplateRef *WorkflowTemplateSpecWorkflowTemplateRef `field:"optional" json:"workflowTemplateRef" yaml:"workflowTemplateRef"`
}

