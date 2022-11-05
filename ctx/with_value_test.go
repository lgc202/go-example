package ctx_test

import (
	"context"
	"fmt"
	"testing"
)

func TestWithaValueString(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "my key", "my value")
	sub := context.WithValue(parent, "my key", "my new value")

	// 如果存放的是普通类型, parent 中拿不到 sub 中已经修改的值
	fmt.Println(parent.Value("my key"))
	fmt.Println(sub.Value("my key"))
}

func TestWithaValueMap(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "my map", map[string]string{"k1": "v1"})

	m := parent.Value("my map").(map[string]string)
	m["k1"] = "v2"
	sub := context.WithValue(parent, "my map", m)

	// 如果存放的是 map 类型, sub 中修改了 map 中的值, parent 中的值也会被修改
	// 但是这样违背了 context 不变性的原则，不到万不得已不要使用
	fmt.Println(parent.Value("my map"))
	fmt.Println(sub.Value("my map"))
}

func TestWithValuePointer(t *testing.T) {
	type valueType string
	var val valueType = "v1"

	ctx := context.Background()
	parent := context.WithValue(ctx, "my key", &val)

	pVal := parent.Value("my key").(*valueType)
	*pVal = "v2"
	sub := context.WithValue(parent, "my key", pVal)

	fmt.Println(*parent.Value("my key").(*valueType)) // 也是 v2
	fmt.Println(*sub.Value("my key").(*valueType))    // v2
}
