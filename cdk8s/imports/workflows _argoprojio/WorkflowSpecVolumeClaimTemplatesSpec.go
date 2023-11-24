package workflows _argoprojio


type WorkflowSpecVolumeClaimTemplatesSpec struct {
	AccessModes *[]*string `field:"optional" json:"accessModes" yaml:"accessModes"`
	DataSource *WorkflowSpecVolumeClaimTemplatesSpecDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	DataSourceRef *WorkflowSpecVolumeClaimTemplatesSpecDataSourceRef `field:"optional" json:"dataSourceRef" yaml:"dataSourceRef"`
	Resources *WorkflowSpecVolumeClaimTemplatesSpecResources `field:"optional" json:"resources" yaml:"resources"`
	Selector *WorkflowSpecVolumeClaimTemplatesSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

