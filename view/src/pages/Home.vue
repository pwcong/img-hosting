<template>
    <div class="home">
        <button type="button" @click="handleSendRequest">Send Request</button>
    </div>
</template>
<style lang="scss">
.home {

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

import sampleAPI from '../network/api/sample';

export default {

    methods: {
        handleSendRequest(e) {

            sampleAPI.test()
                .then(res => {
                    alert('send request successfully !');
                })
                .catch(err => {
                    alert('send request fail !');
                });

        }
    }

}
</script>
