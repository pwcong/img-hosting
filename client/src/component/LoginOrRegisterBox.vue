<template>
    <div class="mask-container">

        <div class="box box-login" v-if="toggle">
            <div class="box-login-avatar">
                <img src="/imgs/logo.png"/>
            </div>
            <div class="box-tools">

                <div class="box-tools-item box-tools-input">
                    <div class="box-tools-input-icon">
                        <span class="fa fa-user"></span>
                    </div>
                    <div class="box-tools-input-main">
                        <input type="text" v-model="loginUID"/>
                    </div>
                </div>

                <div class="box-tools-item box-tools-input">
                    <div class="box-tools-input-icon">
                        <span class="fa fa-key"></span>
                    </div>
                    <div class="box-tools-input-main">
                        <input type="password" v-model="loginPWD"/>
                    </div>
                </div>

                <div 
                    :class="{
                        'box-tools-item': true, 
                        'box-tools-btn': true,
                        'box-tools-btn-one': !logining,
                        'box-tools-btn-disable': logining
                    }"
                    @click="onLoginClickListener">
                    {{ logining ? "登陆中" : "登陆"}}
                </div>
                <div 
                    :class="{
                        'box-tools-item': true, 
                        'box-tools-btn': true,
                        'box-tools-btn-two': !logining,
                        'box-tools-btn-disable': logining
                    }" 
                    @click="toggleBox">
                    没有账号？注册一个
                </div>
                <div 
                    :class="{
                        'box-tools-item': true, 
                        'box-tools-btn': true,
                        'box-tools-btn-cancle': !logining,
                        'box-tools-btn-disable': logining
                    }"
                    @click="onCancleLoginOrRegisterClick">
                    取消
                </div>
            </div>
        </div>


        <div class="box box-register" v-if="!toggle">

            <div class="box-tools">

                <div class="box-tools-item box-tools-input">
                    <div class="box-tools-input-icon">
                        <span class="fa fa-user"></span>
                    </div>
                    <div class="box-tools-input-main">
                        <input type="text" v-model="registerUID"/>
                    </div>
                </div>

                <div class="box-tools-item box-tools-input">
                    <div class="box-tools-input-icon">
                        <span class="fa fa-key"></span>
                    </div>
                    <div class="box-tools-input-main">
                        <input type="password" v-model="registerPWD"/>
                    </div>
                </div>

                <div 
                    :class="{
                        'box-tools-item': true, 
                        'box-tools-btn': true,
                        'box-tools-btn-one': !registering,
                        'box-tools-btn-disable': registering
                    }"
                    @click="onRegisterClickListener">
                    注册
                </div>
                <div 
                    :class="{
                        'box-tools-item': true, 
                        'box-tools-btn': true,
                        'box-tools-btn-cancle': !registering,
                        'box-tools-btn-disable': registering
                    }"
                    @click="toggleBox">
                    返回
                </div>
            </div>

        </div>

    </div>
  
</template>
<style>

    .mask-container{

        position: fixed;
        left: 0;
        top: 0;

        width: 100%;
        height: 100%;

        z-index: 999;

        background-color: rgba(0, 0, 0, 0.5);

        display: flex;
        align-items: center;
        justify-content: center;


    }

    .box{
        background-color: white;
        border-radius: 4px;

        box-shadow: 0 0 10px 0;

        box-sizing: border-box;
        -moz-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -ms-box-sizing: border-box;

        padding: 16px;

        display: flex;
        flex-direction: column;
    }

    .box-login{

        width: 300px;
        height: 350px;

        position: relative;
    }

    .box-register{

        width: 300px;
        height: 200px;

        position: relative;
    }

    .box-login-avatar{
        width: 100%;
        height: 96px;
        text-align: center;
    }
    .box-login-avatar img{
        max-width: 156px;
        position: relative;
        top: -96px;
    }
    .box-tools{
        flex: 1;
        width: 100%;

        display: flex;
        flex-direction: column;
        justify-content: space-around;

    }

    .box-tools-item{

        width: 100%;

    }

    .box-tools-btn{
        display: flex;
        flex-flow: row nowrap;
        align-items: center;

        user-select: none;
        -webkit-user-select: none;
        -ms-user-select: none;
        -moz-user-select: none;

        cursor: pointer;
        
        display: flex;
        align-items: center;
        justify-content: center;

        height: 32px;
        width: 100%;
        
        color: white;


        transition: background-color 0.3s;
    }

    .box-tools-btn-disable{
        pointer-events: none;
        background-color: #ddd;
    }

    .box-tools-btn-one{
        background-color: lightseagreen;
    }
    .box-tools-btn-one:hover{
        background-color: #31c3bb;
    }
    .box-tools-btn-one:active{
        background-color: #10a199;
    }

    .box-tools-btn-two{
        background-color: #2c3e50;
    }
    .box-tools-btn-two:hover{
        background-color: #3d4f61;
    }
    .box-tools-btn-two:active{
        background-color: #1b2d40;
    }

    .box-tools-btn-cancle{
        background-color: #ccc;
    }
    .box-tools-btn-cancle:hover{
        background-color: #ddd;
    }
    .box-tools-btn-cancle:active{
        background-color: #bbb;
    }

    .box-tools-input{
        display: flex;
        flex-flow: row nowrap;
        align-items: center;
        color: #2c3e50;
    }

    .box-tools-input-icon{
        font-size: 24px;


        width: 20%;
        height: 100%;
        text-align: center;
    }

    .box-tools-input-main{
        flex: 1;
        height: 100%;
    }

    .box-tools-input-main input{
        width: 100%;
        height: 32px;
        border: none;
        border: 1px #ccc solid;
        outline: none;

        font-size: 18px;

        box-sizing: border-box;
        -moz-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -ms-box-sizing: border-box;
        padding-left: 8px;

        transition: border 0.3s;

    }
    .box-tools-input-main input:hover{
        border: 1px lightseagreen solid;
    }


</style>
<script>
    export default {

        props: {
            logining: {
                type: Boolean,
                default: false
            },
            registering: {
                type: Boolean,
                default: false
            },
            onLoginClick: {
                type: Function,
                default: function(uid, pwd){
                    console.log(uid + ' ' + pwd);
                }
            },
            onRegisterClick: {
                type: Function,
                default: function(uid, pwd){
                    console.log(uid + ' ' + pwd);
                }
            },
            onCancleLoginOrRegisterClick: {
                type: Function,
                default: function(e){
                    console.log('cancle');
                }
            }
        },
        data(){
            return {
                toggle: true,
                loginUID: '',
                loginPWD: '',
                registerUID: '',
                registerPWD: '',
            }
        },
        methods: {

            toggleBox(e){
                this.$data.toggle = !this.$data.toggle;
                this.$data.loginUID = '';
                this.$data.loginPWD = '';
                this.$data.registerUID = '';
                this.$data.registerPWD = '';
            },
            onLoginClickListener(e){
                this.$props.onLoginClick(this.$data.loginUID, this.$data.loginPWD);
            },
            onRegisterClickListener(e){
                this.$props.onRegisterClick(this.$data.registerUID, this.$data.registerPWD);

            },


        }
    
    }
</script>

