// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASOStatus) DeepCopyInto(out *ASOStatus) {
	*out = *in
	if in.RequestedAt != nil {
		in, out := &in.RequestedAt, &out.RequestedAt
		*out = (*in).DeepCopy()
	}
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASOStatus.
func (in *ASOStatus) DeepCopy() *ASOStatus {
	if in == nil {
		return nil
	}
	out := new(ASOStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlDatabase) DeepCopyInto(out *AzureSqlDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlDatabase.
func (in *AzureSqlDatabase) DeepCopy() *AzureSqlDatabase {
	if in == nil {
		return nil
	}
	out := new(AzureSqlDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureSqlDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlDatabaseList) DeepCopyInto(out *AzureSqlDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureSqlDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlDatabaseList.
func (in *AzureSqlDatabaseList) DeepCopy() *AzureSqlDatabaseList {
	if in == nil {
		return nil
	}
	out := new(AzureSqlDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureSqlDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlDatabaseSpec) DeepCopyInto(out *AzureSqlDatabaseSpec) {
	*out = *in
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(SqlDatabaseSku)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.ShortTermRetentionPolicy != nil {
		in, out := &in.ShortTermRetentionPolicy, &out.ShortTermRetentionPolicy
		*out = new(SQLDatabaseShortTermRetentionPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlDatabaseSpec.
func (in *AzureSqlDatabaseSpec) DeepCopy() *AzureSqlDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(AzureSqlDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlFailoverGroup) DeepCopyInto(out *AzureSqlFailoverGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlFailoverGroup.
func (in *AzureSqlFailoverGroup) DeepCopy() *AzureSqlFailoverGroup {
	if in == nil {
		return nil
	}
	out := new(AzureSqlFailoverGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureSqlFailoverGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlFailoverGroupList) DeepCopyInto(out *AzureSqlFailoverGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureSqlFailoverGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlFailoverGroupList.
func (in *AzureSqlFailoverGroupList) DeepCopy() *AzureSqlFailoverGroupList {
	if in == nil {
		return nil
	}
	out := new(AzureSqlFailoverGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureSqlFailoverGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlFailoverGroupSpec) DeepCopyInto(out *AzureSqlFailoverGroupSpec) {
	*out = *in
	if in.DatabaseList != nil {
		in, out := &in.DatabaseList, &out.DatabaseList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlFailoverGroupSpec.
func (in *AzureSqlFailoverGroupSpec) DeepCopy() *AzureSqlFailoverGroupSpec {
	if in == nil {
		return nil
	}
	out := new(AzureSqlFailoverGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlFirewallRule) DeepCopyInto(out *AzureSqlFirewallRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlFirewallRule.
func (in *AzureSqlFirewallRule) DeepCopy() *AzureSqlFirewallRule {
	if in == nil {
		return nil
	}
	out := new(AzureSqlFirewallRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureSqlFirewallRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlFirewallRuleList) DeepCopyInto(out *AzureSqlFirewallRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureSqlFirewallRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlFirewallRuleList.
func (in *AzureSqlFirewallRuleList) DeepCopy() *AzureSqlFirewallRuleList {
	if in == nil {
		return nil
	}
	out := new(AzureSqlFirewallRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureSqlFirewallRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlFirewallRuleSpec) DeepCopyInto(out *AzureSqlFirewallRuleSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlFirewallRuleSpec.
func (in *AzureSqlFirewallRuleSpec) DeepCopy() *AzureSqlFirewallRuleSpec {
	if in == nil {
		return nil
	}
	out := new(AzureSqlFirewallRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlServer) DeepCopyInto(out *AzureSqlServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlServer.
func (in *AzureSqlServer) DeepCopy() *AzureSqlServer {
	if in == nil {
		return nil
	}
	out := new(AzureSqlServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureSqlServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlServerList) DeepCopyInto(out *AzureSqlServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureSqlServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlServerList.
func (in *AzureSqlServerList) DeepCopy() *AzureSqlServerList {
	if in == nil {
		return nil
	}
	out := new(AzureSqlServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureSqlServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSqlServerSpec) DeepCopyInto(out *AzureSqlServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSqlServerSpec.
func (in *AzureSqlServerSpec) DeepCopy() *AzureSqlServerSpec {
	if in == nil {
		return nil
	}
	out := new(AzureSqlServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericResource) DeepCopyInto(out *GenericResource) {
	*out = *in
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericResource.
func (in *GenericResource) DeepCopy() *GenericResource {
	if in == nil {
		return nil
	}
	out := new(GenericResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericSpec) DeepCopyInto(out *GenericSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericSpec.
func (in *GenericSpec) DeepCopy() *GenericSpec {
	if in == nil {
		return nil
	}
	out := new(GenericSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLDatabaseShortTermRetentionPolicy) DeepCopyInto(out *SQLDatabaseShortTermRetentionPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLDatabaseShortTermRetentionPolicy.
func (in *SQLDatabaseShortTermRetentionPolicy) DeepCopy() *SQLDatabaseShortTermRetentionPolicy {
	if in == nil {
		return nil
	}
	out := new(SQLDatabaseShortTermRetentionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqlDatabaseSku) DeepCopyInto(out *SqlDatabaseSku) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqlDatabaseSku.
func (in *SqlDatabaseSku) DeepCopy() *SqlDatabaseSku {
	if in == nil {
		return nil
	}
	out := new(SqlDatabaseSku)
	in.DeepCopyInto(out)
	return out
}
