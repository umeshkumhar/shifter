/*
copyright 2019 google llc
licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at
    http://www.apache.org/licenses/license-2.0
unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package openshift

import (
	"context"
	"log"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Get all OpenShift Daemon Sets by Namespace
func (c Openshift) GetAllDaemonSets(namespace string) (*v1.DaemonSetList, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.DaemonSetList{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get All OpenShift Daemon Sets By Namespace
	daemonSets, err := cluster.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// Error: Getting All OpenShift Daemon Sets By Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Daemon Sets from Namespace: '%s'. Error Text: '%s'. ", namespace, err.Error())
		return &v1.DaemonSetList{}, err
	} else {
		// Success: Getting All OpenShift Daemon Sets By Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Daemon Sets from Namespace: '%s'.", namespace)
		// Return Daemon Sets
		return daemonSets, err
	}

}

// Get OpenShift Daemon Set by Name from Namespace
func (c Openshift) GetDaemonSet(name string, namespace string) (*v1.DaemonSet, error) {
	// Uses KNative Structs
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		// Error: Getting Cluster Configuration
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Cluster Configuration")
		return &v1.DaemonSet{}, err
	} else {
		// Success: Getting Cluster Configuration
		log.Printf("🧰 ✅ SUCCESS: Getting Cluster Configuration")
	}

	// Get OpenShift Daemon Set By Name from Namespace
	daemonSet, err := cluster.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// Error: Getting OpenShift Daemon Set By Name & Namespace
		log.Printf("🧰 ❌ ERROR: Getting OpenShift Daemon Set with Name: '%s' from Namespace: '%s'. Error Text: '%s'. ", name, namespace, err.Error())
		return &v1.DaemonSet{}, err
	} else {
		// Success: Getting OpenShift Daemon Set By Name & Namespace
		log.Printf("🧰 ✅ SUCCESS: Getting OpenShift Daemon Set with Name: '%s' from Namespace: '%s'.", name, namespace)
		// Return Daemon Set
		return daemonSet, err
	}

}
