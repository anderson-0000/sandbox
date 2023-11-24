package workflows _argoprojio


type WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource struct {
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	Name *string `field:"required" json:"name" yaml:"name"`
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
}

