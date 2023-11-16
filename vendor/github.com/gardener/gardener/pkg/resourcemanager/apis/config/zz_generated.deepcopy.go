//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package config

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentbaseconfig "k8s.io/component-base/config"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRDDeletionProtection) DeepCopyInto(out *CRDDeletionProtection) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRDDeletionProtection.
func (in *CRDDeletionProtection) DeepCopy() *CRDDeletionProtection {
	if in == nil {
		return nil
	}
	out := new(CRDDeletionProtection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConnection) DeepCopyInto(out *ClientConnection) {
	*out = *in
	out.ClientConnectionConfiguration = in.ClientConnectionConfiguration
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CacheResyncPeriod != nil {
		in, out := &in.CacheResyncPeriod, &out.CacheResyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConnection.
func (in *ClientConnection) DeepCopy() *ClientConnection {
	if in == nil {
		return nil
	}
	out := new(ClientConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSliceHintsWebhookConfig) DeepCopyInto(out *EndpointSliceHintsWebhookConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSliceHintsWebhookConfig.
func (in *EndpointSliceHintsWebhookConfig) DeepCopy() *EndpointSliceHintsWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(EndpointSliceHintsWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionValidation) DeepCopyInto(out *ExtensionValidation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionValidation.
func (in *ExtensionValidation) DeepCopy() *ExtensionValidation {
	if in == nil {
		return nil
	}
	out := new(ExtensionValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GarbageCollectorControllerConfig) DeepCopyInto(out *GarbageCollectorControllerConfig) {
	*out = *in
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GarbageCollectorControllerConfig.
func (in *GarbageCollectorControllerConfig) DeepCopy() *GarbageCollectorControllerConfig {
	if in == nil {
		return nil
	}
	out := new(GarbageCollectorControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSServer) DeepCopyInto(out *HTTPSServer) {
	*out = *in
	out.Server = in.Server
	out.TLS = in.TLS
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSServer.
func (in *HTTPSServer) DeepCopy() *HTTPSServer {
	if in == nil {
		return nil
	}
	out := new(HTTPSServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthControllerConfig) DeepCopyInto(out *HealthControllerConfig) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthControllerConfig.
func (in *HealthControllerConfig) DeepCopy() *HealthControllerConfig {
	if in == nil {
		return nil
	}
	out := new(HealthControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HighAvailabilityConfigWebhookConfig) DeepCopyInto(out *HighAvailabilityConfigWebhookConfig) {
	*out = *in
	if in.DefaultNotReadyTolerationSeconds != nil {
		in, out := &in.DefaultNotReadyTolerationSeconds, &out.DefaultNotReadyTolerationSeconds
		*out = new(int64)
		**out = **in
	}
	if in.DefaultUnreachableTolerationSeconds != nil {
		in, out := &in.DefaultUnreachableTolerationSeconds, &out.DefaultUnreachableTolerationSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HighAvailabilityConfigWebhookConfig.
func (in *HighAvailabilityConfigWebhookConfig) DeepCopy() *HighAvailabilityConfigWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(HighAvailabilityConfigWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressControllerSelector) DeepCopyInto(out *IngressControllerSelector) {
	*out = *in
	in.PodSelector.DeepCopyInto(&out.PodSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressControllerSelector.
func (in *IngressControllerSelector) DeepCopy() *IngressControllerSelector {
	if in == nil {
		return nil
	}
	out := new(IngressControllerSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletCSRApproverControllerConfig) DeepCopyInto(out *KubeletCSRApproverControllerConfig) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletCSRApproverControllerConfig.
func (in *KubeletCSRApproverControllerConfig) DeepCopy() *KubeletCSRApproverControllerConfig {
	if in == nil {
		return nil
	}
	out := new(KubeletCSRApproverControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesServiceHostWebhookConfig) DeepCopyInto(out *KubernetesServiceHostWebhookConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesServiceHostWebhookConfig.
func (in *KubernetesServiceHostWebhookConfig) DeepCopy() *KubernetesServiceHostWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(KubernetesServiceHostWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedResourceControllerConfig) DeepCopyInto(out *ManagedResourceControllerConfig) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	if in.AlwaysUpdate != nil {
		in, out := &in.AlwaysUpdate, &out.AlwaysUpdate
		*out = new(bool)
		**out = **in
	}
	if in.ManagedByLabelValue != nil {
		in, out := &in.ManagedByLabelValue, &out.ManagedByLabelValue
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedResourceControllerConfig.
func (in *ManagedResourceControllerConfig) DeepCopy() *ManagedResourceControllerConfig {
	if in == nil {
		return nil
	}
	out := new(ManagedResourceControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyControllerConfig) DeepCopyInto(out *NetworkPolicyControllerConfig) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	if in.NamespaceSelectors != nil {
		in, out := &in.NamespaceSelectors, &out.NamespaceSelectors
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IngressControllerSelector != nil {
		in, out := &in.IngressControllerSelector, &out.IngressControllerSelector
		*out = new(IngressControllerSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyControllerConfig.
func (in *NetworkPolicyControllerConfig) DeepCopy() *NetworkPolicyControllerConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeControllerConfig) DeepCopyInto(out *NodeControllerConfig) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	if in.Backoff != nil {
		in, out := &in.Backoff, &out.Backoff
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeControllerConfig.
func (in *NodeControllerConfig) DeepCopy() *NodeControllerConfig {
	if in == nil {
		return nil
	}
	out := new(NodeControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSchedulerNameWebhookConfig) DeepCopyInto(out *PodSchedulerNameWebhookConfig) {
	*out = *in
	if in.SchedulerName != nil {
		in, out := &in.SchedulerName, &out.SchedulerName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSchedulerNameWebhookConfig.
func (in *PodSchedulerNameWebhookConfig) DeepCopy() *PodSchedulerNameWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(PodSchedulerNameWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTopologySpreadConstraintsWebhookConfig) DeepCopyInto(out *PodTopologySpreadConstraintsWebhookConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTopologySpreadConstraintsWebhookConfig.
func (in *PodTopologySpreadConstraintsWebhookConfig) DeepCopy() *PodTopologySpreadConstraintsWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(PodTopologySpreadConstraintsWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectedTokenMountWebhookConfig) DeepCopyInto(out *ProjectedTokenMountWebhookConfig) {
	*out = *in
	if in.ExpirationSeconds != nil {
		in, out := &in.ExpirationSeconds, &out.ExpirationSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectedTokenMountWebhookConfig.
func (in *ProjectedTokenMountWebhookConfig) DeepCopy() *ProjectedTokenMountWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(ProjectedTokenMountWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceManagerConfiguration) DeepCopyInto(out *ResourceManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.SourceClientConnection.DeepCopyInto(&out.SourceClientConnection)
	if in.TargetClientConnection != nil {
		in, out := &in.TargetClientConnection, &out.TargetClientConnection
		*out = new(ClientConnection)
		(*in).DeepCopyInto(*out)
	}
	out.LeaderElection = in.LeaderElection
	in.Server.DeepCopyInto(&out.Server)
	if in.Debugging != nil {
		in, out := &in.Debugging, &out.Debugging
		*out = new(componentbaseconfig.DebuggingConfiguration)
		**out = **in
	}
	in.Controllers.DeepCopyInto(&out.Controllers)
	in.Webhooks.DeepCopyInto(&out.Webhooks)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceManagerConfiguration.
func (in *ResourceManagerConfiguration) DeepCopy() *ResourceManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ResourceManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceManagerControllerConfiguration) DeepCopyInto(out *ResourceManagerControllerConfiguration) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ResourceClass != nil {
		in, out := &in.ResourceClass, &out.ResourceClass
		*out = new(string)
		**out = **in
	}
	in.GarbageCollector.DeepCopyInto(&out.GarbageCollector)
	in.Health.DeepCopyInto(&out.Health)
	in.KubeletCSRApprover.DeepCopyInto(&out.KubeletCSRApprover)
	in.ManagedResource.DeepCopyInto(&out.ManagedResource)
	in.NetworkPolicy.DeepCopyInto(&out.NetworkPolicy)
	in.Node.DeepCopyInto(&out.Node)
	in.Secret.DeepCopyInto(&out.Secret)
	in.TokenInvalidator.DeepCopyInto(&out.TokenInvalidator)
	in.TokenRequestor.DeepCopyInto(&out.TokenRequestor)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceManagerControllerConfiguration.
func (in *ResourceManagerControllerConfiguration) DeepCopy() *ResourceManagerControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ResourceManagerControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceManagerWebhookConfiguration) DeepCopyInto(out *ResourceManagerWebhookConfiguration) {
	*out = *in
	out.CRDDeletionProtection = in.CRDDeletionProtection
	out.EndpointSliceHints = in.EndpointSliceHints
	out.ExtensionValidation = in.ExtensionValidation
	in.HighAvailabilityConfig.DeepCopyInto(&out.HighAvailabilityConfig)
	out.KubernetesServiceHost = in.KubernetesServiceHost
	in.PodSchedulerName.DeepCopyInto(&out.PodSchedulerName)
	out.PodTopologySpreadConstraints = in.PodTopologySpreadConstraints
	in.ProjectedTokenMount.DeepCopyInto(&out.ProjectedTokenMount)
	out.SeccompProfile = in.SeccompProfile
	in.SystemComponentsConfig.DeepCopyInto(&out.SystemComponentsConfig)
	out.TokenInvalidator = in.TokenInvalidator
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceManagerWebhookConfiguration.
func (in *ResourceManagerWebhookConfiguration) DeepCopy() *ResourceManagerWebhookConfiguration {
	if in == nil {
		return nil
	}
	out := new(ResourceManagerWebhookConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeccompProfileWebhookConfig) DeepCopyInto(out *SeccompProfileWebhookConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeccompProfileWebhookConfig.
func (in *SeccompProfileWebhookConfig) DeepCopy() *SeccompProfileWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(SeccompProfileWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretControllerConfig) DeepCopyInto(out *SecretControllerConfig) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretControllerConfig.
func (in *SecretControllerConfig) DeepCopy() *SecretControllerConfig {
	if in == nil {
		return nil
	}
	out := new(SecretControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfiguration) DeepCopyInto(out *ServerConfiguration) {
	*out = *in
	out.Webhooks = in.Webhooks
	if in.HealthProbes != nil {
		in, out := &in.HealthProbes, &out.HealthProbes
		*out = new(Server)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Server)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfiguration.
func (in *ServerConfiguration) DeepCopy() *ServerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemComponentsConfigWebhookConfig) DeepCopyInto(out *SystemComponentsConfigWebhookConfig) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodNodeSelector != nil {
		in, out := &in.PodNodeSelector, &out.PodNodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodTolerations != nil {
		in, out := &in.PodTolerations, &out.PodTolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemComponentsConfigWebhookConfig.
func (in *SystemComponentsConfigWebhookConfig) DeepCopy() *SystemComponentsConfigWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(SystemComponentsConfigWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSServer) DeepCopyInto(out *TLSServer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSServer.
func (in *TLSServer) DeepCopy() *TLSServer {
	if in == nil {
		return nil
	}
	out := new(TLSServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenInvalidatorControllerConfig) DeepCopyInto(out *TokenInvalidatorControllerConfig) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenInvalidatorControllerConfig.
func (in *TokenInvalidatorControllerConfig) DeepCopy() *TokenInvalidatorControllerConfig {
	if in == nil {
		return nil
	}
	out := new(TokenInvalidatorControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenInvalidatorWebhookConfig) DeepCopyInto(out *TokenInvalidatorWebhookConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenInvalidatorWebhookConfig.
func (in *TokenInvalidatorWebhookConfig) DeepCopy() *TokenInvalidatorWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(TokenInvalidatorWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenRequestorControllerConfig) DeepCopyInto(out *TokenRequestorControllerConfig) {
	*out = *in
	if in.ConcurrentSyncs != nil {
		in, out := &in.ConcurrentSyncs, &out.ConcurrentSyncs
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenRequestorControllerConfig.
func (in *TokenRequestorControllerConfig) DeepCopy() *TokenRequestorControllerConfig {
	if in == nil {
		return nil
	}
	out := new(TokenRequestorControllerConfig)
	in.DeepCopyInto(out)
	return out
}
