package ctx_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWithTimeout(t *testing.T) {
	ctx := context.Background()
	parent, parentCancel := context.WithTimeout(ctx, time.Second)
	sub, subCancel := context.WithTimeout(parent, 3*time.Second)

	// 父可以控制子，但子不能控制父
	// 即使在子 context 中重新设置了 timeout, 也是1秒钟就过期了
	go func() {
		<-sub.Done()
		fmt.Println("time out")
	}()

	time.Sleep(2 * time.Second)
	subCancel()
	parentCancel()
}

func TestTimeoutExample(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ch := make(chan struct{})
	go func() {
		slowBusiness()
		ch <- struct{}{}
	}()

	// 典型用法: 监听两个channel
	// 一个是正常业务的channel, 另一个Done返回的channel
	select {
	case <-ctx.Done():
		fmt.Println("timeout")
	case <-ch:
		fmt.Println("business end")
	}
}

func slowBusiness() {
	time.Sleep(2 * time.Second)
}
