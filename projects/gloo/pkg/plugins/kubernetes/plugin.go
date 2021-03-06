package kubernetes

import (
	"fmt"
	"net/url"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"k8s.io/client-go/kubernetes"
)

type plugin struct {
	kube kubernetes.Interface

	kubeShareFactory KubePluginSharedFactory

	UpstreamConverter UpstreamConverter
}

func (p *plugin) Resolve(u *v1.Upstream) (*url.URL, error) {
	kubeSpec, ok := u.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_Kube)
	if !ok {
		return nil, nil
	}

	return url.Parse(fmt.Sprintf("tcp://%v.%v.svc.cluster.local:%v", kubeSpec.Kube.ServiceName, kubeSpec.Kube.ServiceNamespace, kubeSpec.Kube.ServicePort))
}

func NewPlugin(kube kubernetes.Interface) plugins.Plugin {
	return &plugin{
		kube:              kube,
		UpstreamConverter: DefaultUpstreamConverter(),
	}
}

func (p *plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoyapi.Cluster) error {
	// not ours
	_, ok := in.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_Kube)
	if !ok {
		return nil
	}

	// configure the cluster to use EDS:ADS and call it a day
	out.Type = envoyapi.Cluster_EDS
	out.EdsClusterConfig = &envoyapi.Cluster_EdsClusterConfig{
		EdsConfig: &envoycore.ConfigSource{
			ConfigSourceSpecifier: &envoycore.ConfigSource_Ads{
				Ads: &envoycore.AggregatedConfigSource{},
			},
		},
	}
	return nil
}
