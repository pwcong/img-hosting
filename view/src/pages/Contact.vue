<template>
    <div class="contact markdown-body" v-html="rawHTML">
    </div>
</template>
<style lang="scss">
.contact {
  width: 100%;
  box-sizing: border-box;
  padding: 32px;
}
</style>
<script>
import marked from 'marked';

const content = `
### 联系我

有任何疑问或者问题，请通过以下方式联系我：
* QQ：2209594033
* 邮箱：pwcong@foxmail.com
`;

export default {
  data() {
    return {
      rawHTML: marked(content)
    };
  }
};
</script>
