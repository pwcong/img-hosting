<template>
    <div class="container">
        <h1>
            {{content}}
        </h1>
    </div>
</template>
<style>
    .container{
        width: 100%;
        display: flex;
        justify-content: center;
    }
</style>
<script>

    export default {
        data(){
            return {
                content: 'Hello IMG Hosting'
            }
        }
    }
</script>