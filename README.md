## 빗썸 open api 활용 , 자동매매 프로그램(BackEnd)

### 기획

- **빗썸 api 정리 및 리팩토링**
    - [Info](https://github.com/myungsworld/myungsworld/tree/master/api/bithumb/Info) : 거래대기,사용가능한돈,캔들스틱,코인별잔고,코인별지갑주소,코인시세
    - [Transaction](https://github.com/myungsworld/myungsworld/tree/master/api/bithumb/transaction) : 매도예약,매수예약,시장가매도,시장가매수,예약가매수,원화출금
- **요구사항**
    - 코인마다 매수,매도 최소 수량이 다름 (api 쓸때 구분 필요) -> 소수점 리팩토링
    - 초단위의 모니터링 및 10분 단위로 변수 초기화 -> 10분 데이터를 디비에 저장할지는 미정

- **매도**
    - 10분마다 초기화 되며 시작가보다 시장가 차이가 -3% 이상일시 즉시 50퍼 매도
    - 데이터베이스에 매도 내역 저장
    - sync.WaitGroup{} 으로 나머지 연산들 대기
    - 대기할때 알고리즘 짜기
- **매수**
    - 10분마다 초기화 되며 시작가보다 시장가가 3% 올랐을시 즉시 ~~원 매수
    - 데이터베이스에 매수 내역 저장
    - sync.WaitGroup{} 으로 나머지 연산들 대기
    - 대기할때 알고리즘 짜기

### 기술스택

- MySQL
- GORM
- Go

### USAGE
