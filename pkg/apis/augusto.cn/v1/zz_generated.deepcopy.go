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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	core_v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsCluster) DeepCopyInto(out *EsCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsCluster.
func (in *EsCluster) DeepCopy() *EsCluster {
	if in == nil {
		return nil
	}
	out := new(EsCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EsCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsClusterList) DeepCopyInto(out *EsClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EsCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsClusterList.
func (in *EsClusterList) DeepCopy() *EsClusterList {
	if in == nil {
		return nil
	}
	out := new(EsClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EsClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsClusterSpec) DeepCopyInto(out *EsClusterSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	in.Resource.DeepCopyInto(&out.Resource)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core_v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsClusterSpec.
func (in *EsClusterSpec) DeepCopy() *EsClusterSpec {
	if in == nil {
		return nil
	}
	out := new(EsClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsClusterStatus) DeepCopyInto(out *EsClusterStatus) {
	*out = *in
	if in.PodsStatus != nil {
		in, out := &in.PodsStatus, &out.PodsStatus
		*out = make([]EsInstanceStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ServiceStatus.DeepCopyInto(&out.ServiceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsClusterStatus.
func (in *EsClusterStatus) DeepCopy() *EsClusterStatus {
	if in == nil {
		return nil
	}
	out := new(EsClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsInstanceStatus) DeepCopyInto(out *EsInstanceStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsInstanceStatus.
func (in *EsInstanceStatus) DeepCopy() *EsInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(EsInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
