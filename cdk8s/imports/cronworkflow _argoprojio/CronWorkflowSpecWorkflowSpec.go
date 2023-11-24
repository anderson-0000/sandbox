package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpec struct {
	ActiveDeadlineSeconds *float64 `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	Affinity *CronWorkflowSpecWorkflowSpecAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	Arguments *CronWorkflowSpecWorkflowSpecArguments `field:"optional" json:"arguments" yaml:"arguments"`
	ArtifactGc *CronWorkflowSpecWorkflowSpecArtifactGc `field:"optional" json:"artifactGc" yaml:"artifactGc"`
	ArtifactRepositoryRef *CronWorkflowSpecWorkflowSpecArtifactRepositoryRef `field:"optional" json:"artifactRepositoryRef" yaml:"artifactRepositoryRef"`
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	DnsConfig *CronWorkflowSpecWorkflowSpecDnsConfig `field:"optional" json:"dnsConfig" yaml:"dnsConfig"`
	DnsPolicy *string `field:"optional" json:"dnsPolicy" yaml:"dnsPolicy"`
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	Executor *CronWorkflowSpecWorkflowSpecExecutor `field:"optional" json:"executor" yaml:"executor"`
	Hooks *map[string]*CronWorkflowSpecWorkflowSpecHooks `field:"optional" json:"hooks" yaml:"hooks"`
	HostAliases *[]*CronWorkflowSpecWorkflowSpecHostAliases `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	HostNetwork *bool `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	ImagePullSecrets *[]*CronWorkflowSpecWorkflowSpecImagePullSecrets `field:"optional" json:"imagePullSecrets" yaml:"imagePullSecrets"`
	Metrics *CronWorkflowSpecWorkflowSpecMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	OnExit *string `field:"optional" json:"onExit" yaml:"onExit"`
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	PodDisruptionBudget *CronWorkflowSpecWorkflowSpecPodDisruptionBudget `field:"optional" json:"podDisruptionBudget" yaml:"podDisruptionBudget"`
	PodGc *CronWorkflowSpecWorkflowSpecPodGc `field:"optional" json:"podGc" yaml:"podGc"`
	PodMetadata *CronWorkflowSpecWorkflowSpecPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	PodPriority *float64 `field:"optional" json:"podPriority" yaml:"podPriority"`
	PodPriorityClassName *string `field:"optional" json:"podPriorityClassName" yaml:"podPriorityClassName"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	RetryStrategy *CronWorkflowSpecWorkflowSpecRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	SchedulerName *string `field:"optional" json:"schedulerName" yaml:"schedulerName"`
	SecurityContext *CronWorkflowSpecWorkflowSpecSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Shutdown *string `field:"optional" json:"shutdown" yaml:"shutdown"`
	Suspend *bool `field:"optional" json:"suspend" yaml:"suspend"`
	Synchronization *CronWorkflowSpecWorkflowSpecSynchronization `field:"optional" json:"synchronization" yaml:"synchronization"`
	TemplateDefaults *CronWorkflowSpecWorkflowSpecTemplateDefaults `field:"optional" json:"templateDefaults" yaml:"templateDefaults"`
	Templates *[]*CronWorkflowSpecWorkflowSpecTemplates `field:"optional" json:"templates" yaml:"templates"`
	Tolerations *[]*CronWorkflowSpecWorkflowSpecTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
	TtlStrategy *CronWorkflowSpecWorkflowSpecTtlStrategy `field:"optional" json:"ttlStrategy" yaml:"ttlStrategy"`
	VolumeClaimGc *CronWorkflowSpecWorkflowSpecVolumeClaimGc `field:"optional" json:"volumeClaimGc" yaml:"volumeClaimGc"`
	VolumeClaimTemplates *[]*CronWorkflowSpecWorkflowSpecVolumeClaimTemplates `field:"optional" json:"volumeClaimTemplates" yaml:"volumeClaimTemplates"`
	Volumes *[]*CronWorkflowSpecWorkflowSpecVolumes `field:"optional" json:"volumes" yaml:"volumes"`
	WorkflowMetadata *CronWorkflowSpecWorkflowSpecWorkflowMetadata `field:"optional" json:"workflowMetadata" yaml:"workflowMetadata"`
	WorkflowTemplateRef *CronWorkflowSpecWorkflowSpecWorkflowTemplateRef `field:"optional" json:"workflowTemplateRef" yaml:"workflowTemplateRef"`
}

