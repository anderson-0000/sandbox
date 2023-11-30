package workflows_argoprojio


type WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpec struct {
	AccessModes *[]*string `field:"optional" json:"accessModes" yaml:"accessModes"`
	DataSource *WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	DataSourceRef *WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef `field:"optional" json:"dataSourceRef" yaml:"dataSourceRef"`
	Resources *WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResources `field:"optional" json:"resources" yaml:"resources"`
	Selector *WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

