#### 分页支持

部分接口支持数据分页

分页参数：

* pageSize: 每页个数
* pageNo: 页码

请求示例：

```shell
curl -X GET $API_BASE/imgs/?page_size=20&page_no=1 \
  -H "Token: xxx"
```

相应结果：

```json
{
  "success": true,      // 是否处理成功
  "message": "xxx",     // 错误信息
  "code": 20000,        // 响应码
  "payload": {          // 主要内容
    "data": [],
    "page_size": 5,
    "page_no": 1,
    "total_size": 1,
    "total_no": 1
  }
}


```

#### 统一请求返回内容（JSON）

无论请求是否处理成功均返回内容，内容模板如下：

```json
{
  "success": true,      // 是否处理成功
  "message": "xxx",     // 错误信息
  "code": 20000,        // 响应码
  "payload": {          // 主要内容
    ...
  }
}
```

#### 响应码

| 数值  | 含义      |
| ----- | -------- |
| 20000 | 请求成功  |
| 40000 | 请求失败  |
