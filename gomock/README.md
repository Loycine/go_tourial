## GoMock

```go
go get -u github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen

cd gomock
mockgen -source=db.go -destination=db_mock.go -package=main
```

如何编写可mock的代码
- mock作用的是接口，因此将依赖抽象为接口，而不是直接依赖具体的类。
- 不直接依赖实例，而是使用依赖注入降低耦合性。（依赖注入：给予调用方它需要的事物，通过参数传递，而不是让它自己去new）


### mock不能作用于GetFromDB内部，所以无法测试
```go
func GetFromDB(key string) int {
	db := NewDB()
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}
```

### 如果选择传参，就可以更容易传入Mock对象
```go
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}
```


## 编写xxx_test.go进行测试
```go
go test . -cover -v
```