# go_nomad
https://nomadcoders.co/go-for-beginners/


- `goroutine`을 활용한 `scrapper`와 `gin`을 활용한 웹서버 튜토리얼
  - 원하는 언어 스펙에 대해 검색을 하게 되면, 이에 대한 스크래핑 데이터를 json 또는 csv로 저장할 수 있게 해줍니다.
  - goroutine으로 검색된 언어에 대한 `페이지네이션` * `Max_article_size(50)`의 고루틴이 생성되어 스크래핑을 실시하며, channel을 활용해 하나의 메인 파일에 저장시킵니다.


![](./assets/images/s1.png)

![](./assets/images/s2.png)