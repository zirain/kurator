/*
Copyright Kurator Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	clusterv1alpha1 "kurator.dev/kurator/pkg/client-go/generated/clientset/versioned/typed/cluster/v1alpha1"
	fleetv1alpha1 "kurator.dev/kurator/pkg/client-go/generated/clientset/versioned/typed/fleet/v1alpha1"
	infrastructurev1alpha1 "kurator.dev/kurator/pkg/client-go/generated/clientset/versioned/typed/infra/v1alpha1"
	telemetryv1alpha1 "kurator.dev/kurator/pkg/client-go/generated/clientset/versioned/typed/telemetry/v1alpha1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ClusterV1alpha1() clusterv1alpha1.ClusterV1alpha1Interface
	FleetV1alpha1() fleetv1alpha1.FleetV1alpha1Interface
	InfrastructureV1alpha1() infrastructurev1alpha1.InfrastructureV1alpha1Interface
	TelemetryV1alpha1() telemetryv1alpha1.TelemetryV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	clusterV1alpha1        *clusterv1alpha1.ClusterV1alpha1Client
	fleetV1alpha1          *fleetv1alpha1.FleetV1alpha1Client
	infrastructureV1alpha1 *infrastructurev1alpha1.InfrastructureV1alpha1Client
	telemetryV1alpha1      *telemetryv1alpha1.TelemetryV1alpha1Client
}

// ClusterV1alpha1 retrieves the ClusterV1alpha1Client
func (c *Clientset) ClusterV1alpha1() clusterv1alpha1.ClusterV1alpha1Interface {
	return c.clusterV1alpha1
}

// FleetV1alpha1 retrieves the FleetV1alpha1Client
func (c *Clientset) FleetV1alpha1() fleetv1alpha1.FleetV1alpha1Interface {
	return c.fleetV1alpha1
}

// InfrastructureV1alpha1 retrieves the InfrastructureV1alpha1Client
func (c *Clientset) InfrastructureV1alpha1() infrastructurev1alpha1.InfrastructureV1alpha1Interface {
	return c.infrastructureV1alpha1
}

// TelemetryV1alpha1 retrieves the TelemetryV1alpha1Client
func (c *Clientset) TelemetryV1alpha1() telemetryv1alpha1.TelemetryV1alpha1Interface {
	return c.telemetryV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.clusterV1alpha1, err = clusterv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.fleetV1alpha1, err = fleetv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.infrastructureV1alpha1, err = infrastructurev1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.telemetryV1alpha1, err = telemetryv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.clusterV1alpha1 = clusterv1alpha1.New(c)
	cs.fleetV1alpha1 = fleetv1alpha1.New(c)
	cs.infrastructureV1alpha1 = infrastructurev1alpha1.New(c)
	cs.telemetryV1alpha1 = telemetryv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
