package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplates struct {
	ActiveDeadlineSeconds WorkflowTemplateSpecTemplatesActiveDeadlineSeconds `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	Affinity *WorkflowTemplateSpecTemplatesAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	ArchiveLocation *WorkflowTemplateSpecTemplatesArchiveLocation `field:"optional" json:"archiveLocation" yaml:"archiveLocation"`
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	Container *WorkflowTemplateSpecTemplatesContainer `field:"optional" json:"container" yaml:"container"`
	ContainerSet *WorkflowTemplateSpecTemplatesContainerSet `field:"optional" json:"containerSet" yaml:"containerSet"`
	Daemon *bool `field:"optional" json:"daemon" yaml:"daemon"`
	Dag *WorkflowTemplateSpecTemplatesDag `field:"optional" json:"dag" yaml:"dag"`
	Data *WorkflowTemplateSpecTemplatesData `field:"optional" json:"data" yaml:"data"`
	Executor *WorkflowTemplateSpecTemplatesExecutor `field:"optional" json:"executor" yaml:"executor"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	HostAliases *[]*WorkflowTemplateSpecTemplatesHostAliases `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	Http *WorkflowTemplateSpecTemplatesHttp `field:"optional" json:"http" yaml:"http"`
	InitContainers *[]*WorkflowTemplateSpecTemplatesInitContainers `field:"optional" json:"initContainers" yaml:"initContainers"`
	Inputs *WorkflowTemplateSpecTemplatesInputs `field:"optional" json:"inputs" yaml:"inputs"`
	Memoize *WorkflowTemplateSpecTemplatesMemoize `field:"optional" json:"memoize" yaml:"memoize"`
	Metadata *WorkflowTemplateSpecTemplatesMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Metrics *WorkflowTemplateSpecTemplatesMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	Outputs *WorkflowTemplateSpecTemplatesOutputs `field:"optional" json:"outputs" yaml:"outputs"`
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	Plugin interface{} `field:"optional" json:"plugin" yaml:"plugin"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	Resource *WorkflowTemplateSpecTemplatesResource `field:"optional" json:"resource" yaml:"resource"`
	RetryStrategy *WorkflowTemplateSpecTemplatesRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	SchedulerName *string `field:"optional" json:"schedulerName" yaml:"schedulerName"`
	Script *WorkflowTemplateSpecTemplatesScript `field:"optional" json:"script" yaml:"script"`
	SecurityContext *WorkflowTemplateSpecTemplatesSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Sidecars *[]*WorkflowTemplateSpecTemplatesSidecars `field:"optional" json:"sidecars" yaml:"sidecars"`
	Steps *[]*[]interface{} `field:"optional" json:"steps" yaml:"steps"`
	Suspend *WorkflowTemplateSpecTemplatesSuspend `field:"optional" json:"suspend" yaml:"suspend"`
	Synchronization *WorkflowTemplateSpecTemplatesSynchronization `field:"optional" json:"synchronization" yaml:"synchronization"`
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	Tolerations *[]*WorkflowTemplateSpecTemplatesTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
	Volumes *[]*WorkflowTemplateSpecTemplatesVolumes `field:"optional" json:"volumes" yaml:"volumes"`
}

