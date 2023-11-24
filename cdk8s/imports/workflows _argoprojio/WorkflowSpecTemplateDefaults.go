package workflows _argoprojio


type WorkflowSpecTemplateDefaults struct {
	ActiveDeadlineSeconds WorkflowSpecTemplateDefaultsActiveDeadlineSeconds `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	Affinity *WorkflowSpecTemplateDefaultsAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	ArchiveLocation *WorkflowSpecTemplateDefaultsArchiveLocation `field:"optional" json:"archiveLocation" yaml:"archiveLocation"`
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	Container *WorkflowSpecTemplateDefaultsContainer `field:"optional" json:"container" yaml:"container"`
	ContainerSet *WorkflowSpecTemplateDefaultsContainerSet `field:"optional" json:"containerSet" yaml:"containerSet"`
	Daemon *bool `field:"optional" json:"daemon" yaml:"daemon"`
	Dag *WorkflowSpecTemplateDefaultsDag `field:"optional" json:"dag" yaml:"dag"`
	Data *WorkflowSpecTemplateDefaultsData `field:"optional" json:"data" yaml:"data"`
	Executor *WorkflowSpecTemplateDefaultsExecutor `field:"optional" json:"executor" yaml:"executor"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	HostAliases *[]*WorkflowSpecTemplateDefaultsHostAliases `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	Http *WorkflowSpecTemplateDefaultsHttp `field:"optional" json:"http" yaml:"http"`
	InitContainers *[]*WorkflowSpecTemplateDefaultsInitContainers `field:"optional" json:"initContainers" yaml:"initContainers"`
	Inputs *WorkflowSpecTemplateDefaultsInputs `field:"optional" json:"inputs" yaml:"inputs"`
	Memoize *WorkflowSpecTemplateDefaultsMemoize `field:"optional" json:"memoize" yaml:"memoize"`
	Metadata *WorkflowSpecTemplateDefaultsMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Metrics *WorkflowSpecTemplateDefaultsMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	Outputs *WorkflowSpecTemplateDefaultsOutputs `field:"optional" json:"outputs" yaml:"outputs"`
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	Plugin interface{} `field:"optional" json:"plugin" yaml:"plugin"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	Resource *WorkflowSpecTemplateDefaultsResource `field:"optional" json:"resource" yaml:"resource"`
	RetryStrategy *WorkflowSpecTemplateDefaultsRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	SchedulerName *string `field:"optional" json:"schedulerName" yaml:"schedulerName"`
	Script *WorkflowSpecTemplateDefaultsScript `field:"optional" json:"script" yaml:"script"`
	SecurityContext *WorkflowSpecTemplateDefaultsSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Sidecars *[]*WorkflowSpecTemplateDefaultsSidecars `field:"optional" json:"sidecars" yaml:"sidecars"`
	Steps *[]*[]interface{} `field:"optional" json:"steps" yaml:"steps"`
	Suspend *WorkflowSpecTemplateDefaultsSuspend `field:"optional" json:"suspend" yaml:"suspend"`
	Synchronization *WorkflowSpecTemplateDefaultsSynchronization `field:"optional" json:"synchronization" yaml:"synchronization"`
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	Tolerations *[]*WorkflowSpecTemplateDefaultsTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
	Volumes *[]*WorkflowSpecTemplateDefaultsVolumes `field:"optional" json:"volumes" yaml:"volumes"`
}

