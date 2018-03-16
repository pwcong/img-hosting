<template>
    <div :class="{ 'tips': true, 'tips-active': showTips }">
        <div 
            :class="{
                'tips-content': true,
                'tips-content-info': tipsType == 0,
                'tips-content-check': tipsType == 1,
                'tips-content-warning': tipsType == 2,
            }">
            <span 
                :class="{
                    'fa': true,
                    'fa-info': tipsType == 0,
                    'fa-check': tipsType == 1,
                    'fa-warning': tipsType == 2
                }">
            </span>
            {{ tipsContent }}
        </div>
    </div>
</template>
<style>
.tips {
  z-index: 999;

  position: fixed;
  left: 0;
  bottom: -48px;
  width: 100%;
  height: 48px;
  display: flex;
  justify-content: center;

  transition: bottom 0.3s;
}

.tips-active {
  bottom: 0;
}

.tips-content {
  height: 100%;

  color: white;

  box-sizing: border-box;
  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  -ms-box-sizing: border-box;

  border-top-left-radius: 4px;
  border-top-right-radius: 4px;

  display: flex;
  align-items: center;

  padding-left: 32px;
  padding-right: 32px;

  background-color: #2c3e50;
}
.tips-content-info {
  background-color: #2c3e50;
}
.tips-content-check {
  background-color: lightseagreen;
}
.tips-content-warning {
  background-color: orangered;
}

.tips-content span {
  margin-right: 8px;
}
</style>
<script>
export default {
  props: {
    showTips: {
      type: Boolean,
      default: false
    },
    tipsContent: {
      type: String,
      default: 'tips'
    },
    tipsType: {
      type: Number,
      default: 0
    }
  }
};
</script>


