module netease.com/kubediag/event-metrics

go 1.16

require (
	cloud.google.com/go v0.56.0 // indirect
	github.com/Azure/go-autorest/autorest v0.10.0 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.8.3 // indirect
	github.com/go-logr/logr v0.1.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/googleapis/gnostic v0.4.0 // indirect
	github.com/gophercloud/gophercloud v0.10.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/procfs v0.0.11 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	go.uber.org/atomic v1.6.0 // indirect
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20200422194213-44a606286825 // indirect
	golang.org/x/sys v0.0.0-20200420163511-1957bb5e6d1f // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	golang.org/x/tools v0.0.0-20200422205258-72e4a01eba43 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	k8s.io/api v0.17.9
	k8s.io/apimachinery v0.17.9
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.0.0-00010101000000-000000000000
	sigs.k8s.io/structured-merge-diff v1.0.1-0.20191108220359-b1b620dd3f06 // indirect
)

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.0
	// Temporary until https://github.com/openshift/prom-label-proxy/pull/28 gets merged
	github.com/openshift/prom-label-proxy => github.com/vsliouniaev/prom-label-proxy v0.0.0-20200518104441-4fd7fe13454f
	k8s.io/api => k8s.io/api v0.16.15
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.15
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.5.14
)
