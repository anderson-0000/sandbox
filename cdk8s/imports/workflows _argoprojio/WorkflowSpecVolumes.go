package workflows _argoprojio


type WorkflowSpecVolumes struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	AwsElasticBlockStore *WorkflowSpecVolumesAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	AzureDisk *WorkflowSpecVolumesAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	AzureFile *WorkflowSpecVolumesAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	Cephfs *WorkflowSpecVolumesCephfs `field:"optional" json:"cephfs" yaml:"cephfs"`
	Cinder *WorkflowSpecVolumesCinder `field:"optional" json:"cinder" yaml:"cinder"`
	ConfigMap *WorkflowSpecVolumesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	Csi *WorkflowSpecVolumesCsi `field:"optional" json:"csi" yaml:"csi"`
	DownwardApi *WorkflowSpecVolumesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	EmptyDir *WorkflowSpecVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	Ephemeral *WorkflowSpecVolumesEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	Fc *WorkflowSpecVolumesFc `field:"optional" json:"fc" yaml:"fc"`
	FlexVolume *WorkflowSpecVolumesFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	Flocker *WorkflowSpecVolumesFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	GcePersistentDisk *WorkflowSpecVolumesGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	GitRepo *WorkflowSpecVolumesGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	Glusterfs *WorkflowSpecVolumesGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	HostPath *WorkflowSpecVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	Iscsi *WorkflowSpecVolumesIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	Nfs *WorkflowSpecVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	PersistentVolumeClaim *WorkflowSpecVolumesPersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	PhotonPersistentDisk *WorkflowSpecVolumesPhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	PortworxVolume *WorkflowSpecVolumesPortworxVolume `field:"optional" json:"portworxVolume" yaml:"portworxVolume"`
	Projected *WorkflowSpecVolumesProjected `field:"optional" json:"projected" yaml:"projected"`
	Quobyte *WorkflowSpecVolumesQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	Rbd *WorkflowSpecVolumesRbd `field:"optional" json:"rbd" yaml:"rbd"`
	ScaleIo *WorkflowSpecVolumesScaleIo `field:"optional" json:"scaleIo" yaml:"scaleIo"`
	Secret *WorkflowSpecVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
	Storageos *WorkflowSpecVolumesStorageos `field:"optional" json:"storageos" yaml:"storageos"`
	VsphereVolume *WorkflowSpecVolumesVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

