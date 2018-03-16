<template>
    <div class="markdown-body content" v-html="rawHtml">
    </div>
</template>
<style>

</style>
<script>
import marked from 'marked';

const text = `
    # API 文档

    ## 图片上传

    * URL: http://${window.location.host}/img/upload
    * Method: POST

    **请求参数**

    |参数名称	 |类型	|是否必须	 |描述	     	| 
    |-----------|--------|-----------|----------------|
    |img		|File	 |是 		|上传的图片		 |
    |uid		|String  |否			|标记图片拥有者   |

    **返回数据说明（JSON格式）**

    |名称		   |类型	 |描述					|
    |-----------|-------|------------------------|
    |code		|Number	|响应状态码               |
    |msg		|String |响应信息				  |
    |filename	|String |上传的图片文件名 		   |
    |path		|String |图片文件保存路径		   |
    |url 		|String |生成的图片URL地址  	    |

    **响应示例**

    \`\`\`
        {
            "filename": "59432931_p0.jpg",
            "path": "/2017/05/16/b45b2db6425f.jpg",
            "url": "http://localhost/public/2017/05/16/b45b2db6425f.jpg",
            "code": 200,
            "msg": ""
        }

    \`\`\`

    ## 获取用户上传图片

    * URL: http://${window.location.host}/img/list
    * Method: POST

    **请求参数**

    |参数名称	 |类型	|是否必须	 |描述	     	     | 
    |-----------|--------|-----------|---------------------|
    |token		|String	 |是 		|用户令牌（需登录获取） |

    **返回数据说明（JSON格式）**

    |名称		   |类型	 |描述					|
    |-----------|-------|------------------------|
    |code		|Number	|响应状态码               |
    |msg		|String |响应信息				  |
    |imgs	    |Array  |用户上传图片列表 		   |

    **响应示例**

    \`\`\`
        {
            "imgs": [
                {
                "filename": "59432931_p0.jpg",
                "year": "2017",
                "month": "05",
                "day": "16",
                "storename": "b45b2db6425f.jpg",
                "path": "/2017/05/16/b45b2db6425f.jpg",
                "url": "http://localhost/public/2017/05/16/b45b2db6425f.jpg"
                }
            ],
            "code": 200,
            "msg": ""
        }
    \`\`\`

    ## 删除用户上传图片

    * URL: http://${window.location.host}/img/delete
    * Method: POST

    **请求参数**

    |参数名称	 |类型	|是否必须	 |描述	     	    | 
    |-----------|--------|-----------|--------------------|
    |token		|String	 |是 		|用户令牌（需登录获取）|
    |year		|String	 |是 		|年                  |
    |month		|String	 |是 		|月                  |
    |day		|String	 |是 		|日                  |
    |storename	|String	 |是 		|服务端图片文件名     |

    **返回数据说明（JSON格式）**

    |名称		   |类型	 |描述					|
    |-----------|-------|------------------------|
    |code		|Number	|响应状态码               |
    |msg		|String |响应信息				  |

    **响应示例**

    \`\`\`
        {
            "code": 200,
            "msg": ""
        }
    \`\`\`

    ## 用户登录

    * URL: http://${window.location.host}/user/login
    * Method: POST

    **请求参数**

    |参数名称	 |类型	|是否必须	 |描述	 | 
    |-----------|--------|-----------|---------|
    |uid		|String	 |是 		|用户名    |
    |pwd		|String	 |是 		|密码    |

    **返回数据说明（JSON格式）**

    |名称        |类型   |描述	         |
    |-----------|-------|-----------------|
    |code		|Number	|响应状态码        |
    |msg		|String |响应信息          |
    |token	    |String |用户令牌 		   |

    **响应示例**

    \`\`\`
        {
            "token": "eyJhbGciOiJIUzI1NiIsInxpXVCJ9.eyJ1aWQiOiJwd2NvbmciLCJleHAiOjE1MDIxMjE4NjV9.dmcNHDncLnDt3ISGGRhcsJNGCcSlm6NijGn5efz_KwE",
            "code": 200,
            "message": ""
        }
    \`\`\`

    ## 用户注册

    * URL: http://${window.location.host}/user/register
    * Method: POST

    **请求参数**

    |参数名称	 |类型	|是否必须	 |描述	 | 
    |-----------|--------|-----------|---------|
    |uid		|String	 |是 		|用户名    |
    |pwd		|String	 |是 		|密码    |

    **返回数据说明（JSON格式）**

    |名称		   |类型	 |描述		  |
    |-----------|-------|--------------|
    |code		|Number	|响应状态码     |
    |msg		|String |响应信息       |
    |token	    |String |用户令牌       |

    **响应示例**

    \`\`\`
        {
            "token": "eyJhbGciOiJIUzI1NiIsInxpXVCJ9.eyJ1aWQiOiJwd2NvbmciLCJleHAiOjE1MDIxMjE4NjV9.dmcNHDncLnDt3ISGGRhcsJNGCcSlm6NijGn5efz_KwE",
            "code": 200,
            "message": ""
        }
    \`\`\`

    
    `;

export default {
  data() {
    return {
      rawHtml: marked(text)
    };
  }
};
</script>