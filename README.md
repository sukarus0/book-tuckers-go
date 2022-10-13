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

## chatper 16 - 패키지
`패키지` - Go 언어에서 코드를 묶는 가장 큰 단위  
main 패키지 + 외부 패키지  

- https://golang.org/pkg/
- https://github.com/avelino/awesome-go

### Pacakge 설치하기
import로 패키지를 포함시키면,  
go build를 통해서 빌드할 때 해당하는 패키지를 찾아서 포함한 다음  
실행 파일을 생성  

패키지를 찾는 세 가지 방법  
(1) 기본 제공 패키지: Go 설치 경로  
(2) 외부 저장소 패키지: 다운로드해서 GOPATH/pkg 디렉토리에 설치  
(3) 현재 모듈 아래 위치한 패키지 (현재 디렉토리 아래 있는 패키지)  
```go
import (
        "fmt" // (1) 기본 제공 패키지
        "goproject/usepkg/custompkg" // (2) 현재 모듈 패키지

        "github.com/guptarohit/asciigraph"  // (3) 외부 저장소 패키지
        "github.com/tuckersGo/musthaveGo/ch16/expkg" // (3) 외부 저장소 패키지
)
```

### Go 모듈
**go build를 하려면 반드시 Go 모듈 루트 폴더에 go.mod 파일이 있어야 합니다.**
`go.mod` : 모듈이름 + Go버전 + 필요한 외부패키지 명시되어 있음  
`go.sum` : 외부 저장소 패키지 버전 정보 담겨 있음  
go.mod와 go.sum파일을 통해 외부 패키지와 모듈 내 패키지를 합쳐서 실행 파일을 만듬.  

**Go 모듈은 `go mode init`명령을 통해 만들 수 있습니다.**
```go
go mod init [패키지명]
```

(1) `go mod init [패키지명]` --> (2) `go.mod` 생성 --> (3) `go mod tidy` -->  
(4) `go.sum` 생성  

**go mod tidy 명령은 Go 모듈에 필요한 패키지를 찾아서 다운로드해주고**  
**필요한 패키지 정보를 go.mod 파일과 go.sum 파일에 적어주게 됩니다.**  
*go.sum에는 패키지 위조 여부를 검사하기 위한 체크섬 결과가 담겨 있습니다.*  

<go.mod>
```go
module goproject/usepkg

go 1.19

require (
        github.com/guptarohit/asciigraph v0.5.5
        github.com/tuckersGo/musthaveGo/ch16/expkg v0.0.0-20210809125204-68bca0d80b54
)
```

```go
$ ./usepkg
This is custom package!
This is Github expkg Sample
 10.00 ┤        ╭╮
  9.00 ┤   ╭╮   ││
  8.00 ┤   ││ ╭╮││
  7.00 ┤   │╰╮││││╭╮
  6.00 ┤  ╭╯ │││││││ ╭
  5.00 ┤ ╭╯  ╰╯╰╯│││╭╯
  4.00 ┤╭╯       ││││
  3.00 ┼╯        ││││
  2.00 ┤         ╰╯╰╯
```

다운로드 받은 외부 패키지(asciigraph, expkg)는 `GOPATH/pak/mod`폴더에 저장.  
보통 `GOPATH=${HOME}/go` 임. 

### 패키지 초기화
패키지를 임포트하면 벌어지는 일  
(1) 컴파일러는 패키지 내 전역 변수를 초기화  
(2) 패키지에 init() 함수가 있다면 호출 *init()함수는 반드시 입력 매개변수가 없고 반환값도 없음*  
**만약 어떤 패키지의 초기화 함수인 init() 함수 기능만 사용하기를 원할 경우**  
**밑줄을 이용해서 임포트**  


