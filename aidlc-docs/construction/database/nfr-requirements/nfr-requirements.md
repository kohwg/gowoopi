# NFR Requirements - Unit 1: Database

## 성능
- API 응답 100ms 이하 (DB 쿼리 포함)
- 500명+ 동시접속 지원
- Connection Pool: MaxOpenConns=25, MaxIdleConns=10, ConnMaxLifetime=5m

## 보안
- 비밀번호: bcrypt 해시 저장 (cost=10)
- UUID PK: 외부 노출 엔티티 (Store, TableSession, Order)
- SQL Injection 방지: GORM 파라미터 바인딩
- DB 접속 정보: 환경변수로 관리

## 데이터 무결성
- Soft Delete: Store, Category, Table, Menu, Order
- 외래키 제약조건: 참조 무결성 보장
- 트랜잭션: 주문 생성, 이용 완료, 주문 삭제

## 가용성
- On-premises 단일 인스턴스 (MVP)
- GORM AutoMigrate: 서버 시작 시 자동 스키마 동기화

## 다국어
- utf8mb4 문자셋: 한국어, 영어, 이모지 완벽 지원
- utf8mb4_unicode_ci 정렬: 정확한 다국어 정렬

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
