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

    import * as types from '../store/types.js';

    import MyHeader from '../component/MyHeader.vue';
    import MyFooter from '../component/MyFooter.vue';

    import RouterViewContainer from '../component/RouterViewContainer.vue';
    import LoginOrRegisterBox from '../component/LoginOrRegisterBox.vue';

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
                footerText: "Copyleft © 2017  Pwcong",
                loginOrRegisterBoxActive: false
            }
        },
        computed: {
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
        },
        methods: {
            handleShowLoginOrRegisterBox(e){
                
                this.$data.loginOrRegisterBoxActive = true;

            },
            handleSignOut(e){

                this.$store.dispatch(types.ACTION_USER_TOLOGOUT);

            },
            handleHideLoginOrRegisterBox(e){
                this.$data.loginOrRegisterBoxActive = false;

            },
            handleLogin(uid, pwd){

                if(uid != '' && pwd != ''){
                    this.$store.dispatch(types.ACTION_USER_TOLOGIN, {
                        uid,
                        pwd
                    }).then(() => {

                        this.$data.loginOrRegisterBoxActive = false;


                    }).catch(() => {

                    })
                }

            },
            handleRegister(uid, pwd){

                if(uid != '' && pwd != ''){

                    this.$store.dispatch(types.ACTION_USER_TOREGISTER, {
                        uid,
                        pwd
                    }).then(() => {

                        this.$data.loginOrRegisterBoxActive = false;


                    }).catch(() => {

                    })
                    
                }
            }
        },
        components: {
            MyHeader,
            MyFooter,
            RouterViewContainer,
            LoginOrRegisterBox
        }
        
    }
</script>