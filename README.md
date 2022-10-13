# Tuckers Go Programming

## chapter 03

같은 디렉토리에 있는 .go 파일은 디렉토리 이름명의 같은 패키지에 소속  
(Q).go 파일에서 package를 선언하면?  

Go 1.16 이후로는 Go 모듈이 기본 적용.  
따라서, Go 코드 빌드 전에 모듈 생성해야 함.  
```shell
go mod init <module_name>  
```

GOOS, GOARCH 환경변수로 운영체제와 아키텍처에 맞는 빌드 수행 가능  
```shell
GOOS=linux GOARCH=amd64 go build
```

## chapter 04 - 변수
```c
var a int = 3  // basic
var b int      // 초기값 생략, 기본값 대체
var c = 4      // 타입 생략, 우변 값의 타입으로 지정
d := 5         // 선언 대입문 :=, var생략, 타입생략
```

## chapter 05 - fmt

## chapter 06 - 연산자

## chatper 07 - 함수
```go
func Add(a int, b int) int {
	return a + b
}
```

멀티 값을 리턴할 수 있고,  
리턴할 변수 이름을 지정하면, return시 해당 변수값을 자동으로 리턴.  
다음 두 함수 비교.
```go
package main

import "fmt"

func Divide(a, b int) (int, bool) {
        if b == 0 {
                return 0, false
        }
        return a / b, true
}

func main() {
        c, success := Divide(9,3)
        fmt.Println(c, success)
        d, success := Divide(9,0)
        fmt.Println(d, success)
}
```

```go
200~package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
        if b == 0 {
                result = 0
                success = false
                return
        }
        result = a / b
        success = true
        return
}

func main() {
        c, success := Divide(9,3)
        fmt.Println(c, success)
        d, success := Divide(9,0)
        fmt.Println(d, success)
}
```

