package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplates struct {
	ActiveDeadlineSeconds CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	Affinity *CronWorkflowSpecWorkflowSpecTemplatesAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	ArchiveLocation *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocation `field:"optional" json:"archiveLocation" yaml:"archiveLocation"`
	AutomountServiceAccountToken *bool `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	Container *CronWorkflowSpecWorkflowSpecTemplatesContainer `field:"optional" json:"container" yaml:"container"`
	ContainerSet *CronWorkflowSpecWorkflowSpecTemplatesContainerSet `field:"optional" json:"containerSet" yaml:"containerSet"`
	Daemon *bool `field:"optional" json:"daemon" yaml:"daemon"`
	Dag *CronWorkflowSpecWorkflowSpecTemplatesDag `field:"optional" json:"dag" yaml:"dag"`
	Data *CronWorkflowSpecWorkflowSpecTemplatesData `field:"optional" json:"data" yaml:"data"`
	Executor *CronWorkflowSpecWorkflowSpecTemplatesExecutor `field:"optional" json:"executor" yaml:"executor"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	HostAliases *[]*CronWorkflowSpecWorkflowSpecTemplatesHostAliases `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	Http *CronWorkflowSpecWorkflowSpecTemplatesHttp `field:"optional" json:"http" yaml:"http"`
	InitContainers *[]*CronWorkflowSpecWorkflowSpecTemplatesInitContainers `field:"optional" json:"initContainers" yaml:"initContainers"`
	Inputs *CronWorkflowSpecWorkflowSpecTemplatesInputs `field:"optional" json:"inputs" yaml:"inputs"`
	Memoize *CronWorkflowSpecWorkflowSpecTemplatesMemoize `field:"optional" json:"memoize" yaml:"memoize"`
	Metadata *CronWorkflowSpecWorkflowSpecTemplatesMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Metrics *CronWorkflowSpecWorkflowSpecTemplatesMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	Outputs *CronWorkflowSpecWorkflowSpecTemplatesOutputs `field:"optional" json:"outputs" yaml:"outputs"`
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	Plugin interface{} `field:"optional" json:"plugin" yaml:"plugin"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	Resource *CronWorkflowSpecWorkflowSpecTemplatesResource `field:"optional" json:"resource" yaml:"resource"`
	RetryStrategy *CronWorkflowSpecWorkflowSpecTemplatesRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	SchedulerName *string `field:"optional" json:"schedulerName" yaml:"schedulerName"`
	Script *CronWorkflowSpecWorkflowSpecTemplatesScript `field:"optional" json:"script" yaml:"script"`
	SecurityContext *CronWorkflowSpecWorkflowSpecTemplatesSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Sidecars *[]*CronWorkflowSpecWorkflowSpecTemplatesSidecars `field:"optional" json:"sidecars" yaml:"sidecars"`
	Steps *[]*[]interface{} `field:"optional" json:"steps" yaml:"steps"`
	Suspend *CronWorkflowSpecWorkflowSpecTemplatesSuspend `field:"optional" json:"suspend" yaml:"suspend"`
	Synchronization *CronWorkflowSpecWorkflowSpecTemplatesSynchronization `field:"optional" json:"synchronization" yaml:"synchronization"`
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	Tolerations *[]*CronWorkflowSpecWorkflowSpecTemplatesTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
	Volumes *[]*CronWorkflowSpecWorkflowSpecTemplatesVolumes `field:"optional" json:"volumes" yaml:"volumes"`
}

