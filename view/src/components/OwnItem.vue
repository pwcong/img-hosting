<template>
    <div class="own-item">
        <div class="own-item-preview">
            <img :src="url"/>
        </div>
        <div class="own-item-content">
            <div 
                class="own-item-content-item" 
                style="flex: 2; font-size: 18px;">
                <div style="margin-right: 32px;">
                    <span class="fa fa-file-image-o"></span>
                    {{ filename }}
                </div>
                <div>
                    <span class="fa fa-calendar"></span>
                    {{ year + '-' + month + '-' + day }}
                </div>
            </div>
            <div class="own-item-content-item"  style="font-size: 13px;">
                <span class="fa fa-link" style="margin-right: 16px;"></span>
                {{ url }}
            </div>
            
        </div>
        <div class="own-item-tools">
            <div class="own-item-tools-btn" @click="handleShowURL">
                <span class="fa fa-copy"></span>
            </div>
            <div class="own-item-tools-btn own-item-tools-btn-trash" @click="handleRemoveItem">
                <span class="fa fa-trash"></span>
            </div>
        </div>
    </div>
</template>
<style>
.own-item {
  margin-top: 4px;

  box-sizing: border-box;
  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  -ms-box-sizing: border-box;

  padding: 8px;

  border: 1px rgba(0, 0, 0, 0) solid;

  border-radius: 4px;

  width: 100%;
  height: 128px;

  display: flex;
  flex-flow: row nowrap;

  align-items: center;

  transition: border 0.3s;

  border: 1px #ccc solid;
}

.own-item:hover {
  border: 1px lightseagreen solid;
}

.own-item:hover .own-item-tools {
  border-left: 2px lightseagreen solid;
}

.own-item-preview {
  width: 156px;
  height: 100%;

  display: flex;
  justify-content: center;
}

.own-item-preview img {
  max-height: 100%;
  max-width: 100%;
}

.own-item-content {
  box-sizing: border-box;
  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  -ms-box-sizing: border-box;

  padding-left: 16px;
  padding-right: 16px;

  flex: 1;
  height: 100%;

  display: flex;
  flex-direction: column;
}

.own-item-content-item {
  flex: 1;
  width: 100%;

  display: flex;
  align-items: center;
}

.own-item-tools {
  height: 100%;
  width: 156px;

  border-left: 2px #ccc solid;

  transition: border-left 0.3s;

  font-size: 24px;

  display: flex;
  align-items: center;
  justify-content: space-around;
}

.own-item-tools-btn {
  box-sizing: border-box;
  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  -ms-box-sizing: border-box;

  padding: 8px;

  transition: color 0.3s;

  cursor: pointer;
}

.own-item-tools-btn:hover {
  color: lightseagreen;
}
.own-item-tools-btn:active {
  color: #10a199;
}
.own-item-tools-btn-trash:hover {
  color: orangered;
}

.own-item-tools-btn-trash:active {
  color: #ee3400;
}
</style>
<script>
export default {
  props: {
    url: {
      type: String,
      default: 'http://' + location.host + '/imgs/test.jpg'
    },
    filename: {
      type: String,
      default: 'test.jpg'
    },
    year: {
      type: String,
      default: '2017'
    },
    month: {
      type: String,
      default: '01'
    },
    day: {
      type: String,
      default: '01'
    },
    storename: {
      type: String,
      default: 'test.jpg'
    },
    onRemoveItemClickListener: {
      type: Function,
      default: function(year, month, day, storename) {
        console.log(year + ' ' + month + ' ' + day + ' ' + storename);
      }
    }
  },
  methods: {
    handleShowURL(e) {
      window.prompt('请 Ctrl + C 拷贝至剪贴板', this.$props.url);
    },

    handleRemoveItem(e) {
      let props = this.$props;

      props.onRemoveItemClickListener(props.year, props.month, props.day, props.storename);
    }
  }
};
</script>


