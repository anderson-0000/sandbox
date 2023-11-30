package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaults struct {
	ActiveDeadlineSeconds WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	Affinity *WorkflowTemplateSpecTemplateDefaultsAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	ArchiveLocation *WorkflowTemplateSpecTemplateDefaultsArchiveLocation `field:"optional" json:"archiveLocation" yaml:"archiveLocation"`
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	Container *WorkflowTemplateSpecTemplateDefaultsContainer `field:"optional" json:"container" yaml:"container"`
	ContainerSet *WorkflowTemplateSpecTemplateDefaultsContainerSet `field:"optional" json:"containerSet" yaml:"containerSet"`
	Daemon *bool `field:"optional" json:"daemon" yaml:"daemon"`
	Dag *WorkflowTemplateSpecTemplateDefaultsDag `field:"optional" json:"dag" yaml:"dag"`
	Data *WorkflowTemplateSpecTemplateDefaultsData `field:"optional" json:"data" yaml:"data"`
	Executor *WorkflowTemplateSpecTemplateDefaultsExecutor `field:"optional" json:"executor" yaml:"executor"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	HostAliases *[]*WorkflowTemplateSpecTemplateDefaultsHostAliases `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	Http *WorkflowTemplateSpecTemplateDefaultsHttp `field:"optional" json:"http" yaml:"http"`
	InitContainers *[]*WorkflowTemplateSpecTemplateDefaultsInitContainers `field:"optional" json:"initContainers" yaml:"initContainers"`
	Inputs *WorkflowTemplateSpecTemplateDefaultsInputs `field:"optional" json:"inputs" yaml:"inputs"`
	Memoize *WorkflowTemplateSpecTemplateDefaultsMemoize `field:"optional" json:"memoize" yaml:"memoize"`
	Metadata *WorkflowTemplateSpecTemplateDefaultsMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Metrics *WorkflowTemplateSpecTemplateDefaultsMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	Outputs *WorkflowTemplateSpecTemplateDefaultsOutputs `field:"optional" json:"outputs" yaml:"outputs"`
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	Plugin interface{} `field:"optional" json:"plugin" yaml:"plugin"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	Resource *WorkflowTemplateSpecTemplateDefaultsResource `field:"optional" json:"resource" yaml:"resource"`
	RetryStrategy *WorkflowTemplateSpecTemplateDefaultsRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	SchedulerName *string `field:"optional" json:"schedulerName" yaml:"schedulerName"`
	Script *WorkflowTemplateSpecTemplateDefaultsScript `field:"optional" json:"script" yaml:"script"`
	SecurityContext *WorkflowTemplateSpecTemplateDefaultsSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Sidecars *[]*WorkflowTemplateSpecTemplateDefaultsSidecars `field:"optional" json:"sidecars" yaml:"sidecars"`
	Steps *[]*[]interface{} `field:"optional" json:"steps" yaml:"steps"`
	Suspend *WorkflowTemplateSpecTemplateDefaultsSuspend `field:"optional" json:"suspend" yaml:"suspend"`
	Synchronization *WorkflowTemplateSpecTemplateDefaultsSynchronization `field:"optional" json:"synchronization" yaml:"synchronization"`
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	Tolerations *[]*WorkflowTemplateSpecTemplateDefaultsTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
	Volumes *[]*WorkflowTemplateSpecTemplateDefaultsVolumes `field:"optional" json:"volumes" yaml:"volumes"`
}

