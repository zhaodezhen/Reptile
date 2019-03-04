# golang爬虫

- [x] 城市列表解析器
- [x] 城市解析器
- [x] 用户解析器

- Parser解析器
```$xslt
Requests{URL,对应的parser}列表，Item
```
---
####单任务版运行流程

Seed 送给 Engine 一些request，engine把每个request加入了任务队列。每个任务队列里面的任务(URL)去调Fetch，然后等待Fetch返回。（对于每一个任务，它是单任务版。所以必须一个一个任务等fetch返回，然后在等待parse返回，把返回的request加入任务队列、返回的item打印出来）


