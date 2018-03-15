<template>
    <div class="counter">
        <div class="counts">
            <div>Counts: </div>
            {{counts}}
            <div>Double Counts: </div>
            {{doubleCounts}}
        </div>

        <div class="tools">
            <button type="button" @click="handleCounterPlus">Plus</button>
            <button type="button" @click="handleCounterPlusAsync">Plus Async</button>
        </div>
    </div>
</template>
<style lang="scss">
.counts {
    margin-bottom: 16px;
    font-size: 18px;
    color: #41B883;

    div {
        color: #333;
        display: inline-block;
        margin-left: 16px;
    }
}

.tools {
    button {
        background-color: white;
        padding: 8px 16px;
        color: #41B883;
        border: 1px solid #41B883;
        outline: none;
        cursor: pointer;

        &:active {
            color: white;
            background-color: #41B883;
        }
    }
}
</style>
<script>

import types from '../store/types';

export default {

    data() {
        return {
        }
    },
    computed: {
        counts() {
            return this.$store.state.sample.counts;
        },
        doubleCounts() {
            return this.$store.getters.doubleCounts;
        }
    },
    methods: {
        handleCounterPlus(e) {
            this.$store.dispatch(types.SAMPLE_ACTIONS_PLUS);
        },
        handleCounterPlusAsync(e) {
            this.$store.dispatch(types.SAMPLE_ACTIONS_PLUS_ASYNC, {
                time: 1000
            });
        }
    }

}
</script>
