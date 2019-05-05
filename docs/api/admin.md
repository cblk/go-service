客户任务后台管理系统 APIs v1
===============

##  任务分页查询

### URL
 
GET `/tasks`

### INPUT
| 参数名称  | 是否必填 | 数据类型 | 参数说明 |
| --- | --- | --- | --- |
| page  | 是 | int | 页数 |
| size  | 是 | int | 每页数量 |
| type  | 否 | string | 任务类型(article,image) |
| status  | 否 | string | 任务类型(success, faild, pending) |
| err_type  | 否 | int | 任务错误类型(404，405，自定义等) |
| created_at  | 否 | int | 任务开始时间|
| finished_at  | 否 | int | 任务结束时间|
| app_id  | 否 | string | 客户方ID |



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
    "data": [{
        "app_id": "调用方",
        "task_id": "123456",
        "url": "https://article.luban.site/id",
        "start_at": "123456", 
        "finished_at": "123456", 
        "status": "success|faild|pending",
        "err_type": "404"
    }]
}

```

### 状态码定义

| 状态码 | 描述 |
| --- | --- | 
| 0  | 成功 |
| 10000   | 网络错误 | 


##  任务统计查询

### URL
 
GET `/tasks`

### INPUT
| 参数名称  | 是否必填 | 数据类型 | 参数说明 |
| --- | --- | --- | --- |
| start_time  | 是 | int | 任务开始时间|
| end_time  | 是 | int | 任务结束时间|
| type  | 否 | string | 任务类型(article,image, all) |
| time_type  | 是 | string | 任务类型(小时,天,周) |
| err_type  | 否 | int | 任务错误类型(404，405，0等) |
| status  | 否 | string | 任务类型(success, faild, pending, all) |
| app_id  | 否 | string | 客户方ID |


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
    "data": {
        "total": 100,
        "header": ["name","2019-01-01","2019-01-01"],
        "data":{
            "avg": { // 平均时间
                "2019-01-01":2,
                "2019-01-02":2,
                "2019-01-03":3,
                "2019-01-04":10
            },
            "max": { // 最大时间
                "2019-01-01":2,
                "2019-01-02":2,
                "2019-01-03":3,
                "2019-01-04":10
            },
            "min": { // 最小时间
                "2019-01-01":2,
                "2019-01-02":2,
                "2019-01-03":3,
                "2019-01-04":10
            }
        }
    }
}

```

### 状态码定义

| 状态码 | 描述 |
| --- | --- | 
| 0  | 成功 |
| 10000   | 网络错误 | 