package main

import (
	"context"
	"demo/proto"
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
)

func main() {
	cfg := &config.Configuration{
		ServiceName: "grpc-span",
		// 设置为const时，Param为0代表不采样，1代表一直采样
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		// 将信息发送到哪个服务器
		Reporter: &config.ReporterConfig{
			LogSpans:           true, // 发送详情
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	// 可以将tracer设为全局, 以后用的时候可以直接使用opentracing.StartSpan()
	// 这样就可以避免到处传递tracer了
	opentracing.SetGlobalTracer(tracer)
	defer func() { _ = closer.Close() }()

	conn, err := grpc.Dial("127.0.0.1:50051",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())), // 使用拦截器
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "world"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
