<template>
    <div class="content content-main">
        <div class="upload-manager">
            <div 
                class="upload-manager-main"
                @dragenter.prevent="onUploadManagerDragEnter"
                @dragover.prevent="onUploadManagerDragOver"
                @dragleave.prevent="onUploadManagerDragLeave"
                @drop.prevent="onUploadManagerDrop"
                :style="{
                    backgroundColor: uploadManagerBackgroundColor
                }"
                >
                <image-item 
                    v-for="img in imgList" 
                    :key="img.flag"
                    :flag="img.flag"
                    :img="img.img"
                    :uploading="img.uploading"
                    :uploaded="img.uploaded"
                    :progress="img.progress"
                    :onRemoveButtonClick="onImageItemRemove"
                    :onUploadButtonClick="onImageItemUpload"
                    :img-url="img.url"
                    :onShowImageURL="onImageItemInfo"
                    />
            </div>
            <div v-if="imgListLength <= 0" class="upload-manager-tips">
                <h1>Drag & drop image here</h1>
            </div>
        </div>
        <div class="upload-tools">
            <div class="upload-tools-msg">
                {{ msg }}
            </div>
            <div class="upload-tools-btns">

                <div 
                    @click="copyMsg"
                    class="upload-tools-btn upload-tools-btn-copy" 
                    v-if="msg != ''">
                    复制
                </div>
                <div 
                    @click="onUploadAllImage"
                    class="upload-tools-btn upload-tools-btn-upload" 
                    v-if="canUploadAll">
                    上传
                </div>
                <div class="upload-tools-btn upload-tools-btn-browse">
                    <input type="file" multiple="multipe" @change="onFilesSelect"/>
                    浏览
                </div>

            </div>
        </div>
        
    </div>
</template>
<style>

    .content-main{

        box-sizing: border-box;
        -moz-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -ms-box-sizing: border-box;

        padding: 16px; 

        flex: 1;

        display: flex;
        flex-direction: column;

    }

    .upload-manager{
        width: 100%;
        flex: 1;

        border: 1px #ccc solid;
        border-radius: 4px;

        display: flex;
        flex-direction: column;

        position: relative;


        box-sizing: border-box;
        -moz-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -ms-box-sizing: border-box;

        padding: 16px; 

    }

    .upload-manager-tips{

        color: #ccc;

        text-align: center;

        position: absolute;
        left: 0;
        top: 0;

        width: 100%;
        height: 100%;

        display: flex;
        align-items: center;
        justify-content: center;

        box-sizing: border-box;
        -moz-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -ms-box-sizing: border-box;

        user-select: none;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;

        padding: 32px; 
    }


    .upload-manager-main{

        z-index: 100;

        background-color: rgba(1, 1, 1, 0);

        transition: background-color 0.3s;

        width: 100%;
        flex: 1;

        box-sizing: border-box;
        -moz-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -ms-box-sizing: border-box;

        padding: 8px;

        border-radius: 4px;
        border: 1px #2c3e50 dashed;

        display: flex;
        flex-flow: row wrap;

        overflow: auto;
    }
    
    .upload-manager-main::-webkit-scrollbar{

        width: 4px;
        background-color: #ccc;

    }

    .upload-manager-main::-webkit-scrollbar-thumb{
    
        background-color: lightseagreen;
        
    }
    .upload-tools{

        width: 100%;
        height: 40px;
        min-height: 40px;

        margin-top: 8px;

        display: flex;
        flex-flow: row nowrap;
        align-items: center;
    }

    .upload-tools-msg{

        border: 1px #ccc solid;
        border-top-left-radius: 4px;
        border-bottom-left-radius: 4px;

        height: 100%;
        flex: 1;

        display: flex;
        align-items: center;

        box-sizing: border-box;
        -ms-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -moz-box-sizing: border-box;

        padding-left: 8px;
        
        overflow: auto;

        transition: border 0.3s;

    }

    .upload-tools-msg:hover{
        border: 1px #2c3e50 solid;
    }

    .upload-tools-btns{
        height: 100%;

        display: flex;

    }

    .upload-tools-btn{

        box-sizing: border-box;
        -ms-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -moz-box-sizing: border-box;

        padding-left: 16px;
        padding-right: 16px;

        display: flex;
        align-items: center;

        height: 100%;

        background-color: #2c3e50;

        user-select: none;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;

        color: white;

        cursor: pointer;

        transition: background-color 0.3s;

    }
    .upload-tools-btn-browse{
        border-top-right-radius: 4px;
        border-bottom-right-radius: 4px;
        position: relative;
    }

    .upload-tools-btn-browse:hover{
        background-color: #3d4f61;
    }
    .upload-tools-btn-browse:active{
        background-color: #1b2d40;
    }

    .upload-tools-btn-browse input{
        opacity: 0;
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
    }


    .upload-tools-btn-upload{
        background-color: #10a199;
    }


    .upload-tools-btn-upload:hover{
        background-color: lightseagreen;
    }
    .upload-tools-btn-upload:active{
        background-color: #009088;
    }

    .upload-tools-btn-copy{
        background-color: #10a199;
    }


    .upload-tools-btn-copy:hover{
        background-color: lightseagreen;
    }
    .upload-tools-btn-copy:active{
        background-color: #009088;
    }

</style>
<script>

    import * as types from '../store/types.js';

    import { mapGetters } from 'vuex';

    import ImageItem from '../component/ImageItem.vue';

    import api from '../api.js';

    export default {
        data(){
            return {
                uploadManagerBackgroundColor: 'rgba(1, 1, 1, 0)',
                msg: ''
            }
        },
        computed: mapGetters([
            'imgListLength',
            'imgList',
            'canUploadAll'
        ]),
        methods: {
            onImageItemRemove(e){

                this.$store.commit(types.MUTATION_IMG_REMOVEIMG, {
                    flag: e.target.getAttribute("flag")
                });

            },
            onImageItemUpload(e){

                let ctx = this;

                ctx.$store.dispatch(types.ACTION_IMG_TOUPLOADIMG, {
                    flag: e.target.getAttribute("flag"),
                    uid: ctx.$store.state.user.uid
                });

            },
            onUploadAllImage(e){

                let ctx = this;
                ctx.$store.dispatch(types.ACTION_IMG_TOUPLOADALLIMG, {
                    uid: ctx.$store.state.user.uid
                });
            },
            onUploadManagerDragEnter(e){

                this.$data.uploadManagerBackgroundColor = "rgba(1, 1, 1, 0.5)";
            },
            onUploadManagerDragOver(e){

                this.$data.uploadManagerBackgroundColor = "rgba(1, 1, 1, 0.5)";
            },
            onUploadManagerDragLeave(e){

                this.$data.uploadManagerBackgroundColor = "rgba(1, 1, 1, 0)";
            },
            onUploadManagerDrop(e){

                let ctx = this;

                ctx.$data.uploadManagerBackgroundColor = "rgba(1, 1, 1, 0)";

                let files =  e.dataTransfer.files;

                for(let i = 0; i < files.length; i++){
                    ctx.$store.commit(types.MUTATION_IMG_PUSHIMG, {
                        file: files[i]
                    });
                }
                

            },
            onFilesSelect(e){

                let ctx = this;

                let files =  e.target.files;

                for(let i = 0; i < files.length; i++){
                    ctx.$store.commit(types.MUTATION_IMG_PUSHIMG, {
                        file: files[i]
                    });
                }

            },
            onImageItemInfo(e){
                this.$data.msg = e.target.getAttribute("img-url");
            },
            copyMsg(e){

                window.prompt("请 Ctrl + C 拷贝至剪贴板", this.$data.msg);

            }
        },
        components: {
            ImageItem
        }
        
    }
    
</script>