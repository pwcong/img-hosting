<template>
  <div class="footer">
    <span style="display: inline-block; transform: rotate(180deg); margin-right: 4px;">©</span>{{content}}
  </div>
</template>
<style lang="scss">
.footer {
  position: absolute;
  left: 0;
  bottom: 0;
  width: 100%;
  box-sizing: border-box;
  color: #666;
  font-size: 13px;
  text-align: center;
  padding: 16px 32px;

  transition: opacity 0.3s;

  &:hover {
    opacity: 0;
  }
}
</style>
<script>
export default {
  props: {
    startYear: {
      type: Number,
      default: 2017
    },
    author: {
      type: String,
      default: 'Pwcong'
    }
  },
  data() {
    const ctx = this;

    const now = new Date();

    const t = now.getFullYear() > ctx.startYear ? ctx.startYear + ' - ' + now.getFullYear() : ctx.startYear;

    return {
      content: `Copyleft ${t} ${ctx.author}`
    };
  }
};
</script>
