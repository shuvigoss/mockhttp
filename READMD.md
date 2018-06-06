# 原理
`mockhttp -dir=/User/username/xxxdir -port=8080` 

会解析xxxdir下的所有文件

文件格式：
```
/v1/path
换行
{json} 返回json内容
```

`curl -i http://localhost:8080/v1/path`