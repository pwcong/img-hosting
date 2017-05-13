<template>
    <div class="container">
        <my-header :menu-list="menuList" :brand="brand"/>
        <router-view-container/>
        <my-footer :text="footerText"/>
    </div>
</template>
<style>
    .container{
        position: fixed;
        left: 0;
        top: 0;

        width: 100%;
        height: 100%;

        display: flex;
        flex-direction: column;
    }
</style>
<script>

    import MyHeader from '../component/MyHeader.vue';
    import MyFooter from '../component/MyFooter.vue';

    import RouterViewContainer from '../component/RouterViewContainer.vue';

    export default {
        data(){
            return {
                menuList: [
                    { path: '/', text: '主页' },
                    { path: '/about', text: '关于' },
                    { path: '/contact', text: '联系' },
                    { path: '/api', text: 'API' }
                ],
                brand: "IMG HOSTING",
                footerText: "Copyleft © 2017  Pwcong"
            }
        },
        computed: {
            
        },
        methods: {
            
        },
        components: {
            MyHeader,
            MyFooter,
            RouterViewContainer
        }
        
    }
</script>