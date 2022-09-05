package main

import (
	"fmt"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"time"
)

func main() {
	// 对一个项目只需要同一个tracer
	// 对于不同的api可以使用不同的span
	// 对于同一个api的多个步骤(函数)，可以使用span的嵌套
	cfg := &config.Configuration{
		ServiceName: "hello-world",
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
	defer func() { _ = closer.Close() }()

	span := tracer.StartSpan("go-grpc-web")
	time.Sleep(time.Second)
	defer span.Finish()
}
