# Component Methods - 테이블오더 서비스

## Handler Layer 메서드

### handler/auth.go
| 메서드 | HTTP | 경로 | 입력 | 출력 | 목적 |
|--------|------|------|------|------|------|
| CustomerLogin | POST | /api/customer/login | StoreID, TableNumber, Password | JWT Tokens (Cookie) | 테이블 태블릿 로그인 |
| AdminLogin | POST | /api/admin/login | StoreID, Username, Password | JWT Tokens (Cookie) | 관리자 로그인 |
| RefreshToken | POST | /api/auth/refresh | Refresh Token (Cookie) | New Access Token (Cookie) | 토큰 갱신 |

### handler/menu.go
| 메서드 | HTTP | 경로 | 입력 | 출력 | 목적 |
|--------|------|------|------|------|------|
| GetMenus | GET | /api/customer/menus | StoreID (context) | []Menu | 메뉴 목록 조회 |
| CreateMenu | POST | /api/admin/menus | MenuCreateRequest | Menu | 메뉴 등록 |
| UpdateMenu | PUT | /api/admin/menus/:id | MenuUpdateRequest | Menu | 메뉴 수정 |
| DeleteMenu | DELETE | /api/admin/menus/:id | MenuID | - | 메뉴 삭제 |
| UpdateMenuOrder | PATCH | /api/admin/menus/order | []MenuOrderRequest | - | 메뉴 순서 변경 |

### handler/order.go
| 메서드 | HTTP | 경로 | 입력 | 출력 | 목적 |
|--------|------|------|------|------|------|
| CreateOrder | POST | /api/customer/orders | OrderCreateRequest | Order | 주문 생성 |
| GetCustomerOrders | GET | /api/customer/orders | SessionID (context) | []Order | 고객 주문 내역 |
| GetAdminOrders | GET | /api/admin/orders | StoreID (context) | []Order | 관리자 주문 목록 |
| UpdateOrderStatus | PATCH | /api/admin/orders/:id/status | StatusUpdateRequest | Order | 주문 상태 변경 |
| DeleteOrder | DELETE | /api/admin/orders/:id | OrderID | - | 주문 삭제 |

### handler/table.go
| 메서드 | HTTP | 경로 | 입력 | 출력 | 목적 |
|--------|------|------|------|------|------|
| SetupTable | POST | /api/admin/tables/setup | TableSetupRequest | Table | 테이블 초기 설정 |
| CompleteTable | POST | /api/admin/tables/:id/complete | TableID | - | 테이블 이용 완료 |
| GetTableHistory | GET | /api/admin/tables/:id/history | TableID, DateFilter | []OrderHistory | 과거 내역 조회 |

### handler/sse.go
| 메서드 | HTTP | 경로 | 입력 | 출력 | 목적 |
|--------|------|------|------|------|------|
| StreamOrders | GET | /api/admin/orders/stream | StoreID (context) | SSE Stream | 실시간 주문 스트림 |

---

## Service Layer 메서드

### service/auth.go
| 메서드 | 입력 | 출력 | 목적 |
|--------|------|------|------|
| CustomerLogin | StoreID, TableNumber, Password | TokenPair, error | 테이블 인증 |
| AdminLogin | StoreID, Username, Password | TokenPair, error | 관리자 인증 |
| RefreshToken | RefreshToken string | AccessToken, error | 토큰 갱신 |
| GenerateTokenPair | Claims | TokenPair, error | JWT 토큰 쌍 생성 |
| ValidateToken | Token string | Claims, error | 토큰 검증 |
| HashPassword | Password string | Hash string, error | bcrypt 해싱 |
| ComparePassword | Hash, Password | error | bcrypt 비교 |

### service/menu.go
| 메서드 | 입력 | 출력 | 목적 |
|--------|------|------|------|
| GetMenusByStore | StoreID | []Menu, error | 매장별 메뉴 조회 |
| CreateMenu | MenuCreateInput | Menu, error | 메뉴 등록 |
| UpdateMenu | MenuID, MenuUpdateInput | Menu, error | 메뉴 수정 |
| DeleteMenu | MenuID | error | 메뉴 삭제 |
| UpdateMenuOrder | []MenuOrderInput | error | 메뉴 순서 변경 |

### service/order.go
| 메서드 | 입력 | 출력 | 목적 |
|--------|------|------|------|
| CreateOrder | OrderCreateInput | Order, error | 주문 생성 + SSE 이벤트 발행 |
| GetOrdersBySession | SessionID | []Order, error | 세션별 주문 조회 |
| GetOrdersByStore | StoreID | []Order, error | 매장별 주문 조회 |
| UpdateOrderStatus | OrderID, Status | Order, error | 주문 상태 변경 + SSE 이벤트 |
| DeleteOrder | OrderID | error | 주문 삭제 + 총액 재계산 + SSE 이벤트 |

### service/table.go
| 메서드 | 입력 | 출력 | 목적 |
|--------|------|------|------|
| SetupTable | TableSetupInput | Table, error | 테이블 초기 설정 |
| CompleteTable | TableID | error | 이용 완료 (주문→이력 이동, 리셋) |
| GetTableHistory | TableID, DateFilter | []OrderHistory, error | 과거 내역 조회 |
| StartSession | TableID | TableSession, error | 새 세션 시작 |

### service/sse.go
| 메서드 | 입력 | 출력 | 목적 |
|--------|------|------|------|
| Subscribe | StoreID, ClientChan | - | SSE 클라이언트 등록 |
| Unsubscribe | StoreID, ClientChan | - | SSE 클라이언트 해제 |
| Broadcast | StoreID, Event | - | 매장별 이벤트 브로드캐스트 |

---

## Repository Layer 메서드

### repository/store.go
| 메서드 | 입력 | 출력 | 목적 |
|--------|------|------|------|
| FindByID | StoreID | Store, error | 매장 조회 |
| FindByIDAndUsername | StoreID, Username | Store, error | 관리자 인증용 조회 |

### repository/table.go
| 메서드 | 입력 | 출력 | 목적 |
|--------|------|------|------|
| FindByStoreAndNumber | StoreID, TableNumber | Table, error | 테이블 조회 |
| Create | Table | Table, error | 테이블 생성 |
| Update | Table | error | 테이블 업데이트 |
| CreateSession | TableSession | TableSession, error | 세션 생성 |
| EndSession | SessionID | error | 세션 종료 |
| FindActiveSession | TableID | TableSession, error | 활성 세션 조회 |

### repository/menu.go
| 메서드 | 입력 | 출력 | 목적 |
|--------|------|------|------|
| FindByStore | StoreID | []Menu, error | 매장별 메뉴 조회 |
| FindByID | MenuID | Menu, error | 메뉴 단건 조회 |
| Create | Menu | Menu, error | 메뉴 생성 |
| Update | Menu | error | 메뉴 수정 |
| Delete | MenuID | error | 메뉴 삭제 |
| UpdateOrder | []MenuOrderInput | error | 순서 변경 |

### repository/order.go
| 메서드 | 입력 | 출력 | 목적 |
|--------|------|------|------|
| Create | Order, []OrderItem | Order, error | 주문+항목 생성 (트랜잭션) |
| FindBySession | SessionID | []Order, error | 세션별 주문 조회 |
| FindByStore | StoreID | []Order, error | 매장별 활성 주문 조회 |
| FindByID | OrderID | Order, error | 주문 단건 조회 |
| UpdateStatus | OrderID, Status | error | 상태 변경 |
| Delete | OrderID | error | 주문 삭제 |
| MoveToHistory | SessionID | error | 주문→이력 이동 (트랜잭션) |
| FindHistory | TableID, DateFilter | []OrderHistory, error | 과거 내역 조회 |

---

**참고**: 상세 비즈니스 규칙은 Functional Design (CONSTRUCTION) 단계에서 정의됩니다.

**문서 버전**: 1.0  
**작성일**: 2026-02-25
