package workflows _argoprojio


type WorkflowSpecTemplates struct {
	ActiveDeadlineSeconds WorkflowSpecTemplatesActiveDeadlineSeconds `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	Affinity *WorkflowSpecTemplatesAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	ArchiveLocation *WorkflowSpecTemplatesArchiveLocation `field:"optional" json:"archiveLocation" yaml:"archiveLocation"`
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	Container *WorkflowSpecTemplatesContainer `field:"optional" json:"container" yaml:"container"`
	ContainerSet *WorkflowSpecTemplatesContainerSet `field:"optional" json:"containerSet" yaml:"containerSet"`
	Daemon *bool `field:"optional" json:"daemon" yaml:"daemon"`
	Dag *WorkflowSpecTemplatesDag `field:"optional" json:"dag" yaml:"dag"`
	Data *WorkflowSpecTemplatesData `field:"optional" json:"data" yaml:"data"`
	Executor *WorkflowSpecTemplatesExecutor `field:"optional" json:"executor" yaml:"executor"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	HostAliases *[]*WorkflowSpecTemplatesHostAliases `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	Http *WorkflowSpecTemplatesHttp `field:"optional" json:"http" yaml:"http"`
	InitContainers *[]*WorkflowSpecTemplatesInitContainers `field:"optional" json:"initContainers" yaml:"initContainers"`
	Inputs *WorkflowSpecTemplatesInputs `field:"optional" json:"inputs" yaml:"inputs"`
	Memoize *WorkflowSpecTemplatesMemoize `field:"optional" json:"memoize" yaml:"memoize"`
	Metadata *WorkflowSpecTemplatesMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Metrics *WorkflowSpecTemplatesMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	Outputs *WorkflowSpecTemplatesOutputs `field:"optional" json:"outputs" yaml:"outputs"`
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	Plugin interface{} `field:"optional" json:"plugin" yaml:"plugin"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	Resource *WorkflowSpecTemplatesResource `field:"optional" json:"resource" yaml:"resource"`
	RetryStrategy *WorkflowSpecTemplatesRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	SchedulerName *string `field:"optional" json:"schedulerName" yaml:"schedulerName"`
	Script *WorkflowSpecTemplatesScript `field:"optional" json:"script" yaml:"script"`
	SecurityContext *WorkflowSpecTemplatesSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Sidecars *[]*WorkflowSpecTemplatesSidecars `field:"optional" json:"sidecars" yaml:"sidecars"`
	Steps *[]*[]interface{} `field:"optional" json:"steps" yaml:"steps"`
	Suspend *WorkflowSpecTemplatesSuspend `field:"optional" json:"suspend" yaml:"suspend"`
	Synchronization *WorkflowSpecTemplatesSynchronization `field:"optional" json:"synchronization" yaml:"synchronization"`
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	Tolerations *[]*WorkflowSpecTemplatesTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
	Volumes *[]*WorkflowSpecTemplatesVolumes `field:"optional" json:"volumes" yaml:"volumes"`
}

