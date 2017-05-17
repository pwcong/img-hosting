<template>
    <div class="content content-own">
        <own-item />
    </div>
</template>
<style>
    .content-own{
        border: 1px #ccc solid;

        border-radius: 4px;

        box-sizing: border-box;
        -ms-box-sizing: border-box;
        -webkit-box-sizing: border-box;
        -moz-box-sizing: border-box;

        padding: 8px;

        flex: 1;

        display: flex;
        flex-flow: row wrap;

        overflow: auto;
    }

</style>
<script>

    import OwnItem from '../component/OwnItem.vue';

    export default {
    
        components: {
            OwnItem
        }

    }
</script>

