客户任务管理系统 APIs v1
===============

## 提交任务

### URL
 
POST `/tasks`

### INPUT


| 参数名称  | 是否必填 | 数据类型 | 参数说明 |
| --- | --- | --- | --- |
| type   | 是 | enum | 任务类型--文章(article),图片(image)|
| url    | 是 | string | 任务url |
| app_id | 是 | string | 调用方ID |


```
{
    "type":"article",
    "url":"https://mp.weixin.qq.com/s?src=11&timestamp=1556420421&ver=1573&signature=eeaZ8UDyUIkCiXQJ14rXE6xxAwrYLHdNwaKLHzlpABFTzJWx5TVnakGI1UxpRKlSG7h05PDzHrKJYn7cc-JC7mWFoNi8V9-TL7Uu8wk4gSnovc3oxzTqVY*v*q3xYq8b&new=1",
    "app_id": "id"
}
```

### OUTPUT

| 参数名称  | 是否必填 | 数据类型 | 参数说明 |
| --- | --- | --- | --- |
| code  | 是 | int | 状态码 |
| msg   | 否 | string | 错误消息 |
| data   | 是 | object | 相应数据 |
| data[id] | 是 | string | Task ID |

```
# 全部信息
{
    "code": 0,
    "msg": "",
    "data": {
        "id": "83ud73ue83jd8"
    }
}

{
    "code":   10002,
    "msg": "数据参数转化错误",
    "data": null
}
```

### 状态码定义

| 状态码 | 描述 |
| --- | --- | 
| 0  | 成功 |
| 10000   | 网络错误 | 

##  查询任务

### URL
 
GET `/tasks/:id`

### INPUT
| 参数名称  | 是否必填 | 数据类型 | 参数说明 |
| --- | --- | --- | --- |
| id  | 是 | string | 任务ID |

### OUTPUT

| 参数名称  | 是否必填 | 数据类型 | 参数说明 |
| --- | --- | --- | --- |
| code  | 是 | int | 状态码 |
| msg   | 是 | string | 错误消息 |
| data   | 是 | object | data |
| data[url]   | 是 | string | 新生成的url |

```
# 全部信息
{
    "code": 0,
    "msg": "",
    "data": {
        "url": "https://article.luban.site/id" 
    }
}

```

### 状态码定义

| 状态码 | 描述 |
| --- | --- | 
| 0  | 成功 |
| 10000   | 网络错误 | 



##  更新任务

### URL

PUT `/tasks/:id/status`

### INPUT
| 参数名称  | 是否必填 | 数据类型 | 参数说明 |
| --- | --- | --- | --- |
| id  | 是 | string | 任务ID |
| url | 是 | string | 抓取后的url |
```
# 全部信息
{
  "url": "https://article.luban.site/id",
  "status": "success|failed|pending",
  "err_type": "404|500" 
}
```

### OUTPUT

| 参数名称  | 是否必填 | 数据类型 | 参数说明 |
| --- | --- | --- | --- |
| code  | 是 | int | 状态码 |
| msg   | 是 | string | 错误消息 |

```
# 全部信息
{
    "code": 0,
    "msg": ""
}

```

### 状态码定义

| 状态码 | 描述 |
| --- | --- | 
| 0  | 成功 |
| 10000   | 网络错误 | 