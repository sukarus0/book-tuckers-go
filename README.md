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
```go
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

## chatper 08 - 상수
(Q) `rune`?

`iota`
```go
const (
	Red	int = iota  // 0
	Blue	int = iota  // 1
	Green	int = iota  // 2
)
```

## chapter 09 - if
```go
if 초기문; 조건문 {
  문장
}
```


## chapter 10 - switch
break 없음  
fallthrough 키워드로 다음 case문 실행 가능.  


## chatper 11 - for

## chapter 12 - 배열
```go
# var 변수명 [요소개수]타입
var t [5]float64
```
```go
var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
days := [3]string{"monday", "tuseday", "wednesday"}
x := [...]int{10, 20, 30}
```

`range` keyword로 배열 순회
```go
index := [...]int{10, 20, 30}

for i, v := range index {
	fmt.Println(i, v)
}
```

## chapter 13 - 구조체

## chapter 14 - 포인터

## chapter 15 - 문자열
### UTF-8
Go는 UTF-8 문자코드를 표준 문자코드로 사용.  
Go 언어 창시자인 롭 파이크와 켄 톰슨이 고안한 문자.  
UTF-16이 한 문자에 2바이트를 고정 사용하는 것과 달리,  
UTF-8은 자주 사용되는 영문자, 숫자, 일부 특수 문자를 1바이트로 표현하고,  
그와 다른 문자들은 2~3바이트로 표현.  
**영문자, 숫자 등을 1바이트로 표현해 UTF-16에 비해 크기를 절약할 수 있고**  
**ANSI 코드와 1:1 대응이 되어 ANSI로 바로 변환된다는 장점**

### rune 타입
UTF-8은 한 글자가 1~3바이트 크기이기 때문에 UTF-8 문자값을 가지려면 3비이트 필요.  
Go 언어 기본 타입에서 3바이트 정수 타입은 제공되지 않기 때문에 rune 타입은  
4바이트 정수 타입인 int 32 타입의 별칭 타입  
```go
type run int32 // rune 타입과 int32는 이름만 다를 뿐 같은 타입입니다.
```

*string type을 rune타입으로 변환해, 문자열 길이를 계산할 수 있다*

