package workflows_argoprojio


type WorkflowSpecVolumesEphemeral struct {
	VolumeClaimTemplate *WorkflowSpecVolumesEphemeralVolumeClaimTemplate `field:"optional" json:"volumeClaimTemplate" yaml:"volumeClaimTemplate"`
}

