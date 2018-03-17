<template>
    <div class="container">
        <nav-header :navs="navs">
            <a class="logo" href="/" slot="logo">
                <img :src="logo"/>
                <span>IMG HOSTING</span>
            </a>
            <div class="toolbar" slot="toolbar">
              <span v-if="check">Hi~ {{username}}</span>
              <span class="btn" v-if="!check" @click="switchToLogin">登录</span>
              <span class="btn" v-if="check" @click="handleToLogout">登出</span>
            </div>
        </nav-header>
        <div class="content">
            <transition name="fade">
                <router-view></router-view>
            </transition>
        </div>
        <simple-footer/>

        <transition name="fade">
          <div class="dialog" v-if="dialogVisible">
            <div class="form">
              <div :class="{
                  'form-item': true,
                  'form-item-input': true,
                  'warn': !formValid.username
                }">
                <span class="icon icon-username" :style="{
                  'background-image': 'url(' + usernameIcon + ')'
                }"></span>
                <input v-model="form.username" placeholder="请输入用户名"/>
              </div>
              <div :class="{
                  'form-item': true,
                  'form-item-input': true,
                  'warn': !formValid.password
                }">
                <span class="icon icon-password" :style="{
                  'background-image': 'url(' + passwordIcon + ')'
                }"></span>
                <input v-model="form.password" type="password" placeholder="请输入密码"/>
              </div>

              <div v-if="dialogSwitch" :class="{
                  'form-item': true,
                  'form-item-input': true,
                  'warn': !formValid.password2
                }">
                <input v-model="form.password2" type="password" placeholder="请再次输入密码"/>
              </div>
              <div v-if="!dialogSwitch" class="form-item form-item-tip">
                没有账户?<span @click="switchToRegister">注册</span>一个吧
              </div>

              <div v-if="!dialogSwitch" class="form-item form-item-btn">
                <button @click="handleToLogin">{{ logining ? '登录中' : '登录'}}</button>
                <button class="btn btn-back" @click="dialogVisible = false">取消</button>
              </div>

              <div v-if="dialogSwitch" class="form-item form-item-btn">
                <button @click="handleToRegister">{{ registering ? '注册中' : '注册'}}</button>
                <button class="btn btn-back" @click="switchToLogin">返回</button>
              </div>
            </div>
          </div>
        </transition>
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

  .toolbar {
    float: right;

    padding-top: 10px;
    margin-right: 16px;

    span.btn {
      margin-left: 8px;
      display: inline-block;
      height: 26px;
      line-height: 26px;
      width: 48px;
      border: 1px solid #20b2aa;
      text-align: center;
      font-size: 14px;
      color: #20b2aa;
      cursor: pointer;
      transition: background-color 0.3s, color 0.3s;
      &:hover {
        background-color: #20b2aa;
        color: white;
      }
    }
  }

  .dialog {
    z-index: 999;
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.8);

    .form {
      width: 80%;
      max-width: 280px;

      padding: 16px;

      margin-top: 120px;

      margin-left: 50%;
      transform: translateX(-50%);

      background-color: white;
      border-radius: 4px;

      .form-item {
        margin-top: 8px;

        &.form-item-tip {
          text-align: right;
          font-size: 13px;

          span {
            margin: 0 6px;
            color: #20b2aa;
            cursor: pointer;
          }
        }

        &.form-item-btn {
          margin-top: 16px;
          button {
            border-radius: 2px;
            margin-top: 4px;
            width: 100%;
            border: none;
            outline: none;
            padding: 6px;
            background-color: #20b2aa;
            color: white;
            cursor: pointer;
            font-size: 14px;
            transition: background-color 0.3s;
            &:hover {
              background-color: #10a199;
            }

            &.btn-back {
              background-color: #aaa;
              &:hover,
              &:active {
                background-color: #999;
              }
            }
          }
        }

        &.form-item-input {
          position: relative;

          span {
            position: absolute;
            left: 4px;
            top: 4px;
          }

          input {
            height: 32px;
            border: none;
            outline: none;
            border-bottom: 1px solid #ccc;
            width: 100%;
            padding-left: 32px;

            transition: border 0.3s;
            font-size: 14px;
            box-sizing: border-box;

            &:focus,
            &:hover {
              border-bottom: 1px solid #20b2aa;
            }
          }

          &.warn {
            input {
              border-color: red;
            }
          }
        }
      }
    }
  }

  .icon {
    display: block;
    width: 24px;
    height: 24px;

    background-repeat: no-repeat;
    background-size: cover;
    background-position: 50% 50%;
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
import SimpleFooter from '@/components/SimpleFooter';

import {
  USER_ACTION_LOGIN,
  USER_ACTION_CHECK,
  USER_ACTION_REGISTER,
  USER_ACTION_LOGOUT
} from '@/store/types';

import logo from '@/assets/imgs/logo.png';
import usernameIcon from '@/assets/icons/username.png';
import passwordIcon from '@/assets/icons/password.png';

export default {
  data() {
    return {
      logo,
      usernameIcon,
      passwordIcon,
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
      ],
      dialogVisible: false,
      dialogSwitch: false,

      logining: false,
      registering: false,

      form: {
        username: '',
        password: '',
        password2: ''
      },
      formValid: {
        username: true,
        password: true,
        password2: true
      }
    };
  },
  methods: {
    handleToLogin() {
      const ctx = this;

      if (ctx.logining) {
        return;
      }

      const { username, password } = ctx.form;
      if (!username || !password) {
        ctx.formValid = {
          username: !!username,
          password: !!password,
          password2: true
        };

        ctx.$message({
          type: 'warning',
          message: '用户名和密码不能为空'
        });
        return;
      } else {
        ctx.formValid = {
          username: true,
          password: true,
          password2: true
        };
      }

      ctx.logining = true;

      ctx.$store
        .dispatch(USER_ACTION_LOGIN, {
          username,
          password
        })
        .then(res => {
          ctx.logining = false;
          ctx.dialogVisible = false;

          ctx.$message({
            type: 'success',
            message: '登陆成功'
          });
        })
        .catch(err => {
          ctx.logining = false;
        });
    },
    handleToRegister() {
      const ctx = this;

      if (ctx.registering) {
        return;
      }

      const { username, password, password2 } = ctx.form;
      if (!username || !password || !password2 || password !== password2) {
        ctx.formValid = {
          username: !!username,
          password: !!password,
          password2: password === password2
        };

        ctx.$message({
          type: 'warning',
          message: '请填写所有内容'
        });
        return;
      } else {
        ctx.formValid = {
          username: true,
          password: true,
          password2: true
        };
      }

      ctx.registering = true;

      ctx.$store
        .dispatch(USER_ACTION_REGISTER, {
          username,
          password
        })
        .then(res => {
          ctx.registering = false;
          ctx.dialogVisible = false;

          ctx.$message({
            type: 'success',
            message: '注册成功'
          });
        })
        .catch(err => {
          ctx.registering = false;
        });
    },
    handleToLogout() {
      this.$store.dispatch(USER_ACTION_LOGOUT);
    },
    switchToLogin() {
      this.form = {
        username: '',
        password: '',
        password2: ''
      };
      this.formValid = {
        username: true,
        password: true,
        password2: true
      };

      this.dialogVisible = true;
      this.dialogSwitch = false;
    },
    switchToRegister() {
      this.formValid = {
        username: true,
        password: true,
        password2: true
      };
      this.form = {
        username: '',
        password: '',
        password2: ''
      };
      this.dialogSwitch = true;
    }
  },
  computed: {
    check() {
      return this.$store.getters.check;
    },
    username() {
      return this.$store.getters.username;
    }
  },
  components: {
    NavHeader,
    SimpleFooter
  },
  mounted() {
    const ctx = this;

    ctx.$store
      .dispatch(USER_ACTION_CHECK)
      .then(res => {})
      .catch(err => {});
  }
};
</script>
