# Application Design Plan

## 설계 범위
테이블오더 서비스의 컴포넌트 식별, 서비스 레이어 설계, 의존성 정의

---

## 설계 질문

### Question 1: 프론트엔드 프로젝트 구조
고객용 UI와 관리자용 UI를 어떻게 구성하시겠습니까?

A) 단일 React 프로젝트 (라우팅으로 분리: /customer/*, /admin/*)
B) 별도 React 프로젝트 2개 (customer-app, admin-app)
C) 모노레포 (하나의 저장소에 두 프로젝트, 공통 컴포넌트 공유)
D) Other (please describe after [Answer]: tag below)

[Answer]: C

---

### Question 2: 백엔드 API 구조
백엔드 API를 어떻게 구성하시겠습니까?

A) 단일 Go 서버 (라우트 그룹으로 분리: /api/customer/*, /api/admin/*)
B) 별도 Go 서버 2개 (customer-api, admin-api)
C) 단일 Go 서버 + 도메인별 패키지 분리 (handler, service, repository 레이어)
D) Other (please describe after [Answer]: tag below)

[Answer]: C

---

### Question 3: 데이터베이스 접근 패턴
데이터베이스 접근 레이어를 어떻게 설계하시겠습니까?

A) Repository 패턴 (각 엔티티별 Repository 인터페이스)
B) DAO 패턴 (Data Access Object)
C) GORM 직접 사용 (서비스 레이어에서 직접 호출)
D) Other (please describe after [Answer]: tag below)

[Answer]: A

---

### Question 4: 상태 관리 (프론트엔드)
React 프론트엔드의 상태 관리를 어떻게 하시겠습니까?

A) React Context API만 사용
B) Zustand (경량 상태 관리)
C) Redux Toolkit
D) Jotai / Recoil (Atomic 상태 관리)
E) Other (please describe after [Answer]: tag below)

[Answer]: B

---

### Question 5: API 통신 라이브러리 (프론트엔드)
프론트엔드에서 API 통신에 어떤 라이브러리를 사용하시겠습니까?

A) Axios
B) Fetch API (내장)
C) TanStack Query (React Query) + Fetch/Axios
D) SWR + Fetch/Axios
E) Other (please describe after [Answer]: tag below)

[Answer]: C

---

### Question 6: 라우팅 (프론트엔드)
프론트엔드 라우팅에 어떤 라이브러리를 사용하시겠습니까?

A) React Router v6
B) TanStack Router
C) Next.js (SSR/SSG 포함)
D) Other (please describe after [Answer]: tag below)

[Answer]: C

---

**모든 질문에 답변하신 후 "완료" 또는 "done"이라고 말씀해주세요.**

---

## 설계 실행 계획

답변을 받은 후 다음 산출물을 생성합니다:

- [ ] components.md - 컴포넌트 정의 및 책임
- [ ] component-methods.md - 메서드 시그니처 및 목적
- [ ] services.md - 서비스 정의 및 오케스트레이션
- [ ] component-dependency.md - 의존성 관계 및 통신 패턴

---

**Plan Status**: 답변 대기 중
