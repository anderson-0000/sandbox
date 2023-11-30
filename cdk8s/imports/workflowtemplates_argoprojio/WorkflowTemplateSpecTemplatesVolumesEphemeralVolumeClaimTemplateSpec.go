package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpec struct {
	AccessModes *[]*string `field:"optional" json:"accessModes" yaml:"accessModes"`
	DataSource *WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	DataSourceRef *WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef `field:"optional" json:"dataSourceRef" yaml:"dataSourceRef"`
	Resources *WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResources `field:"optional" json:"resources" yaml:"resources"`
	Selector *WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

