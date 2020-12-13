module github.com/apache/dubbo-go-samples

require (
	github.com/apache/dubbo-go v1.5.4
	github.com/apache/dubbo-go-hessian2 v1.8.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/dubbogo/gost v1.10.0
	github.com/emicklei/go-restful/v3 v3.0.0
	github.com/golang/protobuf v1.4.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5 // indirect
	github.com/openzipkin/zipkin-go v0.2.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.1.0
	github.com/stretchr/testify v1.6.1
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0 // indirect
)

go 1.13

replace github.com/apache/dubbo-go v1.5.4 => ../../../dubbo-go
