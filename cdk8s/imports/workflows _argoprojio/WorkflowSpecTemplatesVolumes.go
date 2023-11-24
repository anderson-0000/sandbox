package workflows _argoprojio


type WorkflowSpecTemplatesVolumes struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	AwsElasticBlockStore *WorkflowSpecTemplatesVolumesAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	AzureDisk *WorkflowSpecTemplatesVolumesAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	AzureFile *WorkflowSpecTemplatesVolumesAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	Cephfs *WorkflowSpecTemplatesVolumesCephfs `field:"optional" json:"cephfs" yaml:"cephfs"`
	Cinder *WorkflowSpecTemplatesVolumesCinder `field:"optional" json:"cinder" yaml:"cinder"`
	ConfigMap *WorkflowSpecTemplatesVolumesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	Csi *WorkflowSpecTemplatesVolumesCsi `field:"optional" json:"csi" yaml:"csi"`
	DownwardApi *WorkflowSpecTemplatesVolumesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	EmptyDir *WorkflowSpecTemplatesVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	Ephemeral *WorkflowSpecTemplatesVolumesEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	Fc *WorkflowSpecTemplatesVolumesFc `field:"optional" json:"fc" yaml:"fc"`
	FlexVolume *WorkflowSpecTemplatesVolumesFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	Flocker *WorkflowSpecTemplatesVolumesFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	GcePersistentDisk *WorkflowSpecTemplatesVolumesGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	GitRepo *WorkflowSpecTemplatesVolumesGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	Glusterfs *WorkflowSpecTemplatesVolumesGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	HostPath *WorkflowSpecTemplatesVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	Iscsi *WorkflowSpecTemplatesVolumesIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	Nfs *WorkflowSpecTemplatesVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	PersistentVolumeClaim *WorkflowSpecTemplatesVolumesPersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	PhotonPersistentDisk *WorkflowSpecTemplatesVolumesPhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	PortworxVolume *WorkflowSpecTemplatesVolumesPortworxVolume `field:"optional" json:"portworxVolume" yaml:"portworxVolume"`
	Projected *WorkflowSpecTemplatesVolumesProjected `field:"optional" json:"projected" yaml:"projected"`
	Quobyte *WorkflowSpecTemplatesVolumesQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	Rbd *WorkflowSpecTemplatesVolumesRbd `field:"optional" json:"rbd" yaml:"rbd"`
	ScaleIo *WorkflowSpecTemplatesVolumesScaleIo `field:"optional" json:"scaleIo" yaml:"scaleIo"`
	Secret *WorkflowSpecTemplatesVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
	Storageos *WorkflowSpecTemplatesVolumesStorageos `field:"optional" json:"storageos" yaml:"storageos"`
	VsphereVolume *WorkflowSpecTemplatesVolumesVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

