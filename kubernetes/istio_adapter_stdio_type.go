package kubernetes

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// stdio is one of the specific types defined by Istio as a Kubernetes extension.
// Preliminar rule is to define one file per type.
// types.go will collect common/shared types.
// This type is extracted from Istio Pilot as models used are not public and it doesn't make sense to fetch all
// Istio project as a dependency.
// Reference: https://github.com/istio/istio/blob/master/pilot/pkg/config/kube/crd/types.go

const stdios = "stdios"
const stdioType = "stdio"
const stdioLabel = "stdio"

// stdio is the generic Kubernetes API object wrapper
// stdio starts with lowercase as it maps a "kind":"stdio" Istio response.
type stdio struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               map[string]interface{} `json:"spec"`
}

// stdioList is the generic Kubernetes API list wrapper
// stdioList starts with lowercase as it maps a "kind":"stdioList" Istio response.
type stdioList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []stdio `json:"items"`
}

// GetSpec from a wrapper
func (in *stdio) GetSpec() map[string]interface{} {
	return in.Spec
}

// SetSpec for a wrapper
func (in *stdio) SetSpec(spec map[string]interface{}) {
	in.Spec = spec
}

// GetObjectMeta from a wrapper
func (in *stdio) GetObjectMeta() meta_v1.ObjectMeta {
	return in.ObjectMeta
}

// SetObjectMeta for a wrapper
func (in *stdio) SetObjectMeta(metadata meta_v1.ObjectMeta) {
	in.ObjectMeta = metadata
}

// GetItems from a wrapper
func (in *stdioList) GetItems() []IstioObject {
	out := make([]IstioObject, len(in.Items))
	for i := range in.Items {
		out[i] = &in.Items[i]
	}
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *stdio) DeepCopyInto(out *stdio) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteRule.
func (in *stdio) DeepCopy() *stdio {
	if in == nil {
		return nil
	}
	out := new(stdio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *stdio) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}

// DeepCopyIstioObject is an autogenerated deepcopy function, copying the receiver, creating a new IstioObject.
func (in *stdio) DeepCopyIstioObject() IstioObject {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *stdioList) DeepCopyInto(out *stdioList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]stdio, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleList.
func (in *stdioList) DeepCopy() *stdioList {
	if in == nil {
		return nil
	}
	out := new(stdioList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *stdioList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}
