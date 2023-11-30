package workflows_argoprojio


type WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpec struct {
	AccessModes *[]*string `field:"optional" json:"accessModes" yaml:"accessModes"`
	DataSource *WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	DataSourceRef *WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef `field:"optional" json:"dataSourceRef" yaml:"dataSourceRef"`
	Resources *WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResources `field:"optional" json:"resources" yaml:"resources"`
	Selector *WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

