// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CmdbService) DeepCopyInto(out *CmdbService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CmdbService.
func (in *CmdbService) DeepCopy() *CmdbService {
	if in == nil {
		return nil
	}
	out := new(CmdbService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CmdbService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CmdbServiceList) DeepCopyInto(out *CmdbServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CmdbService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CmdbServiceList.
func (in *CmdbServiceList) DeepCopy() *CmdbServiceList {
	if in == nil {
		return nil
	}
	out := new(CmdbServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CmdbServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CmdbServiceSpec) DeepCopyInto(out *CmdbServiceSpec) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int32)
		**out = **in
	}
	in.Resource.DeepCopyInto(&out.Resource)
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]corev1.ServicePort, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CmdbServiceSpec.
func (in *CmdbServiceSpec) DeepCopy() *CmdbServiceSpec {
	if in == nil {
		return nil
	}
	out := new(CmdbServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CmdbServiceStatus) DeepCopyInto(out *CmdbServiceStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CmdbServiceStatus.
func (in *CmdbServiceStatus) DeepCopy() *CmdbServiceStatus {
	if in == nil {
		return nil
	}
	out := new(CmdbServiceStatus)
	in.DeepCopyInto(out)
	return out
}
