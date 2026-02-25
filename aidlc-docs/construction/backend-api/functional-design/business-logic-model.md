# Business Logic Model - Unit 2: Backend API

## 서비스 플로우

### Auth Service

#### 고객 로그인
```
1. StoreID + TableNumber로 Table 조회
2. bcrypt.Compare(password, table.PasswordHash)
3. 활성 세션 조회 → 없으면 새 세션 생성
4. JWT Claims 생성 (role=customer, store_id, table_id, session_id)
5. Access Token (15분) + Refresh Token (30일) 발급
6. HTTP-only Cookie에 토큰 설정
```

#### 관리자 로그인
```
1. StoreID + Username으로 Admin 조회
2. bcrypt.Compare(password, admin.PasswordHash)
3. JWT Claims 생성 (role=admin, store_id, admin_id)
4. Access Token + Refresh Token 발급
5. HTTP-only Cookie에 토큰 설정
```

#### 토큰 갱신
```
1. Cookie에서 Refresh Token 추출
2. Refresh Token 검증
3. 새 Access Token 발급
4. Cookie 업데이트
```

### Order Service

#### 주문 생성
```
1. 요청 검증 (items 최소 1개, quantity ≥ 1)
2. 각 메뉴 조회 → is_available 확인
3. OrderItem 생성 (menu_name, price 스냅샷)
4. Order + OrderItem 트랜잭션 생성 (Repository)
5. SSE 이벤트 발행: ORDER_CREATED
6. 응답 반환
```

#### 주문 상태 변경
```
1. Order 조회
2. CanTransitionTo 검증
3. UpdateStatus (Repository)
4. SSE 이벤트 발행: ORDER_STATUS_CHANGED
```

#### 주문 삭제
```
1. Order 조회
2. COMPLETED 상태 삭제 불가 검증
3. Soft Delete (Repository)
4. SSE 이벤트 발행: ORDER_DELETED
```

### Table Service

#### 이용 완료
```
1. 활성 세션 조회
2. MoveToHistory (Repository, 트랜잭션)
3. 세션 종료 (Repository)
4. SSE 이벤트 발행: TABLE_RESET
```

### SSE Manager

#### 이벤트 구조
```go
type SSEEvent struct {
    Type string      `json:"type"` // ORDER_CREATED, ORDER_STATUS_CHANGED, ORDER_DELETED, TABLE_RESET
    Data interface{} `json:"data"`
}
```

#### 연결 관리
```
Subscribe: StoreID → client channel 등록
Unsubscribe: client channel 해제
Broadcast: StoreID의 모든 client에 이벤트 전송
Heartbeat: 15초 간격 ping
```

---

## Middleware 체인
```
Request → Logger → CORS → [Auth Middleware] → Handler → Response
                           ↓
                    JWT 검증 + Claims 추출
                    → context에 Claims 저장
```

---

## 환경변수 (추가)
| 변수 | 기본값 | 설명 |
|------|--------|------|
| JWT_SECRET | - | JWT 서명 키 (필수) |
| SERVER_PORT | 8080 | 서버 포트 |
| CORS_ORIGINS | http://localhost:3000,http://localhost:3001 | CORS 허용 origin |

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
