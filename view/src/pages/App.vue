<template>
    <div class="container">
        <nav-header :navs="navs">
            <a class="logo" href="/" slot="logo">
                <img :src="logo"/>
                <span>IMG HOSTING</span>
            </a>
        </nav-header>
        <div class="content">
            <transition name="fade">
                <router-view></router-view>
            </transition>
        </div>
    </div>
</template>

<style lang="scss">
html,
body {
  width: 100%;
  height: 100%;
  overflow: hidden;
  font-family: 'Microsoft YaHei', sans-serif;
}

.fade-enter-active,
.fade-leave-active {
  position: absolute;
  transition: opacity 0.5s;
}
.fade-enter,
.fade-leave-to {
  position: absolute;
  opacity: 0;
}

.container {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;

  .logo {
    display: inline-block;
    text-decoration: none;
    color: #20b2aa;
    font-size: 20px;
    margin: 0 16px;
    img {
      line-height: 50px;
      vertical-align: middle;
      width: 28px;
      height: 28px;
    }
    span {
      line-height: 50px;
      vertical-align: middle;
    }
  }

  .content {
    width: 100%;
    height: 100%;
    overflow: auto;
    padding-top: 50px;
    box-sizing: border-box;
  }
}
</style>

<script>
import NavHeader from '@/components/NavHeader';

import logo from '@/assets/imgs/logo.png';

export default {
  data() {
    return {
      logo,
      navs: [
        {
          to: '/',
          label: '首页'
        },
        {
          to: '/list',
          label: '个人'
        },
        {
          to: '/about',
          label: '关于'
        },
        {
          to: '/contact',
          label: '联系'
        },
        {
          to: '/api',
          label: 'API'
        }
      ]
    };
  },
  components: {
    NavHeader
  }
};
</script>
