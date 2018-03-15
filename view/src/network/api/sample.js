import request from '../request';

import api from './api';

function test() {

    return request(
        api.sample.test.url(),
        api.sample.test.method,
        {},
        {}
    );

}

export default {
    test
}