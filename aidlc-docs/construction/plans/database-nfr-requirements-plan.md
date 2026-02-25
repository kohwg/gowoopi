# NFR Requirements Plan - Unit 1: Database

## 범위
Database Unit의 비기능 요구사항 및 기술 스택 결정

---

## 질문

### Question 1: MySQL 버전
어떤 MySQL 버전을 사용하시겠습니까?

A) MySQL 8.0 (LTS, 안정적)
B) MySQL 8.4 (최신 LTS)
C) MariaDB 10.x (MySQL 호환)

[Answer]: A

### Question 2: 문자셋/정렬
데이터베이스 문자셋을 어떻게 설정하시겠습니까?

A) utf8mb4 + utf8mb4_unicode_ci (다국어 완벽 지원, 이모지 포함)
B) utf8mb4 + utf8mb4_general_ci (약간 빠름, 정렬 정확도 낮음)

[Answer]: A

### Question 3: Connection Pool 크기
500명+ 동시접속 기준, DB 커넥션 풀 설정은?

A) 기본값 사용 (GORM 기본)
B) MaxOpenConns=25, MaxIdleConns=10, ConnMaxLifetime=5m (권장)
C) 직접 지정 (값을 [Answer]: 뒤에 기재)

[Answer]: B

---

**답변 후 "완료"라고 말씀해주세요.**

---

## 생성 실행 계획
- [x] nfr-requirements.md - NFR 요구사항 정의
- [x] tech-stack-decisions.md - 기술 스택 결정

---

**Plan Status**: 완료
