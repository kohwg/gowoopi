# NFR Design Patterns - Unit 2: Backend API

## Middleware 체인
```
Request → Recovery → Logger → CORS → [Auth] → Handler → Response
```

- **Recovery**: panic 복구, 500 응답
- **Logger**: slog 구조화 로깅 (method, path, status, latency)
- **CORS**: Gin CORS middleware
- **Auth**: JWT 검증, Claims를 context에 저장

## 에러 처리 패턴
- 통일된 ErrorResponse 구조 (`{error: {code, message}}`)
- 서비스 레이어에서 커스텀 에러 반환
- 핸들러에서 에러 코드 → HTTP 상태 매핑

## SSE 패턴
- 매장별 클라이언트 맵: `map[string][]chan SSEEvent`
- sync.RWMutex로 동시성 제어
- 하트비트: 15초 ticker
- Context 취소 시 자동 정리

## 환경변수
| 변수 | 기본값 | 설명 |
|------|--------|------|
| SERVER_PORT | 8080 | 서버 포트 |
| JWT_SECRET | - | JWT 서명 키 (필수) |
| CORS_ORIGINS | http://localhost:3000,http://localhost:3001 | CORS |
| DB_HOST | localhost | MySQL |
| DB_PORT | 3306 | MySQL |
| DB_USER | root | MySQL |
| DB_PASSWORD | - | MySQL |
| DB_NAME | table_order | MySQL |
| DB_SEED | false | 시드 데이터 |
