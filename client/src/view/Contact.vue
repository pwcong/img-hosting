<template>
    <div class="markdown-body" v-html="rawHtml">
    </div>
</template>
<style>

</style>
<script>

    import marked from 'marked';

    const text = `
    # 联系我
    有任何疑问或者问题，请通过以下方式联系我：
    * 扣扣：429734712
    * 邮箱：pwcong@foxmail.com
    `

    export default {
        data(){
            return {
                rawHtml: marked(text)
            }
        }
        
    }
    
</script>