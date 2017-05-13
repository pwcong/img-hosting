<template>
    <div class="markdown-body" v-html="rawHtml">
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
    
    `

    export default {
        data(){
            return {
                rawHtml: marked(text)
            }
        }
        
    }
    
</script>