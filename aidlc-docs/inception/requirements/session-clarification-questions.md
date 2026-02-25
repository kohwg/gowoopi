# Session Management Clarification Questions

세션 관리 방식에 대한 명확화 질문입니다.

---

## Question 1: 관리자 세션 관리 방식
관리자의 16시간 세션 유지 방식을 명확히 해주세요:

A) 16시간 후 무조건 로그아웃 (매일 새로 로그인 필요)
B) 16시간 Access Token + Refresh Token (자동 갱신으로 계속 로그인 유지)
C) 16시간 동안 활동이 있으면 세션 연장 (Sliding Expiration)
D) 브라우저 닫기 전까지 유지 (세션 기반)
E) Other (please describe after [Answer]: tag below)

[Answer]: B

---

## Question 2: 테이블 태블릿 세션 관리 방식
테이블 태블릿의 16시간 세션은 어떻게 관리하시겠습니까?

A) 16시간 후 무조건 재로그인 필요
B) 자동 갱신으로 계속 로그인 유지 (태블릿은 항상 로그인 상태)
C) 매장 이용 완료 시에만 로그아웃, 그 외에는 계속 유지
D) 브라우저 닫기 전까지 유지
E) Other (please describe after [Answer]: tag below)

[Answer]: B

---

## Question 3: Refresh Token 구현 여부
Refresh Token을 구현하시겠습니까?

A) Yes - Access Token (짧은 만료) + Refresh Token (긴 만료) 구조
B) No - Access Token만 사용 (16시간 만료)
C) 관리자만 Refresh Token 사용, 테이블은 Access Token만
D) 테이블만 Refresh Token 사용, 관리자는 Access Token만
E) Other (please describe after [Answer]: tag below)

[Answer]: A

---

**모든 질문에 답변하신 후 "완료" 또는 "done"이라고 말씀해주세요.**
