module github.com/openshift/cluster-kube-descheduler-operator

go 1.16

require (
	github.com/ghodss/yaml v1.0.0
	github.com/imdario/mergo v0.3.7
	github.com/openshift/api v0.0.0-20210706092853-b63d499a70ce
	github.com/openshift/build-machinery-go v0.0.0-20210712174854-1bb7fd1518d3
	github.com/openshift/library-go v0.0.0-20210331235027-66936e2fcc52
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/cobra v1.1.3
	k8s.io/api v0.22.0-beta.2
	k8s.io/apimachinery v0.22.0-beta.2
	k8s.io/client-go v0.22.0-beta.2
	k8s.io/code-generator v0.22.0-beta.2
	k8s.io/klog/v2 v2.9.0
	sigs.k8s.io/controller-tools v0.2.8
	sigs.k8s.io/descheduler v0.20.1-0.20210127064140-241f1325c968
)

replace (
	github.com/openshift/api => github.com/soltysh/api v0.0.0-20210719081803-9091ab00c164
	github.com/openshift/client-go => github.com/soltysh/client-go v0.0.0-20210719082425-f8fde3619384
	github.com/openshift/library-go => github.com/soltysh/library-go v0.0.0-20210719104342-c952f4e07d0b
)
