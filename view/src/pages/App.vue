<template>
    <div class="container">
        <img :src="logo" alt="Logo">
        <div class="navs">
            <router-link class="nav" to="/">Home</router-link>
            <router-link class="nav" to="/counter">Counter</router-link>
        </div>
        <div class="content">
            <router-view></router-view>
        </div>

    </div>
</template>

<style lang="scss">
body {
    font-family: 'Microsoft YaHei', sans-serif;
}

.container {

    text-align: center;
    padding-top: 60px;

    .navs {
        border-bottom: 1px solid #ddd;

        .nav {
            display: inline-block;
            padding: 8px 0;
            width: 100px;
            text-align: center;
            color: #666;
            text-decoration: none;

            border-bottom: 2px solid transparent;

            transition: border 0.3s, color 0.3s;

            &:hover {
                color: #41B883;
                border-bottom: 2px solid #41B883;
            }

            &.router-link-exact-active {
                color: #41B883;
                border-bottom: 2px solid #41B883;
            }
        }
    }

    .content {
        margin-top: 50px;
    }
}
</style>

<script>

import logo from '../assets/imgs/vue.png';

export default {

    data() {
        return {
            logo
        }
    }

}
</script>
