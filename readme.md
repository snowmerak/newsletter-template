# newsletter-generator

간단한 뉴스레터 생성기입니다.

## HowTo

레포를 클론하는 것으로 시작합니다.

이 생성기는 하나의 뉴스레터와 여러개의 아티클을 연결하여 사용합니다.  
일단 뉴스레터를 생성합니다.

```bash
go run . generate <newsletter-name>
```

이 명령을 실행하면 해당 이름으로 뉴스레터가 생성됩니다.

```yaml
title: Test
date: "2022-01-20"
articles: []
template: |-
  <tr>
  		<td align="center">
  			<h1>{{.Title}}</h1>
  		</td>
  	</tr>
  	<tr>
  		<td align="center">
  			<p>{{.Date}}</p>
  		</td>
  	</tr>
```

타이틀과 날짜, 아티클 목록, 그리고 템플릿이 있습니다. 해당 템플릿을 기반으로 입력된 값을 출력합니다.  
여기서 아티클 목록은 문자열 배열이고 파일 확장자를 제외한 이름만 입력하시면 됩니다.

다음으로 아티클을 생성합니다.

```bash
go run . article new <article-name>
```

이 명령을 실행하면 해당 아티클 이름으로 아티클이 생성됩니다.

```yaml
title: index
author: snowmerak
tags: []
image: ""
link: ""
Content: ""
template: |-
  <tr>
  		<td align="center">
  			<h2>{{.Title}}</h2>
  		</td>
  	</tr>
  	<tr>
  		<td align="center">
  			<em>{{.Author}}</em>
  		</td>
  	</tr>
  	<tr>
  		<td align="center">
  			<img src="{{.Image}}" alt="{{.Title}}" width="380">
  		</td>
  	</tr>
  	<tr>
  		<td align="center">
  			<a href="{{.Link}}">symbolic link</a>
  		</td>
  	</tr>
  	<tr>
  		<td>
  			<p>{{.Content}}</p>
  		</td>
  	</tr>
```

타이틀과 저자, 이미지와 링크, 태그와 컨텐트, 그리고 마찬가지로 템플릿이 있습니다.  
이미지는 기본 템플릿에서는 보여줄 이미지의 주소로 쓰입니다.  
링크는 해당 아티클의 외부 링크입니다.  
태그는 문자열 배열로 키워드들입니다.  
컨텐트는 마크다운을 기반으로 하여 아티클의 본문이 됩니다.

마지막으로 생성된 아티클의 이름을 뉴스레터 파일에 추가하고 다음 명령을 실행하면 dist 파일에 뉴스레터 이름의 html 파일이 생성됩니다.

```bash
go run . generate <newsletter-name>
```
