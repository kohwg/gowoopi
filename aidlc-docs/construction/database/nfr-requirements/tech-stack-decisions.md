# Tech Stack Decisions - Unit 1: Database

| 항목 | 선택 | 근거 |
|------|------|------|
| RDBMS | MySQL 8.0 | LTS, 안정적, On-premises 적합 |
| 문자셋 | utf8mb4 | 다국어 + 이모지 지원 |
| 정렬 | utf8mb4_unicode_ci | 정확한 다국어 정렬 |
| ORM | GORM | Go 표준 ORM, AutoMigrate 지원 |
| 마이그레이션 | GORM AutoMigrate | 개발 편의, 서버 시작 시 자동 실행 |
| ID 전략 | UUID(외부) / BIGINT(내부) | 보안 + 성능 균형 |
| 삭제 전략 | Soft Delete (gorm.DeletedAt) | 데이터 보존, 감사 추적 |
| Connection Pool | 25/10/5m | 500명+ 동시접속 대응 |

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
