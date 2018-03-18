<template>
  <div class="list">
    <div v-if="!check" class="tip">
      <img :src="pleaseImg" />
      <div>
        您还未登录，请先登录
        <span>.</span>
        <span>.</span>
        <span>.</span>
      </div>
    </div>
    <div class="list-wrap" v-if="check">
      <div v-if="imgList.length <= 0" class="tip" style="font-size: 16px; color: #aaa; font-weight: normal;">
        没有啦♪(^∇^*)
        <span>.</span>
        <span>.</span>
        <span>.</span>
      </div>
      <el-card 
        v-for="(img, idx) in imgList" 
        :key="'img-item-' + idx"
        :body-style="{ padding: '0px' }">
        <div class="thumbnail" :style="{
          backgroundImage: 'url(' + img.url + ')'
        }" @click="handlePreviewImage(img)"></div>
        <div style="padding: 14px;">
          <span class="label">{{img.filename}}</span>
          <div class="bottom">
            <span>{{img.created_at}}</span>
            <i class="el-icon el-icon-info" @click="handleViewImage(img)"></i>
            <i class="el-icon el-icon-delete" @click="handleRemoveImage(img)"></i>
          </div>
        </div>
      </el-card>

    </div>
    <el-pagination
      v-if="check"
      layout="prev, pager, next"
      @current-change="handleCurrentPageChange"
      :page-size="pageSize"
      :total="totaSize">
    </el-pagination>
    
    <el-dialog title="图片预览" :visible.sync="previewDialogVisible">
      <img :src="previewImage" style="max-width: 100%"/>
    </el-dialog>
  </div>
</template>
<style lang="scss">
.list {
  position: relative;
  width: 100%;
  .tip {
    padding-top: 80px;
    text-align: center;
    font-size: 20px;
    font-weight: bold;
    img {
      max-width: 200px;
    }

    span {
      position: relative;
      animation: dotanim 1.2s infinite;

      &:nth-child(2) {
        animation-delay: 0.3s;
      }
      &:nth-child(3) {
        animation-delay: 0.6s;
      }
    }
  }

  .list-wrap {
    min-height: 300px;
    padding: 16px;

    .el-card {
      display: inline-block;
      width: 180px;
      margin-right: 12px;
      .thumbnail {
        cursor: pointer;
        width: 100%;
        height: 160px;
        background-repeat: no-repeat;
        background-size: cover;
        background-position: center center;
      }

      .label {
        font-size: 15px;
        text-overflow: ellipsis;
        white-space: nowrap;
        overflow: hidden;
        display: block;
      }

      .bottom {
        margin-top: 16px;
        font-size: 11px;
        color: #888;

        .el-icon {
          margin-left: 8px;
          position: relative;
          top: -2px;
          font-size: 18px;
          float: right;
          cursor: pointer;

          &.el-icon-delete:hover {
            color: red;
          }
          &.el-icon-info:hover {
            color: #20b2aa;
          }
        }
      }
    }
  }

  .el-pagination {
    text-align: center;
  }
}

@keyframes dotanim {
  0% {
    top: 0;
  }

  50% {
    top: -8px;
  }

  100% {
    top: 0;
  }
}
</style>
<script>
import Cookies from 'js-cookie';
import moment from 'moment';

import { removeImage } from '@/network/api/img';

import { USER_ACTION_CHECK, IMG_ACTION_GETLIST } from '@/store/types';

import { BASE_API } from '@/const/config';
import pleaseImg from '@/assets/imgs/please.png';

const PAGE_SIZE = 20;

export default {
  data() {
    return {
      pleaseImg,
      totaSize: 0,
      pageSize: PAGE_SIZE,

      previewDialogVisible: false,
      previewImage: ''
    };
  },
  methods: {
    getList(currentPage) {
      const ctx = this;

      if (!ctx.check) {
        return;
      }

      ctx.$store
        .dispatch(IMG_ACTION_GETLIST, {
          pageSize: ctx.pageSize,
          pageNo: currentPage
        })
        .then(res => {
          ctx.totaSize = res.payload.total_size;
        })
        .catch(err => {});
    },
    handlePreviewImage(img) {
      this.previewDialogVisible = true;
      this.previewImage = img.url;
    },
    handleRemoveImage(img) {
      const ctx = this;

      ctx
        .$confirm('此操作将永久删除该图片, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
        .then(() => {
          removeImage(img.id)
            .then(res => {
              ctx.$message({
                type: 'success',
                message: '删除成功',
                duration: 1500
              });

              ctx.getList(1);
            })
            .then(err => {});
        })
        .catch(() => {
          ctx.$message({
            type: 'info',
            message: '已取消删除',
            duration: 1500
          });
        });
    },
    handleCurrentPageChange(currentPage) {
      this.getList(currentPage);
    },
    handleViewImage(img) {
      this.$alert(`<input value="${img.url}" style="width: 100%; border: none; border-bottom: 1px solid #ccc; height: 24px; outline: none;"/>`, '图片链接', {
        confirmButtonText: '确定',
        dangerouslyUseHTMLString: true
      });
    }
  },
  computed: {
    check() {
      return this.$store.getters.check;
    },
    imgList() {
      return this.$store.getters.imgList.map(img => ({
        ...img,
        url: BASE_API + img.url,
        created_at: moment(img.created_at).format('YYYY-MM-DD')
      }));
    }
  },
  mounted() {
    const ctx = this;

    if (Cookies.get('token')) {
      ctx.$store
        .dispatch(USER_ACTION_CHECK)
        .then(res => {
          ctx.getList(1);
        })
        .catch(err => {});
    }
  }
};
</script>
