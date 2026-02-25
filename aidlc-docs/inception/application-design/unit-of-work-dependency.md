# Unit of Work Dependencies - 테이블오더 서비스

## 의존성 매트릭스

| Unit | Database | Backend API | Customer App | Admin App |
|------|----------|-------------|--------------|-----------|
| **Database** | - | 의존됨 | - | - |
| **Backend API** | 의존 | - | 의존됨 | 의존됨 |
| **Customer App** | - | 의존 | - | - |
| **Admin App** | - | 의존 | - | - |

## 의존성 상세

### Unit 1: Database → 없음
- 다른 Unit에 의존하지 않음
- Backend API가 이 Unit에 의존

### Unit 2: Backend API → Database
- MySQL 스키마에 의존 (GORM 모델 매핑)
- Customer App, Admin App이 이 Unit에 의존

### Unit 3: Customer App → Backend API, Shared
- REST API 엔드포인트에 의존
- shared 패키지에 의존

### Unit 4: Admin App → Backend API, Shared
- REST API + SSE 엔드포인트에 의존
- shared 패키지에 의존

## 개발 순서 (병렬)

```
Phase 1 (병렬):
  ├── Unit 1: Database (스키마 + 마이그레이션)
  └── Unit 2: Backend API (모델 + API 구현)
      └── shared: API 계약 정의 (TypeScript 타입)

Phase 2 (병렬, Phase 1과 부분 병렬 가능):
  ├── Unit 3: Customer App (API 계약 기반 개발)
  └── Unit 4: Admin App (API 계약 기반 개발)

Phase 3:
  └── 통합 테스트 (전체 Unit 연동)
```

## 통합 포인트
- **API 계약**: Backend API의 요청/응답 타입을 shared 패키지에 TypeScript interface로 정의
- **인증**: JWT 토큰 (HTTP-only Cookie)을 통한 프론트엔드-백엔드 인증
- **실시간**: SSE 이벤트 형식을 API 계약에 포함

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
