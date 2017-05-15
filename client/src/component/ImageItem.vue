<template>
    <div class="upload-manager-item">
        <div class="upload-manager-item-img">
            <img :src="img"/>
        </div>
        <div class="upload-manager-item-tools">
            <span 
                class="fa fa-upload" 
                v-if="!uploading && !uploaded"
                @click="onUploadButtonClick"
                :flag="flag"
                >
            </span>
            <span 
                class="fa fa-spinner" 
                v-if="uploading"
                >
            </span>
            <span 
                class="fa fa-check" 
                v-if="!uploading && uploaded"
                >
            </span>
            <span 
                :flag="flag"
                class="fa fa-trash" 
                v-if="!uploading"
                @click="onRemoveButtonClick"
                >
            </span>
        </div>
        <div class="upload-manager-item-progress" v-if="uploading">
            <div 
                class="upload-manager-item-progress-bar" 
                :style="{
                    width: progress + '%'
                }">
            </div>
        </div>
    </div>
</template>
<style>

    .upload-manager-item{

        background-color: white;

        margin: 4px;

        border: 1px #ccc solid;

        border-radius: 4px;

        height: 200px;

        box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -ms-box-sizing: border-box;
        -moz-box-sizing: border-box;

        padding: 8px;

        box-shadow: 0 0 0 0 ;

        display: flex;
        flex-direction: column;

        transition: box-shadow 0.1s;

    }

    .upload-manager-item:hover{
        box-shadow: 0 0 1px 0;
    }

    .upload-manager-item-img{
        height: 150px;

        display: flex;
        justify-content: center;

    }

    .upload-manager-item-img img{
        max-height: 100%;
    }

    .upload-manager-item-tools{
        flex: 1;

        display: flex;
        align-items: center;
        justify-content: space-around;

    }

    .upload-manager-item-tools span.fa{
        cursor: pointer;
        transition: color 0.3s;
    }
    .upload-manager-item-tools span.fa-spinner{
        animation: rotation 1s infinite linear;
        cursor: wait;
    }
    .upload-manager-item-tools span.fa-check{
        color: lightseagreen;
        cursor: default;
    }
    .upload-manager-item-tools span.fa-trash:hover{
        color: orangered;
    }
    .upload-manager-item-tools span.fa-upload:hover{
        color: lightseagreen;
    }

    .upload-manager-item-progress{
        height: 4px;

        position: relative;
    }

    .upload-manager-item-progress-bar{
        position: absolute;
        left: 0;
        top: 0;
        height: 100%;
        width: 0%;

        background-color: lightseagreen;
    }

    @keyframes rotation{
        from{
            transform: rotate(0deg);
        }
        to{
            transform: rotate(364deg);
        }
    }

</style>
<script>

    export default {

        props: {
            flag: {
                type: String,
                require: true
            },
            img: {
                type: String,
                require: true
            },
            uploading: {
                type: Boolean,
                require: true
            },
            uploaded: {
                type: Boolean,
                require: true
            },
            progress: {
                type: Number,
                require: true
            },
            onUploadButtonClick: {
                type: Function,
                default: function(e){
                    console.log("upload");
                }
            },
            onRemoveButtonClick: {
                type: Function,
                default: function(e){
                    console.log("remove");
                }
            }
        }
        
    }

</script>
