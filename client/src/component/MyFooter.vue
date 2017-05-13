<template>
    <div class="footer">
       {{ text }}
    </div>
</template>
<style>

    .footer{

        font-size: 12px;

        background-color: #ccc;

        box-sizing: border-box;
        -ms-box-sizing: border-box;
        -moz-box-sizing: border-box;
        -webkit-box-sizing: border-box;

        padding: 8px 32px;

        text-align: center;

    }

</style>
<script>

    export default {
        props: {
            text: {
                type: String,
                require: true
            }
        }
        
    }
    
</script>