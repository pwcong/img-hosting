<template>
    <div class="header">
        <a class="header-brand" href="/">
            {{ brand }}
        </a>
        <router-link 
            class="header-menu"
            v-for="menu in menuList"
            :key="menu.path" 
            :to="menu.path">
            {{ menu.text }}
        </router-link>
        <div class="header-user"></div>
    </div>
</template>
<style>
    .header{
        width: 100%;
        height: 56px;

        background-color: #2c3e50;

        box-sizing: border-box;

        display: flex;
        padding-left: 24px;
        padding-right: 24px;

        position: relative;

    }

    .header-brand{
        font-size: 18px;
        color: white;

        box-sizing: border-box;
        -ms-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -moz-box-sizing: border-box;

        padding-left: 8px;
        padding-right: 8px;

        text-decoration: none;

        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .header-menu{

        color: white;

        font-size: 14px;

        box-sizing: border-box;
        -ms-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -moz-box-sizing: border-box;

        padding-left: 16px;
        padding-right: 16px;

        text-decoration: none;

        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;

        transition: color 0.3s;

    }


    .header-menu:hover{
        color: lightseagreen;
    }

    .router-link-exact-active{
        background-color: lightseagreen;
    }

    .router-link-exact-active:hover{
        color: #2c3e50;
    }

    .header-user{
        height: 100%;
        width: 200px;
        border: 1px red solid;

        position: absolute;
        right: 0px;
    }


</style>
<script>


    export default {
        props: {
            menuList: {
                type: Array,
                require: true
            },
            brand: {
                type: String,
                require: true
            }
        },
        data(){
            return {
            }
        },
        computed: {
            
        },
        methods: {
            
        },
        components: {
        }
        
    }
    
</script>