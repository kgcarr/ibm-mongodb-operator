package resources

import(
  corev1 "k8s.io/api/core/v1"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  appsv1 "k8s.io/api/apps/v1"
  "k8s.io/apimachinery/pkg/util/intstr"
)

// Structures for Statefulset Data
type stsData struct {
  Replicas       int
  ImageRepo      string
  StorageClass   string
  InitImage      string
  BootstrapImage string
  MetricsImage   string
}

// Build functions are used to build generic or common objects

// Constants for Volume Mounts
const mongodbdirWDVolumeMount = corev1.VolumeMount{
  Name: "mongodbdir",
  MountPath: "/work-dir",
  SubPath: "workdir",
}

const configdirVolumeMount = corev1.VolumeMount{
  Name: "configdir",
  MountPath: "/data/configdb",
}

const configVolumeMount = corev1.VolumeMount{
  Name: "config",
  MountPath: "/configdb-readonly",
}

const installVolumeMount = corev1.VolumeMount{
  Name: "install",
  MountPath: "/install",
}

const initVolumeMount = corev1.VolumeMount{
  Name: "init",
  MountPath: "/init",
}

const keydirVolumeMount = corev1.VolumeMount{
  Name: "keydir",
  MountPath: "/keydir-readonly",
}

const caVolumeMount = corev1.VolumeMount{
  Name: "ca",
  MountPath: "/ca-readonly",
}

const mongodbdirDDVolumeMount = corev1.VolumeMount{
  Name: "mongodbdir",
  SubPath: "datadir",
  MountPath: "/data/db",
}

const tmpVolumeMount = corev1.VolumeMount{
  Name: "tmp-mongodb",
  MountPath: "/tmp",
}

const tmpMetricsVolumeMount = corev1.VolumeMount{
  Name: "tmp-metrics",
  MountPath: "/tmp",
}

// Constants for Environment Variables
const podNamespaceEV = corev1.EnvVar{
  Name: "POD_NAMESPACE",
  ValueFrom: &corev1.EnvVarSource{
    FieldRef: &corev1.ObjectFieldSelector{
      APIVersion: "v1",
      FieldPath: "metadata.namespace",
    },
  }
}

const replicaSetEV = corev1.EnvVar{
  Name: "REPLICA_SET",
  Value: "rs0",
}

const authEV = corev1.EnvVar{
  Name: "AUTH",
  Value: "true",
}

const adminUserEV = corev1.EnvVar{
  Name: "ADMIN_USER"
  ValueFrom: &corev1.EnvVarSource{
    SecretKeyRef: &corev1.SecretKeySelector{
      LocalObjectReference: corev1.LocalObjectReference{
        Name: "icp-mongodb-admin",
      },
      Key: "user",
    },
  },
}

const adminPasswordEV = corev1.EnvVar{
  Name: "ADMIN_PASSWORD"
  ValueFrom: &corev1.EnvVarSource{
    SecretKeyRef: &corev1.SecretKeySelector{
      LocalObjectReference: corev1.LocalObjectReference{
        Name: "icp-mongodb-admin",
      },
      Key: "password",
    },
  },
}

const metricsEV = corev1.EnvVar{
  Name: "METRICS",
  Value: "true"
}

const metricsUserEV = corev1.EnvVar{
  Name: "METRICS_USER"
  ValueFrom: &corev1.EnvVarSource{
    SecretKeyRef: &corev1.SecretKeySelector{
      LocalObjectReference: corev1.LocalObjectReference{
        Name: "icp-mongodb-metrics",
      },
      Key: "user",
    },
  },
}

const metricsPasswordEV = corev1.EnvVar{
  Name: "METRICS_PASSWORD"
  ValueFrom: &corev1.EnvVarSource{
    SecretKeyRef: &corev1.SecretKeySelector{
      LocalObjectReference: corev1.LocalObjectReference{
        Name: "icp-mongodb-metrics",
      },
      Key: "password",
    },
  },
}

const networkIPVersionEV = corev1.EnvVar{
  Name: "NETWORK_IP_VERSION",
  Value: "ipv4",
}

// Constant for Resource Limits
const mongoDBResourceLimit = BuildResourceRequirement(5*1024)

// Constant for mongoDB Init SecurityContext
const trueValue = true
const falseValue = false

const mongoDBInitSecurityContext = corev1.SecurityContext{
	AllowPrivilegeEscalation: &falseValue,
	ReadOnlyRootFilesystem:   &trueValue,
}

// Create the labels map for kubernetes resources
func BuildLabels() map[string]string {
  var labels = make(map[string]string)
  labels["app.kubernetes.io/name"] = "icp-mongodb"
  labels["app.kubernetes.io/instance"] = "icp-mongodb"
  labels["app.kubernetes.io/version"] = "4.0.12"
  labels["app.kubernetes.io/component"] = "database"
  labels["app.kubernetes.io/part-of"] = "common-services-cloud-pak"
  labels["app.kubernetes.io/managed-by"] = "operator"
  labels["release"] = "mongodb"

  return labels
}

//Create Object Metadata
func BuildObjectMeta(name string) metav1.ObjectMeta {
  labels := BuildLabels()
  return metav1.ObjectMeta{
    Name: name,
    Labels: labels,
  }
}

// Create Service Object
func BuildServiceObject(name string, ports []corev1.ServicePort, selector map[string]string) *corev1.Service {
  service := &corev1.Service{
    ObjectMeta: BuildObjectMeta(name),
    Spec: corev1.ServiceSpec{
      Ports: ports,
      Selector: selector,
    },
  }

  return service
}

// Create ConfigMap objects
func BuildConfigmapObject(name string, dataMap map[string]string) *corev1.ConfigMap {
  return &corev1.ConfigMap{
    ObjectMeta: BuildObjectMeta(name),
    Data: dataMap,
  }
}

// Create StatefulSet
func BuildStatefulsetObject(name string) *appsv1.StatefulSet {
  return &appsv1.StatefulSet{
    ObjectMeta: BuildObjectMeta(name),
    Spec: BuildStatefulsetSpec()
  }

}

func BuildStatefulsetSpec(sts stsData) appsv1.StatefulSetSpec {
  return appsv1.StatefulSetSpec{
    Replicas: replicas,
    Selector: &metav1.LabelSelector{
      MatchLabels: BuildMatchLabels(),
    },
    Template: BuildPodTemplateSpec(sts),

  }
}

// Build the match label selector for the StatefulSet
func BuildMatchLabels() map[string]string {
  labels := make(map[string]string)
  labels["app"] = "icp-mongodb"
  labels["release"] = "mongodb"
  return labels
}

// Build the Pod spec for the statefulset
func BuildPodTemplateSpec(sts stsData) corev1.PodSpec{
  return corev1.PodSpec{
    Volumes: BuildVolumeArray(),
    InitContainers: BuildInitContainersArray(sts),
  }
}

// Build Array of Volumes for StatefulSet
func BuildVolumeArray() []corev1.Volume{
  var volumeArray [8]corev1.Volume
  volumeArray[0] = BuildConfigmapVolume("config", "icp-mongodb")
  volumeArray[1] = BuildConfigmapVolume("init", "icp-mongodb-init")
  volumeArray[2] = BuildConfigmapVolume("install", "icp-mongodb-install")
  volumeArray[3] = BuildConfigmapVolume("init", "icp-mongodb-init")
  volumeArray[4] = BuildSecretVolume("ca", "mongodb-root-ca-cert")
  volumeArray[5] = BuildSecretVolume("keydir", "icp-mongodb-keyfile")
  volumeArray[6] = BuildSecretVolume("tmp-mongodb")
  volumeArray[7] = BuildSecretVolume("tmp-metrics")

  return volumeArray
}

// Build a Configmap Volume
func BuildConfigmapVolume(volumeName string, configmapName string) corev1.Volume{
  return corev1.Volume{
    Name: volumeName,
    VolumeSource: corev1.VolumeSource{
      ConfigMap: &corev1.ConfigMapVolumeSource{
        LocalObjectReference: corev1.LocalObjectReference{
          Name: configmapName,
          DefaultMode: 0755,
        }
      }
    }
  }
}

// Build a Secret Volume
func BuildSecretVolume(volumeName string, secretName string) corev1.Volume{
  return corev1.Volume{
    Name: volumeName,
    VolumeSource: corev1.VolumeSource{
      Secret: &corev1.SecretVolumeSource{
        SecretName: secretName,
        DefaultMode: 0755,
      }
    }
  }
}

// Build an Empty Directory Volume
func BuildEmptyDirVolume(volumeName string) corev1.Volume{
  return corev1.Volume{
    Name: volumeName,
    VolumeSource: corev1.VolumeSource{
      EmptyDir: &corev1.EmptyDirVolumeSource{}
    }
  }
}

func BuildInitContainersArray(sts stsData) []corev1.Container{
  var initContainerArray []corev1.Container
  initContainerArray[0] = BuildInstallInitContainer(sts)
  initContainerArray[1] = BuildBootStrapInitContainer(sts)
}

func BuildContainer(containerName string, imageName string, commandArray []string, commandArgsArray []string, envVarArray []corev1.EnvVar, resources corev1.ResourceRequirements, volumeMountsArray []corev1.VolumeMount, liveProbe *corev1.Probe, readyProbe *corev1.Probe, securityContext *corev1.SecurityContext) corev1.Container {
  return corev1.Container{
    Name: containerName,
    Image: imageName,
    Command: commandArray,
    Args: commandArgsArray,
    Env: envVarArray,
    Resources: resources,
    VolumeMounts: volumeMountsArray,
    LivenessProbe: liveProbe,
    ReadinessProbe: readyProbe,
    ImagePullPolicy: corev1.PullPolicy.PullIfNotPresent,
    SecurityContext: securityContext,
  }
}

// Build the Install Init Container
func BuildInstallInitContainer(sts stsData) corev1.Container {
  commandArray := [1]string{"/install/install.sh"}
  argsArray := [2]string{"-on-start=/init/on-start.sh", "\"-service=icp-mongodb\""}
  volumeMountsArray := [8]corev1.VolumeMount(mongodbdirWDVolumeMount, configdirVolumeMount, configVolumeMount, installVolumeMount, keydirVolumeMount, caVolumeMount, mongodbdirDDVolumeMount, tmpVolumeMount)

  return BuildContainer("install", sts.ImageRepo + "/" + sts.InitImage, commandArray, argsArray, nil, mongoDBResourceLimit, volumeMountsArray, nil, nil, nil)
}

// Build the BootStrap Init Container
func BuildBootstrapInitContainer(sts stsData) corev1.Container {
  commandArray := [1]string{"/work-dir/peer-finder"}
  argsArray := [2]string{"--work-dir=/work-dir", "--config-dir=/data/configdb"}
  resource := BuildResourceRequirement(5*1024)
  volumeMountsArray := [5]corev1.VolumeMount(mongodbdirWDVolumeMount, configdirVolumeMount, initVolumeMount, mongodbdirDDVolumeMount, tmpVolumeMount)
  envVarArray := [9]corev1.EnvVar(podNamespaceEV, replicaSetEV, authEV, adminUserEV, adminPasswordEV, metricsEV, metricsUserEV, metricsPasswordEV, networkIPVersionEV)
  return BuildContainer("bootstrap", sts.ImageRepo + "/" + sts.BootstrapImage, commandArray, argsArray, envVarArray, mongoDBResourceLimit, volumeMountsArray, nil, nil, &mongoDBInitSecurityContext)
}

// Build MongoDB container
func BuildMongoDBContainer(sts stsData) corev1.Container {
  commandArray := [1]string{"mongod"}
  argsArray := [1]string{"--config=/data/configdb/mongod.conf"}
  envVarArray := [3]corev1.EnvVar(authEV, adminUserEV, adminPasswordEV)
  volumeMountsArray := [4]corev1.VolumeMount(mongodbdirWDVolumeMount, configdirVolumeMount, mongodbdirDDVolumeMount, tmpVolumeMount)
  mongodbProbe := &corev1.Probe{
    Handler: &corev1.Handler{
      Exec: &corev1.ExecAction{
        Command: []string("mongo", "--ssl", "--sslCAFile=/data/configdb/tls.crt", "--sslPEMKeyFile=/work-dir/mongo.pem", "--eval", "\"db.adminCommand('ping')\""),
      },
    },
    InitialDelaySeconds: 30,
    TimeoutSeconds: 5,
    PeriodSeconds: 10,
    SuccessThreshold: 1,
    FailureThreshold: 3,
  }

  return BuildContainer("icp-mongodb", sts.ImageRepo + "/" + sts.BootstrapImage, commandArray, argsArray, envVarArray, mongoDBResourceLimit, volumeMountsArray, mongodbProbe, mongodbProbe, &mongoDBInitSecurityContext)
}

// Build Metrics Container
func BuildMetricsContainer(sts stsData) corev1.Container {
  commandArray := [3]string("sh", "-ec", "/bin/mongodb_exporter")
  argsArray := [8]string("--mongodb.uri mongodb://$METRICS_USER:$METRICS_PASSWORD@localhost:27017", "--mongodb.tls", "--mongodb.tls-ca=/data/configdb/tls.crt", "--mongodb.tls-cert=/work-dir/mongo.pem", "--mongodb.socket-timeout=5s", "--mongodb.sync-timeout=1m", "--web.telemetry-path=/metrics", "--web.listen-address=:9216")
  envVarArray := [2]corev1.EnvVar(metricsUserEV, metricsPasswordEV)
  volumeMountsArray := [4]corev1.VolumeMount(mongodbdirWDVolumeMount, configdirVolumeMount, tmpMetricsVolumeMount)
  metricsProbe := &corev1.Probe{
    Handler: &corev1.Handler{
      Exec: &corev1.ExecAction{
        Command: []string("sh", "-ec", "/bin/mongodb_exporter", "--mongodb.uri mongodb://$METRICS_USER:$METRICS_PASSWORD@localhost:27017", "--mongodb.tls", "--mongodb.tls-ca=/data/configdb/tls.crt", "--mongodb.tls-cert=/work-dir/mongo.pem", "--test"),
      },
    },
    InitialDelaySeconds: 30,
    PeriodSeconds: 10,
  }

  return BuildContainer("metrics", sts.ImageRepo + "/" + sts.MetricsImage, commandArray, argsArray, envVarArray, BuildResourceRequirement(256), volumeMountsArray, metricsProbe, metricsProbe, &mongoDBInitSecurityContext)
}

// Build Resource Requirements
func BuildResourceRequirement(memorySizeInMB int) corev1.ResourceRequirements {
  return corev1.ResourceRequirements{
		Limits: map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceMemory: resource.NewQuantity(memorySizeInMB*1024*1024, resource.BinarySI)},
    }
}


// Make functions are used to make specific objects for the MongoDB deployment


func MakeICPMongoDBService() *corev1.Service {
  var port1 = corev1.ServicePort{
    Port: 27017,
    Protocol: corev1.ProtocolTCP,
    TargetPort: intstr.FromInt(27017),
  }
  var icpMongoDBServicePorts = []corev1.ServicePort {port1}

  var selectorMap = make(map[string]string)
  selectorMap["app"] = "icp-mongodb"
  selectorMap["release"] = "mongodb"

  return BuildServiceObject("icp-mongodb", icpMongoDBServicePorts, selectorMap)
}

func MakeMongoDBService() *corev1.Service {
  var port1 = corev1.ServicePort{
    Port: 27017,
    Protocol: corev1.ProtocolTCP,
    TargetPort: intstr.FromInt(27017),
  }

  var icpMongoDBServicePorts = []corev1.ServicePort {port1}

  var selectorMap = make(map[string]string)
  selectorMap["app"] = "icp-mongodb"
  selectorMap["release"] = "mongodb"

  return BuildServiceObject("mongodb", icpMongoDBServicePorts, selectorMap)
}

func MakeInitConfigmap() *corev1.ConfigMap {
  dataMap := make(map[string]string)
  dataMap["on-start.sh"] = onStartConstant
  return BuildConfigmapObject("icp-mongodb-init", dataMap)
}

func MakeInstallConfigmap() *corev1.ConfigMap {
  dataMap := make(map[string]string)
  dataMap["install.sh"] = installScriptConstant
  return BuildConfigmapObject("icp-mongodb-install", dataMap)
}

func MakeMongodConfConfigmap() *corev1.ConfigMap {
  dataMap := make(map[string]string)
  dataMap["mongod.conf"] = mongodConfConstant
  return BuildConfigmapObject("icp-mongodb", dataMap)
}
