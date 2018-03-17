<template>
    <div class="api markdown-body" v-html="rawHTML">
    </div>
</template>
<style lang="scss">
.api {
  padding: 32px;
}
</style>
<script>
import marked from 'marked';

const content = `

### API文档

> API文档地址 -> [Doc For IMG HOSTING](/doc)

#### 分页支持

部分接口支持数据分页

分页参数：

* pageSize: 每页个数
* pageNo: 页码

请求示例：

\`\`\`shell
curl -X GET $API_BASE/imgs/?page_size=20&page_no=1 \
  -H "Token: xxx"
\`\`\`

相应结果：

\`\`\`json
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


\`\`\`

#### 统一请求返回内容（JSON）

无论请求是否处理成功均返回内容，内容模板如下：

\`\`\`json
{
  "success": true,      // 是否处理成功
  "message": "xxx",     // 错误信息
  "code": 20000,        // 响应码
  "payload": {          // 主要内容
    ...
  }
}
\`\`\`

#### 响应码

| 数值  | 含义      |
| ----- | -------- |
| 20000 | 请求成功  |
| 40000 | 请求失败  |

******

### 用户接口

* 用户注册
* 用户登录
* 验证令牌

#### 用户注册

路径：\`/user/register\`

方法: \`POST\`

请求参数:

* username: 用户名
* password: 密码

请求示例：

\`\`\`shell
curl -X POST \
  $API_BASE/user/register \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "xxx",
    "password": "yyy"
  }'
\`\`\`

返回值：

\`\`\`json
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
\`\`\`

---

#### 用户登录

路径：\`/user/login\`

方法：\`POST\`

请求参数：

* username: 用户名
* password: 密码

请求示例：

\`\`\`shell
curl -X POST \
  $API_BASE/user/login \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "xxx",
    "password": "yyy"
  }'
\`\`\`

返回值：

\`\`\`json
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
\`\`\`

---

#### 验证令牌

路径：\`/user/check\`

方法：\`POST\`

请求头：

* Token: 用户令牌

请求示例：

\`\`\`shell
curl -X POST \
  $API_BASE/user/check \
  -H 'Token: xxx'
\`\`\`

返回值：

\`\`\`json
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
\`\`\`

******

### 图片接口

* 上传图片
* 删除图片

******

#### 上传图片

路径：\`/img/upload\`

方法: \`POST\`

请求参数:

* img: 图片文件

请求头：

* Token: 用户令牌(可选)

请求示例：

\`\`\`shell
curl -X POST $API_BASE/img/upload \
  -H 'Token: xxx' \
  -H 'content-type: multipart/form-data; boundary=xxx' \
  -F 'img=@xxx'
\`\`\`

返回值：

\`\`\`json
{
  "success": true,
  "message": "上传成功",
  "code": 20000,
  "payload": {
    "id": 1,
    "url": "xxx"
  }
}
\`\`\`

---

### 删除图片

路径：\`/img/remove\`

方法: \`POST\`

请求参数:

* ids: 图片ID数组

请求头：

* Token: 用户令牌

请求示例：

\`\`\`shell
curl -X POST $API_BASE/img/remove \
  -H 'Token: xxx' \
  -d '{
    ids: []
  }'
\`\`\`

返回值：

\`\`\`json
{
  "success": true,
  "message": "删除成功",
  "code": 20000,
  "payload": {}
}
\`\`\`

`;

export default {
  data() {
    return {
      rawHTML: marked(content)
    };
  }
};
</script>