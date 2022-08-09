# icl2-external-server-sample

## 개요
i Connect Live 2.0(이하 iCL2.0) 서비스 자체 인증 방식(external)을 위한 서버에 대한 예제들이 존재하는 Repository 입니다.

iCL2.0 인증 방식에는 커넥트라이브 인증 방식(internal)과 서비스 자체 인증 방식(external)을 지원하며 원하는 인증방식 중 하나를 선택해 서비스를 이용할 수 있습니다.

서비스 자체 인증 방식(external)이란, 서비스 제공자의 서버에서 자체 발급한 토큰을 이용해 iCL 2.0 서비스를 이용할 수 있는 인증 방식입니다.

서비스 자체 인증 방식(external)을 사용하기 위해서는 서비스 서버에서 두 가지 기능을 제공해야 합니다. REST API로 제공한다면 다음과 같이 API를 작성해볼 수 있습니다.

- /generate
  - 클라이언트가 iCL2.0으로 보낼 토큰을 생성하는 API입니다.
- /validate
  - iCL2.0 서버가 generate에서 생성한 토큰을 검증할 API입니다.

위 두 API를 구현한 서비스 자체 인증 방식(external)예제서버를 사용하여 쉽고 빠르게 사용해보실 수 있습니다.

예제서버는 Go, Node.js, Python, SpringBoot/Java 로 제공됩니다.
## 세부정보
서비스 자체 인증 방식(external) 서버 예제는 JWT를 사용한 Token 인증 방식을 제공하고 있습니다.

JWT에 관련된 자세한 정보는 [jwt.io](https://jwt.io/introduction)를 참고해주세요.

서비스 자체 인증 방식(external)을 통해 발급받은 토큰을 사용하여 iCL 2.0 서비스를 사용하는 방법은 각 SDK 문서를 참고해주시길 바랍니다.