# Unit 4: Admin App - Test Plan

## Test Cases

### 1. SSE Infrastructure
| ID | 테스트 | 기대 결과 |
|----|--------|----------|
| SSE-01 | SSE 연결 성공 | isConnected = true |
| SSE-02 | order_created 이벤트 | newOrderIds에 추가 |
| SSE-03 | order_updated 이벤트 | Query 무효화 |
| SSE-04 | 연결 끊김 시 재연결 | 3초 후 재연결 시도 |

### 2. Authentication
| ID | 테스트 | 기대 결과 |
|----|--------|----------|
| AUTH-01 | 로그인 성공 | /dashboard 리다이렉트 |
| AUTH-02 | 로그인 실패 | 에러 메시지 표시 |
| AUTH-03 | 미인증 접근 | /login 리다이렉트 |
| AUTH-04 | 로그아웃 | auth 클리어, /login 이동 |

### 3. Order Monitoring
| ID | 테스트 | 기대 결과 |
|----|--------|----------|
| ORD-01 | 주문 목록 표시 | 테이블별 그룹화 |
| ORD-02 | 신규 주문 강조 | isNew=true 시 강조 스타일 |
| ORD-03 | 주문 상세 Modal | 메뉴 목록 표시 |
| ORD-04 | 상태 변경 | API 호출 + UI 업데이트 |
| ORD-05 | 주문 삭제 | 확인 후 삭제 |

### 4. Table Management
| ID | 테스트 | 기대 결과 |
|----|--------|----------|
| TBL-01 | 테이블 목록 표시 | 번호, 상태, 총액 |
| TBL-02 | 테이블 설정 | 번호, 비밀번호 저장 |
| TBL-03 | 이용 완료 | 확인 Modal → 세션 종료 |
| TBL-04 | 과거 내역 조회 | 날짜 필터, 주문 목록 |

### 5. Menu Management
| ID | 테스트 | 기대 결과 |
|----|--------|----------|
| MNU-01 | 메뉴 목록 표시 | 카테고리별 그룹화 |
| MNU-02 | 메뉴 등록 | 폼 제출 → 목록 추가 |
| MNU-03 | 메뉴 수정 | 기존 값 로드 → 저장 |
| MNU-04 | 메뉴 삭제 | 확인 후 삭제 |
| MNU-05 | 순서 변경 | displayOrder 업데이트 |

### 6. i18n
| ID | 테스트 | 기대 결과 |
|----|--------|----------|
| I18N-01 | 기본 언어 (ko) | 한국어 텍스트 |
| I18N-02 | 언어 변경 (en) | 영어 텍스트 |

---

## Total: 21 Test Cases

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
