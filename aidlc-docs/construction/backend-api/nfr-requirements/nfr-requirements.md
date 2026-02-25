# NFR Requirements - Unit 2: Backend API

## 성능
- API 응답: 100ms 이하
- 동시접속: 500명+
- SSE 연결: 매장당 다수 관리자 동시 연결

## 보안
- JWT: Access 15분 + Refresh 30일, HTTP-only Secure Cookie
- bcrypt: cost=10
- CORS: localhost:3000, localhost:3001
- 입력 검증: Gin binding 태그
- SQL Injection: GORM 파라미터 바인딩

## 가용성
- Graceful shutdown
- SSE 연결 해제 시 자동 정리
- 구조화된 로깅 (요청/응답/에러)

## 기술 스택
| 항목 | 선택 | 근거 |
|------|------|------|
| HTTP Framework | Gin | 고성능, Go 표준 |
| JWT | golang-jwt/jwt/v5 | 표준 JWT 라이브러리 |
| Password | bcrypt | 단방향 해시 |
| Logging | log/slog | Go 표준 구조화 로깅 |
| Config | 환경변수 | 단순, On-premises 적합 |
