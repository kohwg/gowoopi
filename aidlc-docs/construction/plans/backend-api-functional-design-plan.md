# Functional Design Plan - Unit 2: Backend API

## 범위
5개 서비스 (Auth, Menu, Order, Table, SSE) + 14개 API 엔드포인트 + Middleware

---

## 질문

### Question 1: JWT 토큰 만료 시간
Access Token과 Refresh Token의 만료 시간을 어떻게 설정하시겠습니까?

A) Access: 15분, Refresh: 16시간 (테이블 영업시간 기준)
B) Access: 30분, Refresh: 24시간
C) Access: 1시간, Refresh: 7일
D) Other

[Answer]: D 명시적 로그아웃 없으면 무한유지

### Question 2: 관리자 인증 방식
관리자 로그인 시 Store 테이블의 admin_username/password를 사용합니다. 관리자 계정은 매장당 1개로 충분합니까?

A) 매장당 1개 관리자 계정 (현재 설계대로)
B) 매장당 다수 관리자 계정 (별도 Admin 테이블 필요)

[Answer]: B

### Question 3: SSE 하트비트 간격
SSE 연결 유지를 위한 하트비트 전송 간격은?

A) 15초
B) 30초
C) 60초

[Answer]: A

### Question 4: API 에러 응답 형식
API 에러 응답을 어떤 형식으로 통일하시겠습니까?

A) `{"error": {"code": "NOT_FOUND", "message": "주문을 찾을 수 없습니다"}}` (구조화)
B) `{"error": "주문을 찾을 수 없습니다"}` (단순)

[Answer]: A

### Question 5: CORS 설정
프론트엔드 개발 시 CORS 허용 origin은?

A) localhost:3000, localhost:3001 (개발용 고정)
B) 환경변수로 설정 가능 (유연)

[Answer]: A

---

**답변 후 "완료"라고 말씀해주세요.**

---

## 생성 실행 계획
- [x] business-logic-model.md - 서비스별 비즈니스 로직 상세
- [x] business-rules.md - API 비즈니스 규칙 및 검증
- [x] domain-entities.md - DTO/Request/Response 타입 정의

---

**Plan Status**: 완료
