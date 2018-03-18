<template>
    <div class="home">

      <div class="upload-box">
        <div class="upload-box-wrap"
          @dragenter.prevent="handleUploadDragEnter"
          @dragover.prevent="handleUploadDragOver"
          @dragleave.prevent="handleUploadDragLeave"
          @drop.prevent="handleUploadDrop"
          >

          <div class="upload-box-tip" v-if="uploadList.length <= 0">
            Drag &amp; drop here
          </div>

          <el-card 
            v-for="(img, idx) in uploadList" 
            :key="'img-item-' + idx"
            :body-style="{ padding: '0px' }">
            <div class="thumbnail" :style="{
              backgroundImage: 'url(' + img.img + ')'
            }" @click="handlePreviewImage(img)"></div>
            <div style="padding: 8px;">
              <div class="bottom">
                <i v-if="img.uploading" class="el-icon el-icon-loading"></i>
                <i v-if="!img.uploading && !img.success" class="el-icon el-icon-close" style="color: red;"></i>
                <i v-if="!img.uploading && img.uploaded && img.success" class="el-icon el-icon-check" style="color: #20b2aa;"></i>
                <i v-if="img.uploaded" class="el-icon el-icon-info" @click="handleViewImage(img)"></i>
                <i v-if="!img.uploading && !img.uploaded" class="el-icon el-icon-upload" @click="handleUploadImage(img)"></i>
                <i v-if="!img.uploading" class="el-icon el-icon-delete" @click="handleRemoveImage(img)"></i>
              </div>
            </div>
            <div class="progress">
              <div class="progress-wrap" :style="{
                width: img.progress + '%'
              }"></div>
            </div>
          </el-card>

        </div>
      </div>
      <div class="upload-toolbar">
        <input id="urlResult" class="upload-toolbar-input" type="text" v-model="previewUrl">

        <button id="copyBtn" data-clipboard-target="#urlResult" class="upload-toolbar-btn btn-copy" type="button" @click="handleCopyResultUrl">复制</button>
        <button v-if="canUploadAll" class="upload-toolbar-btn btn-upload" type="button" @click="handleUploadAllImages">上传</button>

        <label class="upload-toolbar-btn btn-browser">
          <input type="file" multiple="multipe" accept="image/*" @change="handleSelectImages"/>
          浏览
        </label>

      </div>

      <el-dialog title="图片预览" :visible.sync="previewDialogVisible">
        <img :src="previewImage" style="max-width: 100%"/>
      </el-dialog>

    </div>
</template>
<style lang="scss">
.home {
  width: 100%;
  box-sizing: border-box;
  padding: 32px;

  .upload-box {
    border: 1px solid #ccc;
    border-radius: 2px;

    padding: 16px;

    .upload-box-wrap {
      border: 1px dashed #888;
      height: 400px;
      border-radius: 2px;
      overflow: auto;

      box-sizing: border-box;
      padding: 8px;

      transition: border 0.3s;

      &:hover {
        border-color: #20b2aa;
      }

      .upload-box-tip {
        padding-top: 180px;
        font-size: 32px;
        color: #ccc;
        text-align: center;
      }

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
          color: #888;

          text-align: center;

          .el-icon {
            font-size: 18px;
            cursor: pointer;
            margin: 0 4px;
            transition: color 0.3s;
            &.el-icon-delete:hover {
              color: red;
            }
            &.el-icon-upload:hover {
              color: #20b2aa;
            }
            &.el-icon-info:hover {
              color: #666;
            }
          }
        }

        .progress {
          position: relative;
          .progress-wrap {
            width: 0;
            height: 2px;
            background-color: #20b2aa;
          }
        }
      }

      &::-webkit-scrollbar {
        width: 4px;
        background-color: #ccc;
      }

      &::-webkit-scrollbar-thumb {
        background-color: #20b2aa;
      }
    }
  }

  .upload-toolbar {
    margin-top: 8px;

    display: flex;
    align-items: center;
    flex-flow: row nowrap;

    .upload-toolbar-input {
      height: 38px;
      flex: 1;
      padding: 0 8px;
      border: 1px solid #aaa;
    }

    .upload-toolbar-btn {
      display: flex;
      align-items: center;
      justify-content: center;
      height: 40px;
      position: relative;
      outline: none;
      width: 80px;
      border: none;
      cursor: pointer;
      font-size: 13px;
      color: white;
      transition: background-color 0.3s;

      input {
        opacity: 0;
        visibility: hidden;
        position: absolute;
        left: 0;
        top: 0;
        height: 100%;
        width: 100%;
      }

      &.btn-upload {
        background-color: #20b2aa;

        &:hover {
          background-color: #10a199;
        }
      }
      &.btn-browser {
        background-color: #2c3e50;
        &:hover {
          background-color: #1b2d40;
        }
      }
      &.btn-copy {
        background-color: #aaa;
        &:hover {
          background-color: #999;
        }
      }
    }
  }
}
</style>
<script>
import chance from 'chance';

import clipboard from 'clipboard';

import { IMG_ACTION_UPLOAD, IMG_MUTATION_PUSH_IMG, IMG_MUTATION_POP_IMG } from '@/store/types';

export default {
  data() {
    return {
      previewUrl: '',
      previewDialogVisible: false,
      previewImage: ''
    };
  },
  methods: {
    handleSelectImages(e) {
      const ctx = this;

      const files = e.target.files;

      for (let i = 0; i < files.length; i++) {
        ctx.$store.commit(IMG_MUTATION_PUSH_IMG, {
          symbol: chance().string(),
          file: files[i],
          img: window.URL.createObjectURL(files[i]),
          uploading: false,
          uploaded: false,
          success: true,
          url: ''
        });
      }
    },
    handlePreviewImage(img) {
      this.previewDialogVisible = true;
      this.previewImage = img.img;
    },
    handleRemoveImage(img) {
      const ctx = this;
      ctx.$store.commit(IMG_MUTATION_POP_IMG, img);
    },
    handleViewImage(img) {
      this.previewUrl = img.url;
    },
    handleUploadImage(img) {
      this.$store
        .dispatch(IMG_ACTION_UPLOAD, img)
        .then(res => {})
        .catch(err => {});
    },
    handleUploadDragEnter(e) {},
    handleUploadDragOver(e) {},
    handleUploadDragLeave(e) {},
    handleUploadDrop(e) {
      const ctx = this;
      const files = e.dataTransfer.files;

      for (let i = 0; i < files.length; i++) {
        ctx.$store.commit(IMG_MUTATION_PUSH_IMG, {
          symbol: chance().string(),
          file: files[i],
          img: window.URL.createObjectURL(files[i]),
          uploading: false,
          uploaded: false,
          success: true,
          url: ''
        });
      }
    },
    handleUploadAllImages() {
      this.$store
        .dispatch(IMG_ACTION_UPLOAD)
        .then(res => {})
        .catch(err => {});
    },
    handleCopyResultUrl() {
      if (!this.previewUrl) {
        return;
      }

      this.$message({
        type: 'success',
        message: '图片链接复制成功'
      });
    }
  },
  computed: {
    uploadList() {
      return this.$store.getters.uploadList;
    },
    canUploadAll() {
      const uploadList = this.$store.getters.uploadList || [];
      const l = uploadList.length;
      for (let i = 0; i < l; i++) {
        if (!uploadList[i].uploaded && !uploadList[i].uploading) {
          return true;
        }
      }

      return false;
    }
  },
  mounted() {
    new clipboard('#copyBtn');
  }
};
</script>
