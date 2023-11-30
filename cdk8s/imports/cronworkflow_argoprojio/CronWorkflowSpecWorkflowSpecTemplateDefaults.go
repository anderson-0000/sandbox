package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaults struct {
	ActiveDeadlineSeconds CronWorkflowSpecWorkflowSpecTemplateDefaultsActiveDeadlineSeconds `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	Affinity *CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	ArchiveLocation *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocation `field:"optional" json:"archiveLocation" yaml:"archiveLocation"`
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	Container *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainer `field:"optional" json:"container" yaml:"container"`
	ContainerSet *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSet `field:"optional" json:"containerSet" yaml:"containerSet"`
	Daemon *bool `field:"optional" json:"daemon" yaml:"daemon"`
	Dag *CronWorkflowSpecWorkflowSpecTemplateDefaultsDag `field:"optional" json:"dag" yaml:"dag"`
	Data *CronWorkflowSpecWorkflowSpecTemplateDefaultsData `field:"optional" json:"data" yaml:"data"`
	Executor *CronWorkflowSpecWorkflowSpecTemplateDefaultsExecutor `field:"optional" json:"executor" yaml:"executor"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	HostAliases *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsHostAliases `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	Http *CronWorkflowSpecWorkflowSpecTemplateDefaultsHttp `field:"optional" json:"http" yaml:"http"`
	InitContainers *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainers `field:"optional" json:"initContainers" yaml:"initContainers"`
	Inputs *CronWorkflowSpecWorkflowSpecTemplateDefaultsInputs `field:"optional" json:"inputs" yaml:"inputs"`
	Memoize *CronWorkflowSpecWorkflowSpecTemplateDefaultsMemoize `field:"optional" json:"memoize" yaml:"memoize"`
	Metadata *CronWorkflowSpecWorkflowSpecTemplateDefaultsMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Metrics *CronWorkflowSpecWorkflowSpecTemplateDefaultsMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	Outputs *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputs `field:"optional" json:"outputs" yaml:"outputs"`
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	Plugin interface{} `field:"optional" json:"plugin" yaml:"plugin"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	Resource *CronWorkflowSpecWorkflowSpecTemplateDefaultsResource `field:"optional" json:"resource" yaml:"resource"`
	RetryStrategy *CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	SchedulerName *string `field:"optional" json:"schedulerName" yaml:"schedulerName"`
	Script *CronWorkflowSpecWorkflowSpecTemplateDefaultsScript `field:"optional" json:"script" yaml:"script"`
	SecurityContext *CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Sidecars *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecars `field:"optional" json:"sidecars" yaml:"sidecars"`
	Steps *[]*[]interface{} `field:"optional" json:"steps" yaml:"steps"`
	Suspend *CronWorkflowSpecWorkflowSpecTemplateDefaultsSuspend `field:"optional" json:"suspend" yaml:"suspend"`
	Synchronization *CronWorkflowSpecWorkflowSpecTemplateDefaultsSynchronization `field:"optional" json:"synchronization" yaml:"synchronization"`
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	Tolerations *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
	Volumes *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumes `field:"optional" json:"volumes" yaml:"volumes"`
}

