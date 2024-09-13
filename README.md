# gomatcher
Matcher Collection for gomock

## Installation
Use go get.

```shell
go get github.com/bickyeric/gomatcher
```
Then import the validator package into your own code.

```golang
import "github.com/bickyeric/gomatcher"
```

## Working with Struct

### Struct Including

Validating half of the data is more effective than using `gomock.Any()` matcher.

```go
uc := mock.NewMockUserUsecase(ctrl)

// expect passed user to have field "Name" with value "John Doe"
uc.EXPECT().CreateUser(context.Background(), gomatcher.StructIncludes{
  "Name": "John Doe",
})

// this statement will pass
uc.CreateUser(context.Background(), usecase.User{
  ID:        123,
  Name:      "John Doe",
  CreatedAt: time.Now()
})
```

## How to Contribute
Just make a pull request...
