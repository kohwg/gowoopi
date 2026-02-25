# Requirements Verification Questions

요구사항을 명확히 하기 위한 질문입니다. 각 질문에 대해 [Answer]: 태그 뒤에 선택한 옵션의 문자를 입력해주세요.

---

## Question 1
프론트엔드 기술 스택으로 어떤 프레임워크를 사용하시겠습니까? 

A) React
B) Vue.js
C) Angular
D) Vanilla JavaScript (프레임워크 없음)
E) Other (please describe after [Answer]: tag below)

[Answer]: A

---

## Question 2
백엔드 기술 스택으로 어떤 언어/프레임워크를 사용하시겠습니까?

A) Node.js (Express/NestJS)
B) Python (FastAPI/Django)
C) Java (Spring Boot)
D) Go
E) Other (please describe after [Answer]: tag below)

[Answer]: Go

---

## Question 3
데이터베이스로 어떤 기술을 사용하시겠습니까?

A) Relational Database (PostgreSQL, MySQL)
B) NoSQL Document (MongoDB, DynamoDB)
C) NoSQL Key-Value (Redis, DynamoDB)
D) 혼합 (Relational + NoSQL)
E) Other (please describe after [Answer]: tag below)

[Answer]: A

---

## Question 4
배포 환경은 어디입니까?

A) AWS (Lambda, ECS, EC2 등)
B) 로컬 서버 (On-premises)
C) 다른 클라우드 (Azure, GCP)
D) 컨테이너 환경 (Docker, Kubernetes)
E) Other (please describe after [Answer]: tag below)

[Answer]: B

---

## Question 5
실시간 주문 모니터링을 위한 Server-Sent Events(SSE) 외에 다른 실시간 통신 방식을 고려하시겠습니까?

A) SSE만 사용 (요구사항대로)
B) WebSocket 추가 고려
C) Polling 방식 사용
D) 실시간 통신 불필요
E) Other (please describe after [Answer]: tag below)

[Answer]: A

---

## Question 6
메뉴 이미지는 어떻게 관리하시겠습니까?

A) 외부 URL 직접 입력 (요구사항대로)
B) 파일 업로드 및 저장소 관리 (S3 등)
C) 이미지 없이 텍스트만
D) 혼합 (URL + 업로드)
E) Other (please describe after [Answer]: tag below)

[Answer]: 일단 A로 하고 추가될 수도 있음

---

## Question 7
인증 방식으로 JWT를 사용한다고 명시되어 있습니다. JWT 저장 위치는 어디입니까?

A) LocalStorage
B) SessionStorage
C) HTTP-only Cookie
D) Memory (Redux/Vuex 등 상태 관리)
E) Other (please describe after [Answer]: tag below)

[Answer]: C

---

## Question 8
테이블 태블릿의 자동 로그인 정보는 어디에 저장하시겠습니까?

A) LocalStorage
B) SessionStorage
C) Cookie
D) IndexedDB
E) Other (please describe after [Answer]: tag below)

[Answer]: C

---

## Question 9
동시 접속 사용자 수 예상 규모는 어느 정도입니까?

A) 소규모 (10개 미만 테이블, 동시 접속 ~50명)
B) 중규모 (10-50개 테이블, 동시 접속 ~200명)
C) 대규모 (50개 이상 테이블, 동시 접속 500명 이상)
D) 확장 가능하도록 설계 (규모 미정)
E) Other (please describe after [Answer]: tag below)

[Answer]: C + D

---

## Question 10
API 응답 시간 목표는 어느 정도입니까?

A) 100ms 이하 (매우 빠름)
B) 500ms 이하 (빠름)
C) 1초 이하 (보통)
D) 특별한 요구사항 없음
E) Other (please describe after [Answer]: tag below)

[Answer]: A

---

## Question 11
에러 처리 및 로깅 수준은 어느 정도로 하시겠습니까?

A) 기본 에러 처리만 (콘솔 로그)
B) 구조화된 로깅 (파일 또는 로깅 서비스)
C) 모니터링 도구 연동 (CloudWatch, Datadog 등)
D) 최소한의 에러 처리
E) Other (please describe after [Answer]: tag below)

[Answer]: B

---

## Question 12
테스트 커버리지 목표는 어느 정도입니까?

A) Unit Test만 (핵심 비즈니스 로직)
B) Unit + Integration Test
C) Unit + Integration + E2E Test
D) 테스트 코드 불필요
E) Other (please describe after [Answer]: tag below)

[Answer]: C

---

## Question 13
UI/UX 디자인 시스템이나 컴포넌트 라이브러리를 사용하시겠습니까?

A) Material-UI / MUI
B) Ant Design
C) Tailwind CSS
D) 커스텀 디자인 (라이브러리 없음)
E) Other (please describe after [Answer]: tag below)

[Answer]: C

---

## Question 14
다국어 지원이 필요합니까? (constraints.md에서 제외되었지만 확인)

A) 한국어만
B) 한국어 + 영어
C) 다국어 지원 필요
D) 향후 확장 가능하도록 설계
E) Other (please describe after [Answer]: tag below)

[Answer]: C

---

## Question 15
매장 식별자는 어떤 형식입니까?

A) 숫자 ID (예: 1, 2, 3)
B) 문자열 코드 (예: "store001", "gangnam-branch")
C) UUID
D) 사용자 정의 형식
E) Other (please describe after [Answer]: tag below)

[Answer]: C 

---

**모든 질문에 답변하신 후 "완료" 또는 "done"이라고 말씀해주세요.**
