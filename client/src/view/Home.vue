<template>
    <div class="content content-main">
        <div class="upload-manager">
            <div class="upload-manager-main">
                <image-item 
                    v-for="img in imgList" 
                    :key="img.flag"
                    :file="img.file"
                    :uploading="img.uploading"
                    :uploaded="img.uploaded"
                    :progress="img.progress"
                    />
            </div>
            <div v-if="imgListLength <= 0" class="upload-manager-tips">
                <h1>Drag & drop image here</h1>
            </div>
        </div>
        <div class="upload-tools">

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

        z-index: 999;

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
    }

    .upload-tools{

        margin-top: 8px;

        height: 50px;
        width: 100%;

        border: 1px red solid;
    }

</style>
<script>

    import { mapGetters } from 'vuex';

    import ImageItem from '../component/ImageItem.vue';

    import api from '../api.js';

    export default {
        data(){
            return {
                
            }
        },
        computed: mapGetters([
            'imgListLength',
            'imgList'
        ]),
        methods: {
            
        },
        components: {
            ImageItem
        }
        
    }
    
</script>