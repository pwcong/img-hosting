# 用户接口

* 用户注册
* 用户登录
* 验证令牌

## 用户注册

路径：`/user/register`

方法: `POST`

请求参数:

* username: 用户名
* password: 密码

请求示例：

```shell
curl -X POST \
  $API_BASE/user/register \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "xxx",
    "password": "yyy"
  }'
```

返回值：

```json
{
  "success": true,
  "message": "注册成功",
  "code": 20000,
  "payload": {
    "token": "xxx",
    "id": 1,
    "username": "xxx"
  }
}
```

---

## 用户登录

路径：`/user/login`

方法：`POST`

请求参数：

* username: 用户名
* password: 密码

请求示例：

```shell
curl -X POST \
  $API_BASE/user/login \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "xxx",
    "password": "yyy"
  }'
```

返回值：

```json
{
  "success": true,
  "message": "登录成功",
  "code": 20000,
  "payload": {
    "token": "xxx",
    "id": 1,
    "username": "xxx"
  }
}
```

---

## 验证令牌

路径：`/user/check`

方法：`POST`

请求头：

* Token: 用户令牌

请求示例：

```shell
curl -X POST \
  $API_BASE/user/check \
  -H 'Token: xxx'
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "验证成功",
  "payload": {
    "id": 10000,
    "token": "xxx",
    "username": "xxx"
  }
}
```
