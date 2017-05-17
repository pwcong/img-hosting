<template>
    <div class="router-view-container">
        <router-view></router-view>
    </div>
</template>
<style>

    .router-view-container{
        width: 100%;
        flex: 1;

        box-sizing: border-box;
        -ms-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -moz-box-sizing: border-box;

        padding: 16px 32px;

        overflow: auto;

        position: relative;

        display: flex;
        flex-direction: column;

    }

    .content{
        width: 100%;
        color: #2c3e50;
    }

    .router-view-container::-webkit-scrollbar{

        width: 4px;
        background-color: #ccc;

    }

    .router-view-container::-webkit-scrollbar-thumb{
    
        background-color: lightseagreen;
        
    }

</style>
<script>

    export default {
        
        
    }
    
</script>