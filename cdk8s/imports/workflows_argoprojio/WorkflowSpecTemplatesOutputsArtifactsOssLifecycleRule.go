package workflows_argoprojio


type WorkflowSpecTemplatesOutputsArtifactsOssLifecycleRule struct {
	MarkDeletionAfterDays *float64 `field:"optional" json:"markDeletionAfterDays" yaml:"markDeletionAfterDays"`
	MarkInfrequentAccessAfterDays *float64 `field:"optional" json:"markInfrequentAccessAfterDays" yaml:"markInfrequentAccessAfterDays"`
}

