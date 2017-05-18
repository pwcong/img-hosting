<template>
    <div class="container">
        <my-header 
            :menu-list="menuList" 
            :brand="brand"
            :uid="uid"
            :logined="logined"
            :onSignInClickListener="handleShowLoginOrRegisterBox"
            :onSignOutClickListener="handleSignOut"
            />
        <router-view-container/>
        <my-footer :text="footerText"/>
        <transition name="fade">

            <login-or-register-box 
                :logining="logining"
                :registering="registering"
                :onRegisterClick="handleRegister"
                :onLoginClick="handleLogin"
                :onCancleLoginOrRegisterClick="handleHideLoginOrRegisterBox"
                v-if="loginOrRegisterBoxActive"
                />
        </transition>
        <tips 
            :showTips="showTips"
            :tipsContent="tipsContent"
            :tipsType="tipsType"/>
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

    .fade-enter-active, .fade-leave-active {
        transition: opacity .3s
    }

    .fade-enter, .fade-leave-active {
        opacity: 0
    }
</style>
<script>

    // import { router } from 'vue-router';

    import * as types from '../store/types.js';

    import MyHeader from '../component/MyHeader.vue';
    import MyFooter from '../component/MyFooter.vue';

    import RouterViewContainer from '../component/RouterViewContainer.vue';
    import LoginOrRegisterBox from '../component/LoginOrRegisterBox.vue';
    import Tips from '../component/Tips.vue';

    export default {
        data(){
            return {
                brand: "IMG HOSTING",
                footerText: "Copyleft © 2017  Pwcong",
                loginOrRegisterBoxActive: false
            }
        },
        computed: {

            menuList(){
                return [
                    { path: '/', text: '主页', active: true },
                    { path: '/own', text: '个人图库', active: this.$store.state.user.logined },
                    { path: '/about', text: '关于', active: true },
                    { path: '/contact', text: '联系', active: true },
                    { path: '/api', text: 'API', active: true },
                ]
            },
            uid(){
                return this.$store.state.user.uid;
            },
            logined(){
                return this.$store.state.user.logined;
            },
            logining(){
                return this.$store.state.user.logining;
            },
            registering(){
                return this.$store.state.user.registering;
            },
            showTips(){
                return this.$store.state.tip.show;
            },
            tipsContent(){
                return this.$store.state.tip.content;
            },
            tipsType(){
                return this.$store.state.tip.type;
            }
        },
        methods: {
            handleShowLoginOrRegisterBox(e){
                
                this.$data.loginOrRegisterBoxActive = true;

            },
            handleSignOut(e){

                this.$router.push("/");
                this.$store.dispatch(types.ACTION_USER_TOLOGOUT);

                this.$store.commit(types.MUTATION_TIP_SHOWTIPS, {
                    type: 0,
                    content: '已成功登出'
                })

            },
            handleHideLoginOrRegisterBox(e){
                this.$data.loginOrRegisterBoxActive = false;

            },
            handleLogin(uid, pwd){

                let ctx = this;

                if(uid != '' && pwd != ''){
                    ctx.$store.dispatch(types.ACTION_USER_TOLOGIN, {
                        uid,
                        pwd
                    }).then(() => {

                        ctx.$data.loginOrRegisterBoxActive = false;

                        ctx.$store.commit(types.MUTATION_TIP_SHOWTIPS, {
                            type: 1,
                            content: '已成功登陆'
                        })

                    }).catch(() => {

                        ctx.$store.commit(types.MUTATION_TIP_SHOWTIPS, {
                            type: 2,
                            content: '登陆失败，密码错误或用户不存在'
                        });
                    })
                }

            },
            handleRegister(uid, pwd){

                let ctx = this;

                if(uid != '' && pwd != ''){

                    ctx.$store.dispatch(types.ACTION_USER_TOREGISTER, {
                        uid,
                        pwd
                    }).then(() => {

                        ctx.$data.loginOrRegisterBoxActive = false;

                        ctx.$store.commit(types.MUTATION_TIP_SHOWTIPS, {
                            type: 1,
                            content: '注册成功'
                        });

                    }).catch(() => {

                        ctx.$store.commit(types.MUTATION_TIP_SHOWTIPS, {
                            type: 2,
                            content: '注册失败，用户已存在或服务器出了小差'
                        });
                    })
                    
                }
            },
        },
        components: {
            MyHeader,
            MyFooter,
            RouterViewContainer,
            LoginOrRegisterBox,
            Tips
        },
        mounted: function() {
            this.$store.dispatch(types.ACTION_USER_CHECK);
        }
        
    }
</script>