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



