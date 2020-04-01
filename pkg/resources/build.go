//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package resources

import (
	//"reflect"
	//"strconv"

	//logf "sigs.k8s.io/controller-runtime/pkg/log"

	operatorv1alpha1 "github.com/IBM/ibm-mongodb-operator/pkg/apis/operator/v1alpha1"
	//certmgr "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1"
	//yaml "gopkg.in/yaml.v2"
	//appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	//rbacv1 "k8s.io/api/rbac/v1"
	//extv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/apimachinery/pkg/util/intstr"
)

//constant values for annotations
const productName = "IBM Cloud Platform Common Services"
const productID = "068a62892a1e4db39641342e592daa25"
const productVersion = "3.3.0"
const productMetric = "FREE"

//constant values for metadata labels
const meta_name = "icp-mongodb"
const instance = "icp-mongodb"
const label_version = "app.kubernetes.io/version: 4.0.12"
const component = "database"
const label_partOf = "app.kubernetes.io/part-of: common-services-cloud-pak"
const managedBy = "operator"
const release = "mongoDB"

func BuildConfigMap(instance *operatorv1alpha1.MongoDB, name string) (*corev1.ConfigMap, err) {
  labels := LabelsForMetadata()
  dataSection := make(map[string]string)

  dataSection["install.sh"] = installScript
  configMap := &corev1.ConfigMap{
    Object.metav1.ObjectMeta{
      Name: name,
      Labels: labels,
    },
    Data: dataSection,
  }

  return configMap, nil
}

// Helper function to generate the metadata labels
func LabelsForMetadata() map[string]string {
	return map[string]string{"app.kubernetes.io/name": meta_name, "app.kubernetes.io/component": component,
		"app.kubernetes.io/managed-by": managedBy, "app.kubernetes.io/instance": instance, "release": "mongodb"}

}
