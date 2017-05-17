<template>
    <div class="content content-own">
        <own-item 
            v-for="item in ownList" 
            :key="item.url" 
            :url="item.url"
            :filename="item.filename"
            :year="item.year"
            :month="item.month"
            :day="item.day"
            :storename="item.storename"
            :onRemoveItemClickListener="handleRemoveOwnItem"
            />
        
        <div class="content-own-tips" v-if="ownListLength < 1">
            <h1>Your have no own picture</h1>
        </div>
        
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

        overflow: auto;

        position: relative;
    }

    .content-own::-webkit-scrollbar{
        width: 4px;
        background-color: #ccc;

    }

    .content-own::-webkit-scrollbar-thumb{
    
        background-color: lightseagreen;
        
    }


    .content-own-tips{
        position: absolute;
        left: 0;
        top: 0;
        height: 100%;
        width: 100%;

        display: flex;
        align-items: center;
        justify-content: center;

        color: #ccc;
    }

</style>
<script>

    import { mapGetters } from 'vuex';

    import * as types from '../store/types.js';

    import OwnItem from '../component/OwnItem.vue';

    export default {

        computed: mapGetters([
            'ownList',
            'ownListLength'
        ]),
        methods: {
            handleRemoveOwnItem(year, month, day, storename){
                console.log(year,month,day,storename);
                this.$store.dispatch(types.ACTION_IMG_REMOVEOWN, {
                    token: localStorage.token,
                    year,
                    month,
                    day,
                    storename
                });
            }
        },
        components: {
            OwnItem
        },
        mounted: function() {

            let ctx = this;

            ctx.$store.dispatch(types.ACTION_IMG_GETLIST, {
                token: localStorage.token
            });
        }

    }
</script>

