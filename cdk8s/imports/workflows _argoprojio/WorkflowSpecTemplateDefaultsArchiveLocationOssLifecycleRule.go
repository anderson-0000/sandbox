package workflows _argoprojio


type WorkflowSpecTemplateDefaultsArchiveLocationOssLifecycleRule struct {
	MarkDeletionAfterDays *float64 `field:"optional" json:"markDeletionAfterDays" yaml:"markDeletionAfterDays"`
	MarkInfrequentAccessAfterDays *float64 `field:"optional" json:"markInfrequentAccessAfterDays" yaml:"markInfrequentAccessAfterDays"`
}

