// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&Cluster{}, func(obj interface{}) { SetObjectDefaults_Cluster(obj.(*Cluster)) })
	scheme.AddTypeDefaultingFunc(&ClusterList{}, func(obj interface{}) { SetObjectDefaults_ClusterList(obj.(*ClusterList)) })
	scheme.AddTypeDefaultingFunc(&ClusterProviderConfigSpec{}, func(obj interface{}) { SetObjectDefaults_ClusterProviderConfigSpec(obj.(*ClusterProviderConfigSpec)) })
	scheme.AddTypeDefaultingFunc(&Machine{}, func(obj interface{}) { SetObjectDefaults_Machine(obj.(*Machine)) })
	scheme.AddTypeDefaultingFunc(&MachineList{}, func(obj interface{}) { SetObjectDefaults_MachineList(obj.(*MachineList)) })
	scheme.AddTypeDefaultingFunc(&MachineSet{}, func(obj interface{}) { SetObjectDefaults_MachineSet(obj.(*MachineSet)) })
	scheme.AddTypeDefaultingFunc(&MachineSetList{}, func(obj interface{}) { SetObjectDefaults_MachineSetList(obj.(*MachineSetList)) })
	return nil
}

func SetObjectDefaults_Cluster(in *Cluster) {
	SetDefaults_ClusterSpec(&in.Spec)
}

func SetObjectDefaults_ClusterList(in *ClusterList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Cluster(a)
	}
}

func SetObjectDefaults_ClusterProviderConfigSpec(in *ClusterProviderConfigSpec) {
	SetDefaults_ClusterSpec(&in.ClusterSpec)
}

func SetObjectDefaults_Machine(in *Machine) {
	SetDefaults_MachineSpec(&in.Spec)
}

func SetObjectDefaults_MachineList(in *MachineList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Machine(a)
	}
}

func SetObjectDefaults_MachineSet(in *MachineSet) {
	SetDefaults_MachineSetSpec(&in.Spec)
}

func SetObjectDefaults_MachineSetList(in *MachineSetList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_MachineSet(a)
	}
}
