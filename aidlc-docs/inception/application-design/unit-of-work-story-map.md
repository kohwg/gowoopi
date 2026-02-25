# Unit of Work Story Map - 테이블오더 서비스

## Story → Unit 매핑

### Unit 1: Database
| Story ID | 스토리 | 관련 테이블 |
|----------|--------|-------------|
| 전체 | 모든 데이터 관련 스토리의 기반 | Store, Table, TableSession, Menu, Order, OrderItem, OrderHistory |

### Unit 2: Backend API
| Story ID | 스토리 | 관련 서비스 |
|----------|--------|-------------|
| US-001 | 테이블 태블릿 자동 로그인 | Auth Service |
| US-006 | 매장 인증 | Auth Service |
| US-002 | 메뉴 조회 및 탐색 | Menu Service |
| US-010 | 메뉴 관리 | Menu Service |
| US-003 | 장바구니 관리 | - (클라이언트 전용) |
| US-004 | 주문 생성 | Order Service |
| US-005 | 주문 내역 조회 | Order Service |
| US-007 | 실시간 주문 모니터링 | Order Service, SSE Manager |
| US-008 | 주문 상태 관리 | Order Service |
| US-009 | 테이블 관리 | Table Service, Order Service |
| US-011 | API 성능 | 전체 |
| US-012 | 실시간 통신 성능 | SSE Manager |
| US-013 | 보안 | Auth Service, Middleware |
| US-014 | 데이터 무결성 | Repository Layer |
| US-015 | 로깅 및 모니터링 | Middleware |

### Unit 3: Customer App
| Story ID | 스토리 | 관련 페이지/컴포넌트 |
|----------|--------|---------------------|
| US-001 | 테이블 태블릿 자동 로그인 | 자동 로그인 로직 |
| US-002 | 메뉴 조회 및 탐색 | 메뉴 페이지 (기본 화면) |
| US-003 | 장바구니 관리 | 장바구니 (Zustand store) |
| US-004 | 주문 생성 | 주문 확인/생성 페이지 |
| US-005 | 주문 내역 조회 | 주문 내역 페이지 |
| US-017 | 다국어 지원 | i18n 설정 |

### Unit 4: Admin App
| Story ID | 스토리 | 관련 페이지/컴포넌트 |
|----------|--------|---------------------|
| US-006 | 매장 인증 | 로그인 페이지 |
| US-007 | 실시간 주문 모니터링 | 대시보드 (SSE 연동) |
| US-008 | 주문 상태 관리 | 대시보드 상태 변경 |
| US-009 | 테이블 관리 | 테이블 관리 페이지 |
| US-010 | 메뉴 관리 | 메뉴 관리 페이지 |
| US-017 | 다국어 지원 | i18n 설정 |

---

## 커버리지 확인
- ✅ 모든 Feature Story (US-001 ~ US-010) 매핑 완료
- ✅ 모든 NFR Story (US-011 ~ US-017) 매핑 완료
- ✅ US-016 (테스트 커버리지)는 모든 Unit에 적용
- ✅ US-003 (장바구니)는 클라이언트 전용 (Unit 3)

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
