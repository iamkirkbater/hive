//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/openshift/hive/apis/hive/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSync) DeepCopyInto(out *ClusterSync) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSync.
func (in *ClusterSync) DeepCopy() *ClusterSync {
	if in == nil {
		return nil
	}
	out := new(ClusterSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSync) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSyncCondition) DeepCopyInto(out *ClusterSyncCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSyncCondition.
func (in *ClusterSyncCondition) DeepCopy() *ClusterSyncCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterSyncCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSyncLease) DeepCopyInto(out *ClusterSyncLease) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSyncLease.
func (in *ClusterSyncLease) DeepCopy() *ClusterSyncLease {
	if in == nil {
		return nil
	}
	out := new(ClusterSyncLease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSyncLease) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSyncLeaseList) DeepCopyInto(out *ClusterSyncLeaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSyncLease, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSyncLeaseList.
func (in *ClusterSyncLeaseList) DeepCopy() *ClusterSyncLeaseList {
	if in == nil {
		return nil
	}
	out := new(ClusterSyncLeaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSyncLeaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSyncLeaseSpec) DeepCopyInto(out *ClusterSyncLeaseSpec) {
	*out = *in
	in.RenewTime.DeepCopyInto(&out.RenewTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSyncLeaseSpec.
func (in *ClusterSyncLeaseSpec) DeepCopy() *ClusterSyncLeaseSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSyncLeaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSyncList) DeepCopyInto(out *ClusterSyncList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSync, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSyncList.
func (in *ClusterSyncList) DeepCopy() *ClusterSyncList {
	if in == nil {
		return nil
	}
	out := new(ClusterSyncList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSyncList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSyncSpec) DeepCopyInto(out *ClusterSyncSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSyncSpec.
func (in *ClusterSyncSpec) DeepCopy() *ClusterSyncSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSyncStatus) DeepCopyInto(out *ClusterSyncStatus) {
	*out = *in
	if in.SyncSets != nil {
		in, out := &in.SyncSets, &out.SyncSets
		*out = make([]SyncStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SelectorSyncSets != nil {
		in, out := &in.SelectorSyncSets, &out.SelectorSyncSets
		*out = make([]SyncStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterSyncCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FirstSuccessTime != nil {
		in, out := &in.FirstSuccessTime, &out.FirstSuccessTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSyncStatus.
func (in *ClusterSyncStatus) DeepCopy() *ClusterSyncStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterSyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FakeClusterInstall) DeepCopyInto(out *FakeClusterInstall) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FakeClusterInstall.
func (in *FakeClusterInstall) DeepCopy() *FakeClusterInstall {
	if in == nil {
		return nil
	}
	out := new(FakeClusterInstall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FakeClusterInstall) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FakeClusterInstallList) DeepCopyInto(out *FakeClusterInstallList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FakeClusterInstall, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FakeClusterInstallList.
func (in *FakeClusterInstallList) DeepCopy() *FakeClusterInstallList {
	if in == nil {
		return nil
	}
	out := new(FakeClusterInstallList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FakeClusterInstallList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FakeClusterInstallSpec) DeepCopyInto(out *FakeClusterInstallSpec) {
	*out = *in
	out.ImageSetRef = in.ImageSetRef
	out.ClusterDeploymentRef = in.ClusterDeploymentRef
	if in.ClusterMetadata != nil {
		in, out := &in.ClusterMetadata, &out.ClusterMetadata
		*out = new(v1.ClusterMetadata)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FakeClusterInstallSpec.
func (in *FakeClusterInstallSpec) DeepCopy() *FakeClusterInstallSpec {
	if in == nil {
		return nil
	}
	out := new(FakeClusterInstallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FakeClusterInstallStatus) DeepCopyInto(out *FakeClusterInstallStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.ClusterInstallCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FakeClusterInstallStatus.
func (in *FakeClusterInstallStatus) DeepCopy() *FakeClusterInstallStatus {
	if in == nil {
		return nil
	}
	out := new(FakeClusterInstallStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncResourceReference) DeepCopyInto(out *SyncResourceReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncResourceReference.
func (in *SyncResourceReference) DeepCopy() *SyncResourceReference {
	if in == nil {
		return nil
	}
	out := new(SyncResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncStatus) DeepCopyInto(out *SyncStatus) {
	*out = *in
	if in.ResourcesToDelete != nil {
		in, out := &in.ResourcesToDelete, &out.ResourcesToDelete
		*out = make([]SyncResourceReference, len(*in))
		copy(*out, *in)
	}
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	if in.FirstSuccessTime != nil {
		in, out := &in.FirstSuccessTime, &out.FirstSuccessTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncStatus.
func (in *SyncStatus) DeepCopy() *SyncStatus {
	if in == nil {
		return nil
	}
	out := new(SyncStatus)
	in.DeepCopyInto(out)
	return out
}
