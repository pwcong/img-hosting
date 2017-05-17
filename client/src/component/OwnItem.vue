<template>
    <div class="own-item">
        <div class="own-item-preview">
            <img :src="url"/>
        </div>
        <div class="own-item-content">
            <div 
                class="own-item-content-item" 
                style="flex: 2; font-size: 20px">
                <div style="margin-right: 32px">
                    <span class="fa fa-file-image-o"></span>
                    {{ filename }}
                </div>
                <div>
                    <span class="fa fa-calendar"></span>
                    {{ filename }}
                </div>
            </div>
            <div class="own-item-content-item">

            </div>
            
        </div>
        <div class="own-item-tools">
            <span class="fa fa-trash"></span>
        </div>
    </div>
</template>
<style>
    .own-item{

        box-sizing: border-box;
        -moz-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -ms-box-sizing: border-box;

        padding: 8px;

        border: 1px #ccc solid;

        width: 100%;
        height: 128px;

        display: flex;
        flex-flow: row nowrap;

        align-items: center;

    }

    .own-item-preview{
        width: 200px;
        height: 100%;
        border-right: 2px #ccc solid;

        display: flex;
        justify-content: center;

    }

    .own-item-preview img{
        max-height: 100%;
        max-width: 100%;
    }

    .own-item-content{
        box-sizing: border-box;
        -moz-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -ms-box-sizing: border-box;

        padding: 8px;

        flex: 1;
        height: 100%;

        display: flex;
        flex-direction: column;

        
    }

    .own-item-content-item{
        border: 1px green solid;
        flex: 1;
        width: 100%;

        display: flex;
        align-items: center;
    }

    .own-item-tools{
        height: 100%;
        width: 156px;

        border: 1px green solid;

        display: flex;
        align-items: center;
        justify-content: center;
        
    }


</style>
<script>
    export default {
        props: {
            url: {
                type: String,
                default: '/imgs/test.jpg'
            },
            filename: {
                type: String,
                default: 'test.jpg'
            },
        }
    }
</script>


