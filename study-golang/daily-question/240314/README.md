1. 变量作用域
2. main 函数中的 `p, err := foo()` 使用 `:=` 会新定义一个变量，全局变量 p 并不会赋值，所以在 `bar()` 中全局变量是 `nil`，所以会提示 `runtime error`