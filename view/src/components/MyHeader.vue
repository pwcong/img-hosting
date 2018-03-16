<template>
    <div class="header">
        <a class="header-brand" href="/">
            {{ brand }}
        </a>
        <router-link 
            class="header-menu"
            v-for="menu in menuList"
            v-if="menu.active"
            :key="menu.path" 
            :to="menu.path">
            {{ menu.text }}
        </router-link>
        <div class="header-user">

            <div 
                class="header-user-item header-user-item-btn" 
                v-if="!logined"
                @click="onSignInClickListener">
                <span class="fa fa-lg fa-sign-in"></span>
            </div>

            <div class="header-user-item" v-if="logined">
                {{ "Hi~ o(*￣▽￣*)ブ " +  uid }}
            </div>
            <div 
                class="header-user-item header-user-item-btn" 
                v-if="logined"
                @click="onSignOutClickListener">
                <span class="fa fa-lg fa-sign-out"></span>
            </div>

        </div>
    </div>
</template>
<style>
.header {
  width: 100%;
  height: 56px;

  background-color: #2c3e50;

  box-sizing: border-box;

  display: flex;
  padding-left: 24px;
  padding-right: 24px;

  position: relative;
}

.header-brand {
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

.header-menu {
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

.header-menu:hover {
  color: lightseagreen;
}

.router-link-exact-active {
  background-color: lightseagreen;
}

.router-link-exact-active:hover {
  color: #2c3e50;
}

.header-user {
  height: 100%;

  position: absolute;
  right: 32px;

  display: flex;
  align-items: center;
}

.header-user-item {
  font-size: 14px;

  color: white;

  box-sizing: border-box;
  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  -ms-box-sizing: border-box;

  padding: 8px;

  display: flex;
  align-items: center;
}
.header-user-item-btn {
  cursor: pointer;

  user-select: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;

  transition: color 0.3s;
}

.header-user-item-btn:hover {
  color: lightseagreen;
}

.header-user-item-btn:active {
  color: #10a199;
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
    },
    logined: {
      type: Boolean,
      require: true
    },
    uid: {
      type: String,
      require: true
    },
    onSignInClickListener: {
      type: Function,
      default: function(e) {
        console.log('sign in');
      }
    },
    onSignOutClickListener: {
      type: Function,
      default: function(e) {
        console.log('sign out');
      }
    }
  },
  data() {
    return {};
  },
  computed: {},
  methods: {},
  components: {}
};
</script>