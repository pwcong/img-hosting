# 图片接口

* 上传图片
* 删除图片
* 获取图片列表（需登录）

## 上传图片

路径：`/img/upload`

方法: `POST`

请求参数:

* img: 图片文件

请求头：

* Token: 用户令牌(可选)

请求示例：

```shell
curl -X POST $API_BASE/img/upload \
  -H 'Token: xxx' \
  -H 'content-type: multipart/form-data; boundary=xxx' \
  -F 'img=@xxx'
```

返回值：

```json
{
  "success": true,
  "message": "上传成功",
  "code": 20000,
  "payload": {
    "id": 1,
    "url": "xxx"
  }
}
```

---

## 删除图片

路径：`/img/remove`

方法: `POST`

请求参数:

* ids: 图片 ID 数组

请求头：

* Token: 用户令牌

请求示例：

```shell
curl -X POST $API_BASE/img/remove \
  -H 'Token: xxx' \
  -d '{
    ids: []
  }'
```

返回值：

```json
{
  "success": true,
  "message": "删除成功",
  "code": 20000,
  "payload": {}
}
```

---

## 获取图片列表（需登录）

路径：`/imgs`

方法: `GET`

请求头：

* Token: 用户令牌

其他：

* 支持分页

请求示例：

```shell
curl -X GET $API_BASE/imgs \
  -H 'Token: xxx'
```

返回值：

```json
{
  "success": true,
  "message": "查询成功",
  "code": 20000,
  "payload": {
    "data": [
      {
        "id": 1,
        "url": "/public/2018/03/17/xxx.jpg",
        "created_at": "2018-03-17T12:22:59+08:00"
      }
    ],
    "page_size": 1,
    "page_no": 1,
    "total_size": 1,
    "total_no": 1
  }
}
```

---
