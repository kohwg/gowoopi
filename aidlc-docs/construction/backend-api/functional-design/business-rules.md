# Business Rules - Unit 2: Backend API

## 인증 규칙

### JWT 토큰
- Access Token: 15분 만료 (자동 갱신으로 사실상 무한 유지)
- Refresh Token: 30일 만료
- 저장: HTTP-only Secure Cookie
- 자동 갱신: Access Token 만료 시 Refresh Token으로 자동 재발급
- 명시적 로그아웃: Cookie 삭제

### 고객 로그인
- StoreID + TableNumber + Password로 인증
- 성공 시: 활성 세션 확인 → 없으면 새 세션 생성
- Claims: store_id, role="customer", table_id, session_id

### 관리자 로그인
- StoreID + Username + Password로 인증 (Admin 테이블)
- Claims: store_id, role="admin", admin_id

---

## API 검증 규칙

### Menu
- name: 필수, 1~100자
- price: 필수, 0 이상 정수
- category_id: 유효한 카테고리 참조 필수
- 삭제: Soft Delete (기존 주문의 스냅샷 영향 없음)

### Order
- items: 최소 1개 항목 필수
- quantity: 1 이상
- 메뉴 가용성 확인 (is_available=true)
- 주문 생성 시 메뉴명/가격 스냅샷 저장

### Order Status
- 상태 전이: PENDING→CONFIRMED→PREPARING→COMPLETED
- 역방향/건너뛰기 불가 (model.CanTransitionTo 사용)
- COMPLETED 상태에서 삭제 불가

### Table
- 매장 내 테이블 번호 유일
- 이용 완료: 활성 주문 → OrderHistory 이동, 세션 종료

---

## 권한 규칙

| 엔드포인트 | 고객 | 관리자 |
|-----------|------|--------|
| POST /api/customer/login | ✅ (인증 불필요) | - |
| GET /api/customer/menus | ✅ | - |
| POST /api/customer/orders | ✅ | - |
| GET /api/customer/orders | ✅ | - |
| POST /api/admin/login | - | ✅ (인증 불필요) |
| GET /api/admin/orders/stream | - | ✅ |
| GET /api/admin/orders | - | ✅ |
| PATCH /api/admin/orders/:id/status | - | ✅ |
| DELETE /api/admin/orders/:id | - | ✅ |
| POST /api/admin/tables/setup | - | ✅ |
| POST /api/admin/tables/:id/complete | - | ✅ |
| GET /api/admin/tables/:id/history | - | ✅ |
| POST/PUT/DELETE /api/admin/menus | - | ✅ |
| POST /api/auth/refresh | ✅ | ✅ |

---

## SSE 규칙
- 매장별 채널 분리 (StoreID 기반)
- 하트비트: 15초 간격
- 이벤트 타입: ORDER_CREATED, ORDER_STATUS_CHANGED, ORDER_DELETED, TABLE_RESET
- 연결 해제 시 자동 Unsubscribe

---

## 에러 코드
| 코드 | HTTP | 설명 |
|------|------|------|
| UNAUTHORIZED | 401 | 인증 실패 |
| FORBIDDEN | 403 | 권한 없음 |
| NOT_FOUND | 404 | 리소스 없음 |
| VALIDATION_ERROR | 400 | 입력 검증 실패 |
| INVALID_STATUS_TRANSITION | 400 | 잘못된 상태 전이 |
| INTERNAL_ERROR | 500 | 서버 내부 오류 |

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
