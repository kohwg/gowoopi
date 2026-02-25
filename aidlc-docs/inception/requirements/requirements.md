# 테이블오더 서비스 요구사항 문서

## Intent Analysis Summary

### User Request
테이블오더 서비스 구축 - 고객용 주문 인터페이스, 관리자용 관리 인터페이스, 서버 시스템, 데이터 저장소 포함

### Request Type
New Project (신규 프로젝트)

### Scope Estimate
System-wide - 다중 인터페이스, 실시간 통신, 세션 관리, 데이터베이스 포함

### Complexity Estimate
Complex - 실시간 주문 모니터링(SSE), 다중 사용자 세션 관리, 고성능 요구사항

---

## 1. 기술 스택

### 1.1 프론트엔드
- **프레임워크**: React
- **UI 라이브러리**: Tailwind CSS
- **상태 관리**: React Context API 또는 Redux (필요시)

### 1.2 백엔드
- **언어/프레임워크**: Go (Gin 또는 Echo 프레임워크)
- **인증**: JWT (HTTP-only Cookie 저장)
- **실시간 통신**: Server-Sent Events (SSE)

### 1.3 데이터베이스
- **타입**: Relational Database (PostgreSQL 또는 MySQL)
- **ORM**: GORM (Go) 또는 직접 SQL

### 1.4 배포 환경
- **환경**: 로컬 서버 (On-premises)
- **컨테이너화**: Docker (선택사항)

---

## 2. 기능 요구사항

### 2.1 고객용 기능 (Customer Interface)

#### 2.1.1 테이블 태블릿 자동 로그인 및 세션 관리
**목적**: 고객이 별도 로그인 절차 없이 즉시 주문

**기능**:
- 초기 설정 (관리자 1회 수행):
  - 매장 식별자 입력 (UUID 형식)
  - 테이블 번호 입력
  - 테이블 비밀번호 입력
  - 로그인 정보 Cookie에 저장
- 자동 로그인: 저장된 정보로 자동 인증
- 세션 유지: 16시간

#### 2.1.2 메뉴 조회 및 탐색
**기능**:
- 메뉴 화면이 기본 화면
- 카테고리별 메뉴 분류
- 메뉴 상세 정보: 메뉴명, 가격, 설명, 이미지 (외부 URL)
- 카드 형태 레이아웃
- 터치 친화적 UI (최소 44x44px 버튼)

#### 2.1.3 장바구니 관리
**기능**:
- 메뉴 추가/삭제
- 수량 조절
- 총 금액 실시간 계산
- 장바구니 비우기
- 로컬 저장 (페이지 새로고침 시 유지)

#### 2.1.4 주문 생성
**기능**:
- 주문 내역 최종 확인
- 주문 확정
- 주문 성공 시:
  - 주문 번호 표시
  - 장바구니 자동 비우기
  - 메뉴 화면으로 자동 리다이렉트
- 주문 실패 시: 에러 메시지 표시, 장바구니 유지

**주문 정보**:
- 매장 식별자 (UUID)
- 테이블 식별자
- 주문 메뉴 목록 (메뉴명, 수량, 단가)
- 총 주문 금액
- 세션 ID

#### 2.1.5 주문 내역 조회
**기능**:
- 주문 시간 순 정렬
- 주문별 상세 정보: 주문 번호, 시각, 메뉴/수량, 금액, 상태
- 현재 테이블 세션 주문만 표시
- 주문 상태 실시간 업데이트 (선택사항)

### 2.2 관리자용 기능 (Admin Interface)

#### 2.2.1 매장 인증
**기능**:
- 매장 식별자 입력 (UUID)
- 사용자명 및 비밀번호 입력
- JWT 토큰 기반 인증 (HTTP-only Cookie)
- 16시간 세션 유지
- 브라우저 새로고침 시 세션 유지
- 비밀번호 bcrypt 해싱

#### 2.2.2 실시간 주문 모니터링
**기능**:
- Server-Sent Events (SSE) 기반 실시간 업데이트
- 그리드/대시보드 레이아웃:
  - 테이블별 카드 형태
  - 각 테이블 총 주문액 표시
  - 최신 주문 n개 미리보기
- 주문 카드 클릭 시 전체 메뉴 목록 상세 보기
- 주문 상태 변경 (대기중/준비중/완료)
- 신규 주문 시각적 강조
- 2초 이내 주문 표시
- 테이블별 필터링

#### 2.2.3 테이블 관리
**기능**:
1. **테이블 태블릿 초기 설정**:
   - 테이블 번호 및 비밀번호 설정
   - 16시간 세션 생성
2. **주문 삭제**:
   - 특정 주문 삭제
   - 확인 팝업
   - 테이블 총 주문액 재계산
3. **테이블 세션 처리**:
   - 세션 시작/종료 관리
   - 이용 완료 시 주문 내역 과거 이력으로 이동
   - 테이블 현재 주문 목록 및 총 주문액 리셋
4. **과거 주문 내역 조회**:
   - 테이블별 과거 주문 목록 (시간 역순)
   - 날짜 필터링

#### 2.2.4 메뉴 관리
**기능**:
- 메뉴 조회 (카테고리별)
- 메뉴 등록: 메뉴명, 가격, 설명, 카테고리, 이미지 URL
- 메뉴 수정
- 메뉴 삭제
- 메뉴 노출 순서 조정
- 필수 필드 검증

---

## 3. 비기능 요구사항

### 3.1 성능
- **API 응답 시간**: 100ms 이하 목표
- **동시 접속**: 대규모 (50개 이상 테이블, 500명 이상) 지원
- **확장성**: 더 큰 규모로 확장 가능한 아키텍처
- **실시간 업데이트**: 2초 이내 주문 표시

### 3.2 보안
- **인증**: JWT 기반, HTTP-only Cookie 저장
- **비밀번호**: bcrypt 해싱
- **세션**: 16시간 자동 만료
- **로그인 시도 제한**: 구현 권장

### 3.3 데이터 관리
- **장바구니**: 클라이언트 측 로컬 저장 (Cookie)
- **주문 이력**: 데이터베이스 영구 저장
- **세션 추적**: 세션 ID로 주문 그룹화

### 3.4 로깅 및 모니터링
- **로깅 수준**: 구조화된 로깅 (파일 또는 로깅 서비스)
- **에러 처리**: 명확한 에러 메시지 및 로깅

### 3.5 테스트
- **테스트 커버리지**: Unit + Integration + E2E Test
- **핵심 비즈니스 로직**: Unit Test 필수
- **API 엔드포인트**: Integration Test
- **사용자 플로우**: E2E Test

### 3.6 다국어
- **MVP 지원**: 한국어 + 영어
- **향후 확장**: 추가 언어 지원 가능하도록 설계

### 3.7 이미지 관리
- **MVP**: 외부 URL만 지원
- **향후 확장**: 파일 업로드 기능 추가 가능하도록 설계

---

## 4. 데이터 모델 (개념적)

### 4.1 주요 엔티티

#### Store (매장)
- store_id (UUID, PK)
- store_name
- admin_username
- admin_password_hash
- created_at

#### Table (테이블)
- table_id (PK)
- store_id (FK)
- table_number
- table_password_hash
- current_session_id
- created_at

#### TableSession (테이블 세션)
- session_id (UUID, PK)
- table_id (FK)
- start_time
- end_time (nullable)
- is_active

#### Menu (메뉴)
- menu_id (PK)
- store_id (FK)
- menu_name
- price
- description
- category
- image_url
- display_order
- created_at

#### Order (주문)
- order_id (UUID, PK)
- session_id (FK)
- table_id (FK)
- store_id (FK)
- total_amount
- status (대기중/준비중/완료)
- created_at

#### OrderItem (주문 항목)
- order_item_id (PK)
- order_id (FK)
- menu_id (FK)
- menu_name (스냅샷)
- quantity
- unit_price (스냅샷)

#### OrderHistory (주문 이력)
- 완료된 세션의 주문 데이터
- Order 테이블과 동일 구조 + completed_at

---

## 5. API 엔드포인트 (개념적)

### 5.1 고객용 API
- `POST /api/customer/login` - 테이블 로그인
- `GET /api/customer/menus` - 메뉴 목록 조회
- `POST /api/customer/orders` - 주문 생성
- `GET /api/customer/orders` - 주문 내역 조회

### 5.2 관리자용 API
- `POST /api/admin/login` - 관리자 로그인
- `GET /api/admin/orders/stream` - SSE 실시간 주문 스트림
- `GET /api/admin/orders` - 주문 목록 조회
- `PATCH /api/admin/orders/:id/status` - 주문 상태 변경
- `DELETE /api/admin/orders/:id` - 주문 삭제
- `POST /api/admin/tables/:id/complete` - 테이블 이용 완료
- `GET /api/admin/tables/:id/history` - 과거 주문 내역
- `GET /api/admin/menus` - 메뉴 목록 조회
- `POST /api/admin/menus` - 메뉴 등록
- `PUT /api/admin/menus/:id` - 메뉴 수정
- `DELETE /api/admin/menus/:id` - 메뉴 삭제

---

## 6. 제외 기능 (constraints.md 기반)

다음 기능은 MVP에서 구현하지 않습니다:
- 실제 결제 처리
- 복잡한 사용자 인증 (OAuth, SNS 로그인)
- 이미지 업로드/리사이징
- 푸시 알림, SMS, 이메일
- 주방 기능 (주방 디스플레이, 재고 관리)
- 데이터 분석 대시보드
- 매출 리포트
- 직원 관리
- 예약 시스템
- 고객 리뷰
- 외부 연동 (배달 플랫폼, POS)

---

## 7. 주요 요구사항 요약

### 핵심 기능
1. **고객**: 자동 로그인, 메뉴 조회, 장바구니, 주문 생성, 주문 내역
2. **관리자**: 인증, 실시간 주문 모니터링(SSE), 테이블 관리, 메뉴 관리

### 핵심 기술
- React + Tailwind CSS (프론트엔드)
- Go + Gin/Echo (백엔드)
- PostgreSQL/MySQL (데이터베이스)
- JWT + HTTP-only Cookie (인증)
- Server-Sent Events (실시간 통신)

### 핵심 품질 속성
- 고성능 (API 응답 100ms 이하)
- 확장성 (500명 이상 동시 접속)
- 보안 (JWT, bcrypt, 세션 관리)
- 테스트 (Unit + Integration + E2E)
- 다국어 (한국어 + 영어)

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25  
**상태**: 승인 대기
