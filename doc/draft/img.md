# 图片接口

* 上传图片
* 删除图片

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

* ids: 图片ID数组

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
