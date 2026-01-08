# BasicLearning Note

## Slice Learning

取元素

```go
    package main
    import (
    	"fmt"
    )

    func main() {
		arr := []int{1, 3, 3}
        s1 := arr[0:1]
        // 注意这个是左开右闭
        for _, val := range s1 {
			fmt.Println(val)
		}
    }
```

## Method Learning

Golang里面没有**类**的概念，很自然的，那也没有继承这种说法；

那么Golang怎么实现这种功能？通过“**组合**”的方式
