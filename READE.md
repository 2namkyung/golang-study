# [ Golang ]

## Package

- Golang 필수 디렉토리 : bin / pkg / src

- 패키지를 import 할 때,
  Go 컴파일러는 GOROOT 혹은 GOPATH 환경변수를 검색한다

- 표준 패키지 : GOROOT/pkg
- 사용자 패키지/3rd Party 패키지 : GOPATH/pkg

- Package Scope
- 패키지 내에는 함수, 구조체, 인터페이스, 메소드 등이 존재  
  -> 이름의 첫 문자를 대문자로 시작하면 : Public  
  -> 이름의 첫 문자를 소문자로 시작하면 : Private  
  -> private은 패키지 내부에서만 사용할 수 있다

- init() 함수

  - 패키지 실행 시 처음으로 호출되는 init() 함수 작성 가능
    ```
    func init(){
    pop = make(map[string]string) // 패키지 로드 시 map을 초기화한다
    }
    ```

- alias

  - 패키지를 import하면서 패키지 안의 init() 함수만을 호출하는 경우  
    -> import _ "other/xlib" -> _ 라는 alias
  - 패키지 이름이 동일하지만, 서로 다른 버전이거나 서로 다른 위치에서 로딩하고자 할 때 alias로 구분 가능

- 사용자 정의 패키지 생성

  - 일반적으로 Directory를 생성하고 .go 파일들을 만들어 구성
  - 서브 디렉토리 안에 있는 .go 파일들은 동일한 패키지명을 가진다
  - 패키지명은 해당 디렉토리의 이름과 같게 한다

- Standard I/O

  - fmt.Fprintln(os.Stdout, s) == fmt.Println(s)
  - fmt.Fscan(os.Stdin, &s) == fmt.Scan(&s)

- GO111MODULE

  - off : Go Module 기능 사용하지 않는다
  - on : Go Module 기능을 사용한다
  - auto : 현재 디렉터리가 $GOPATH 외부에 있고 go.mod 파일이 포함된 경우 모듈을 사용
    그렇지 않으면 $GOPATH의 방법으로 사용한다
