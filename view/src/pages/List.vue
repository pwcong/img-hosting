<template>
  <div class="list">
    <div v-if="!check" class="tip">
      <img :src="pleaseImg" />
      <div>
        您还未登录，请先登录
        <span>.</span>
        <span>.</span>
        <span>.</span>
      </div>
    </div>
  </div>
</template>
<style lang="scss">
.list {
  position: relative;
  width: 100%;
  .tip {
    padding-top: 80px;
    text-align: center;
    font-size: 20px;
    font-weight: bold;
    img {
      max-width: 200px;
    }

    span {
      position: relative;
      animation: dotanim 1.2s infinite;

      &:nth-child(2) {
        animation-delay: 0.3s;
      }
      &:nth-child(3) {
        animation-delay: 0.6s;
      }
    }
  }
}

@keyframes dotanim {
  0% {
    top: 0;
  }

  50% {
    top: -8px;
  }

  100% {
    top: 0;
  }
}
</style>
<script>
import pleaseImg from '@/assets/imgs/please.png';

export default {
  data() {
    return {
      pleaseImg
    };
  },
  methods: {},
  computed: {
    check() {
      return this.$store.getters.check;
    }
  }
};
</script>
