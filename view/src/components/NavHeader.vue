<template>
  <div class="nav">
    <slot name="logo"></slot>
    <router-link 
      class="nav-item" 
      v-for="(nav, idx) in navs"
      :to="nav.to" 
      :key="'nav-item-' + idx">
      {{nav.label}}
    </router-link>
    <slot name="toolbar"></slot>
  </div>
</template>

<style lang="scss">
.nav {
  background: white;
  z-index: 100;
  height: 50px;
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  box-shadow: 0 1px 4px 1px #eee;

  .nav-item {
    display: inline-block;
    height: 50px;
    line-height: 50px;
    width: 80px;
    text-align: center;
    color: #666;
    text-decoration: none;

    box-sizing: border-box;

    font-size: 15px;

    border-bottom: 2px solid transparent;

    transition: border 0.3s, color 0.3s;

    &:hover {
      color: #20b2aa;
      border-bottom: 2px solid #20b2aa;
    }

    &.router-link-exact-active {
      color: #20b2aa;
      border-bottom: 2px solid #20b2aa;
    }
  }
}
</style>

<script>
export default {
  props: {
    navs: {
      type: Array,
      default() {
        return [
          {
            to: '/',
            label: '首页'
          }
        ];
      }
    }
  }
};
</script>

