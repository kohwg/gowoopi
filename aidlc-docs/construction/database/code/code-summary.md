# Code Summary - Unit 1: Database

## 생성된 파일

### Model Layer
| 파일 | 설명 |
|------|------|
| `backend/internal/model/base.go` | BaseModel (CreatedAt, UpdatedAt, DeletedAt) |
| `backend/internal/model/store.go` | Store 엔티티 + UUID BeforeCreate 훅 |
| `backend/internal/model/category.go` | Category 엔티티 |
| `backend/internal/model/table.go` | Table 엔티티 |
| `backend/internal/model/session.go` | TableSession 엔티티 + UUID BeforeCreate 훅 |
| `backend/internal/model/menu.go` | Menu 엔티티 + MenuOrderInput |
| `backend/internal/model/order.go` | Order, OrderItem, OrderStatus (IsValid, CanTransitionTo) |
| `backend/internal/model/order_history.go` | OrderHistory 엔티티 |

### Repository Layer
| 파일 | 설명 |
|------|------|
| `backend/internal/repository/repository.go` | 6개 Repository 인터페이스 정의 |
| `backend/internal/repository/impl/store.go` | StoreRepository 구현 |
| `backend/internal/repository/impl/category.go` | CategoryRepository 구현 |
| `backend/internal/repository/impl/table.go` | TableRepository 구현 |
| `backend/internal/repository/impl/session.go` | SessionRepository 구현 |
| `backend/internal/repository/impl/menu.go` | MenuRepository 구현 (UpdateOrder 트랜잭션) |
| `backend/internal/repository/impl/order.go` | OrderRepository 구현 (Create, MoveToHistory 트랜잭션) |

### Database Layer
| 파일 | 설명 |
|------|------|
| `backend/internal/database/database.go` | Connect (Pool 설정) + AutoMigrate |
| `backend/internal/database/seed.go` | 개발용 시드 데이터 |

### Test Files
| 파일 | 테스트 수 |
|------|-----------|
| `backend/internal/model/order_test.go` | 2 (IsValid, CanTransitionTo) |
| `backend/internal/model/store_test.go` | 2 (UUID 생성, 기존 ID 보존) |
| `backend/internal/repository/impl/store_test.go` | 3 |
| `backend/internal/repository/impl/category_test.go` | 3 |
| `backend/internal/repository/impl/table_test.go` | 2 |
| `backend/internal/repository/impl/session_test.go` | 3 |
| `backend/internal/repository/impl/menu_test.go` | 4 |
| `backend/internal/repository/impl/order_test.go` | 7 |

### Utility
| 파일 | 설명 |
|------|------|
| `backend/internal/testutil/testutil.go` | 테스트 DB 설정/정리 헬퍼 |
| `backend/go.mod` | Go module 정의 |

---

**Total Files**: 21  
**Total Test Cases**: 25 (Unit: 4, Integration: 21)
