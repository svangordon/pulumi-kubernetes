// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ServiceSpec describes the attributes that a user creates on a service.
type ServiceSpec struct {
	// clusterIP is the IP address of the service and is usually assigned randomly by the master. If an
	// address is specified manually and is not in use by others, it will be allocated to the service;
	// otherwise, creation of the service will fail. This field can not be changed through updates.
	// Valid values are "None", empty string (""), or a valid IP address. "None" can be specified for
	// headless services when proxying is not required. Only applies to types ClusterIP, NodePort, and
	// LoadBalancer. Ignored if type is ExternalName. More info:
	// https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	ClusterIP *string `pulumi:"clusterIP"`

	// externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic
	// for this service.  These IPs are not managed by Kubernetes.  The user is responsible for
	// ensuring that traffic arrives at a node with this IP.  A common example is external
	// load-balancers that are not part of the Kubernetes system.
	ExternalIPs []string `pulumi:"externalIPs"`

	// externalName is the external reference that kubedns or equivalent will return as a CNAME record
	// for this service. No proxying will be involved. Must be a valid RFC-1123 hostname
	// (https://tools.ietf.org/html/rfc1123) and requires Type to be ExternalName.
	ExternalName *string `pulumi:"externalName"`

	// externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or
	// cluster-wide endpoints. "Local" preserves the client source IP and avoids a second hop for
	// LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading.
	// "Cluster" obscures the client source IP and may cause a second hop to another node, but should
	// have good overall load-spreading.
	ExternalTrafficPolicy *string `pulumi:"externalTrafficPolicy"`

	// healthCheckNodePort specifies the healthcheck nodePort for the service. If not specified,
	// HealthCheckNodePort is created by the service api backend with the allocated nodePort. Will use
	// user-specified nodePort value if specified by the client. Only effects when Type is set to
	// LoadBalancer and ExternalTrafficPolicy is set to Local.
	HealthCheckNodePort *int `pulumi:"healthCheckNodePort"`

	// ipFamily specifies whether this Service has a preference for a particular IP family (e.g. IPv4
	// vs. IPv6).  If a specific IP family is requested, the clusterIP field will be allocated from
	// that family, if it is available in the cluster.  If no IP family is requested, the cluster's
	// primary IP family will be used. Other IP fields (loadBalancerIP, loadBalancerSourceRanges,
	// externalIPs) and controllers which allocate external load-balancers should use the same IP
	// family.  Endpoints for this Service will be of this family.  This field is immutable after
	// creation. Assigning a ServiceIPFamily not available in the cluster (e.g. IPv6 in IPv4 only
	// cluster) is an error condition and will fail during clusterIP assignment.
	IpFamily *string `pulumi:"ipFamily"`

	// Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified
	// in this field. This feature depends on whether the underlying cloud-provider supports specifying
	// the loadBalancerIP when a load balancer is created. This field will be ignored if the
	// cloud-provider does not support the feature.
	LoadBalancerIP *string `pulumi:"loadBalancerIP"`

	// If specified and supported by the platform, this will restrict traffic through the
	// cloud-provider load-balancer will be restricted to the specified client IPs. This field will be
	// ignored if the cloud-provider does not support the feature." More info:
	// https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/
	LoadBalancerSourceRanges []string `pulumi:"loadBalancerSourceRanges"`

	// The list of ports that are exposed by this service. More info:
	// https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Ports []ServicePort `pulumi:"ports"`

	// publishNotReadyAddresses, when set to true, indicates that DNS implementations must publish the
	// notReadyAddresses of subsets for the Endpoints associated with the Service. The default value is
	// false. The primary use case for setting this field is to use a StatefulSet's Headless Service to
	// propagate SRV records for its Pods without respect to their readiness for purpose of peer
	// discovery.
	PublishNotReadyAddresses *bool `pulumi:"publishNotReadyAddresses"`

	// Route service traffic to pods with label keys and values matching this selector. If empty or not
	// present, the service is assumed to have an external process managing its endpoints, which
	// Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored
	// if type is ExternalName. More info:
	// https://kubernetes.io/docs/concepts/services-networking/service/
	Selector map[string]string `pulumi:"selector"`

	// Supports "ClientIP" and "None". Used to maintain session affinity. Enable client IP based
	// session affinity. Must be ClientIP or None. Defaults to None. More info:
	// https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	SessionAffinity *string `pulumi:"sessionAffinity"`

	// sessionAffinityConfig contains the configurations of session affinity.
	SessionAffinityConfig *SessionAffinityConfig `pulumi:"sessionAffinityConfig"`

	// type determines how the Service is exposed. Defaults to ClusterIP. Valid options are
	// ExternalName, ClusterIP, NodePort, and LoadBalancer. "ExternalName" maps to the specified
	// externalName. "ClusterIP" allocates a cluster-internal IP address for load-balancing to
	// endpoints. Endpoints are determined by the selector or if that is not specified, by manual
	// construction of an Endpoints object. If clusterIP is "None", no virtual IP is allocated and the
	// endpoints are published as a set of endpoints rather than a stable IP. "NodePort" builds on
	// ClusterIP and allocates a port on every node which routes to the clusterIP. "LoadBalancer"
	// builds on NodePort and creates an external load-balancer (if supported in the current cloud)
	// which routes to the clusterIP. More info:
	// https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	Type *string `pulumi:"type"`

}

var _ServiceSpecType = reflect.TypeOf((*ServiceSpec)(nil)).Elem()

// ServiceSpecInput represents an input type that resolves to a ServiceSpec.
type ServiceSpecInput interface {
	ElementType() reflect.Type

	ToServiceSpecOutput() ServiceSpecOutput
	ToServiceSpecOutputWithContext(ctx context.Context) ServiceSpecOutput
}

// ServiceSpecArgs is a ServiceSpecInput whose fields are all Input types.
type ServiceSpecArgs struct {
	// clusterIP is the IP address of the service and is usually assigned randomly by the master. If an
	// address is specified manually and is not in use by others, it will be allocated to the service;
	// otherwise, creation of the service will fail. This field can not be changed through updates.
	// Valid values are "None", empty string (""), or a valid IP address. "None" can be specified for
	// headless services when proxying is not required. Only applies to types ClusterIP, NodePort, and
	// LoadBalancer. Ignored if type is ExternalName. More info:
	// https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	ClusterIP pulumi.StringInput `pulumi:"clusterIP"`

	// externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic
	// for this service.  These IPs are not managed by Kubernetes.  The user is responsible for
	// ensuring that traffic arrives at a node with this IP.  A common example is external
	// load-balancers that are not part of the Kubernetes system.
	ExternalIPs pulumi.StringArrayInput `pulumi:"externalIPs"`

	// externalName is the external reference that kubedns or equivalent will return as a CNAME record
	// for this service. No proxying will be involved. Must be a valid RFC-1123 hostname
	// (https://tools.ietf.org/html/rfc1123) and requires Type to be ExternalName.
	ExternalName pulumi.StringInput `pulumi:"externalName"`

	// externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or
	// cluster-wide endpoints. "Local" preserves the client source IP and avoids a second hop for
	// LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading.
	// "Cluster" obscures the client source IP and may cause a second hop to another node, but should
	// have good overall load-spreading.
	ExternalTrafficPolicy pulumi.StringInput `pulumi:"externalTrafficPolicy"`

	// healthCheckNodePort specifies the healthcheck nodePort for the service. If not specified,
	// HealthCheckNodePort is created by the service api backend with the allocated nodePort. Will use
	// user-specified nodePort value if specified by the client. Only effects when Type is set to
	// LoadBalancer and ExternalTrafficPolicy is set to Local.
	HealthCheckNodePort pulumi.IntInput `pulumi:"healthCheckNodePort"`

	// ipFamily specifies whether this Service has a preference for a particular IP family (e.g. IPv4
	// vs. IPv6).  If a specific IP family is requested, the clusterIP field will be allocated from
	// that family, if it is available in the cluster.  If no IP family is requested, the cluster's
	// primary IP family will be used. Other IP fields (loadBalancerIP, loadBalancerSourceRanges,
	// externalIPs) and controllers which allocate external load-balancers should use the same IP
	// family.  Endpoints for this Service will be of this family.  This field is immutable after
	// creation. Assigning a ServiceIPFamily not available in the cluster (e.g. IPv6 in IPv4 only
	// cluster) is an error condition and will fail during clusterIP assignment.
	IpFamily pulumi.StringInput `pulumi:"ipFamily"`

	// Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified
	// in this field. This feature depends on whether the underlying cloud-provider supports specifying
	// the loadBalancerIP when a load balancer is created. This field will be ignored if the
	// cloud-provider does not support the feature.
	LoadBalancerIP pulumi.StringInput `pulumi:"loadBalancerIP"`

	// If specified and supported by the platform, this will restrict traffic through the
	// cloud-provider load-balancer will be restricted to the specified client IPs. This field will be
	// ignored if the cloud-provider does not support the feature." More info:
	// https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/
	LoadBalancerSourceRanges pulumi.StringArrayInput `pulumi:"loadBalancerSourceRanges"`

	// The list of ports that are exposed by this service. More info:
	// https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Ports ServicePortArrayInput `pulumi:"ports"`

	// publishNotReadyAddresses, when set to true, indicates that DNS implementations must publish the
	// notReadyAddresses of subsets for the Endpoints associated with the Service. The default value is
	// false. The primary use case for setting this field is to use a StatefulSet's Headless Service to
	// propagate SRV records for its Pods without respect to their readiness for purpose of peer
	// discovery.
	PublishNotReadyAddresses pulumi.BoolInput `pulumi:"publishNotReadyAddresses"`

	// Route service traffic to pods with label keys and values matching this selector. If empty or not
	// present, the service is assumed to have an external process managing its endpoints, which
	// Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored
	// if type is ExternalName. More info:
	// https://kubernetes.io/docs/concepts/services-networking/service/
	Selector pulumi.StringMapInput `pulumi:"selector"`

	// Supports "ClientIP" and "None". Used to maintain session affinity. Enable client IP based
	// session affinity. Must be ClientIP or None. Defaults to None. More info:
	// https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	SessionAffinity pulumi.StringInput `pulumi:"sessionAffinity"`

	// sessionAffinityConfig contains the configurations of session affinity.
	SessionAffinityConfig SessionAffinityConfigInput `pulumi:"sessionAffinityConfig"`

	// type determines how the Service is exposed. Defaults to ClusterIP. Valid options are
	// ExternalName, ClusterIP, NodePort, and LoadBalancer. "ExternalName" maps to the specified
	// externalName. "ClusterIP" allocates a cluster-internal IP address for load-balancing to
	// endpoints. Endpoints are determined by the selector or if that is not specified, by manual
	// construction of an Endpoints object. If clusterIP is "None", no virtual IP is allocated and the
	// endpoints are published as a set of endpoints rather than a stable IP. "NodePort" builds on
	// ClusterIP and allocates a port on every node which routes to the clusterIP. "LoadBalancer"
	// builds on NodePort and creates an external load-balancer (if supported in the current cloud)
	// which routes to the clusterIP. More info:
	// https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	Type pulumi.StringInput `pulumi:"type"`

}

func (a ServiceSpecArgs) ElementType() reflect.Type {
	return _ServiceSpecType
}

func (a ServiceSpecArgs) ToServiceSpecOutput() ServiceSpecOutput {
	return pulumi.ToOutput(a).(ServiceSpecOutput)
}

func (a ServiceSpecArgs) ToServiceSpecOutputWithContext(ctx context.Context) ServiceSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServiceSpecOutput)
}

// ServiceSpecOutput is an output type that resolves to a Input.
type ServiceSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ServiceSpecOutput{}) }

func (ServiceSpecOutput) ElementType() reflect.Type {
	return _ServiceSpecType
}

func (o ServiceSpecOutput) ClusterIP() pulumi.StringOutput {
	return o.Apply(func(v ServiceSpec) *string {
		return v.ClusterIP
	}).(pulumi.StringOutput)
}

func (o ServiceSpecOutput) ExternalIPs() pulumi.StringArrayOutput {
	return o.Apply(func(v ServiceSpec) []string {
		return v.ExternalIPs
	}).(pulumi.StringArrayOutput)
}

func (o ServiceSpecOutput) ExternalName() pulumi.StringOutput {
	return o.Apply(func(v ServiceSpec) *string {
		return v.ExternalName
	}).(pulumi.StringOutput)
}

func (o ServiceSpecOutput) ExternalTrafficPolicy() pulumi.StringOutput {
	return o.Apply(func(v ServiceSpec) *string {
		return v.ExternalTrafficPolicy
	}).(pulumi.StringOutput)
}

func (o ServiceSpecOutput) HealthCheckNodePort() pulumi.IntOutput {
	return o.Apply(func(v ServiceSpec) *int {
		return v.HealthCheckNodePort
	}).(pulumi.IntOutput)
}

func (o ServiceSpecOutput) IpFamily() pulumi.StringOutput {
	return o.Apply(func(v ServiceSpec) *string {
		return v.IpFamily
	}).(pulumi.StringOutput)
}

func (o ServiceSpecOutput) LoadBalancerIP() pulumi.StringOutput {
	return o.Apply(func(v ServiceSpec) *string {
		return v.LoadBalancerIP
	}).(pulumi.StringOutput)
}

func (o ServiceSpecOutput) LoadBalancerSourceRanges() pulumi.StringArrayOutput {
	return o.Apply(func(v ServiceSpec) []string {
		return v.LoadBalancerSourceRanges
	}).(pulumi.StringArrayOutput)
}

func (o ServiceSpecOutput) Ports() ServicePortArrayOutput {
	return o.Apply(func(v ServiceSpec) []ServicePort {
		return v.Ports
	}).(ServicePortArrayOutput)
}

func (o ServiceSpecOutput) PublishNotReadyAddresses() pulumi.BoolOutput {
	return o.Apply(func(v ServiceSpec) *bool {
		return v.PublishNotReadyAddresses
	}).(pulumi.BoolOutput)
}

func (o ServiceSpecOutput) Selector() pulumi.StringMapOutput {
	return o.Apply(func(v ServiceSpec) map[string]string {
		return v.Selector
	}).(pulumi.StringMapOutput)
}

func (o ServiceSpecOutput) SessionAffinity() pulumi.StringOutput {
	return o.Apply(func(v ServiceSpec) *string {
		return v.SessionAffinity
	}).(pulumi.StringOutput)
}

func (o ServiceSpecOutput) SessionAffinityConfig() SessionAffinityConfigOutput {
	return o.Apply(func(v ServiceSpec) *SessionAffinityConfig {
		return v.SessionAffinityConfig
	}).(SessionAffinityConfigOutput)
}

func (o ServiceSpecOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v ServiceSpec) *string {
		return v.Type
	}).(pulumi.StringOutput)
}
