<template>
    <div class="markdown-body content" v-html="rawHtml">
        
    </div>
</template>
<style>

</style>
<script>
import marked from 'marked';

const text = `
    # 关于 IMG Hosting
    IMG Hosting 是一个免费的公共图床，请在法律允许范围内随意使用。
    
    项目开源地址：[https://github.com/pwcong/img-hosting](https://github.com/pwcong/img-hosting)
    # 严禁上传及分享图片类型
    * 含有色情、暴力、宣扬恐怖主义的图片
    * 侵犯版权、未经授权的图片
    * 其他违反中华人民共和国法律的图片
    * 其他违反香港法律的图片
    `;

export default {
  data() {
    return {
      rawHtml: marked(text)
    };
  }
};
</script>