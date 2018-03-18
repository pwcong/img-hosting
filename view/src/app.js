import Vue from 'vue';

import { Message, MessageBox, Icon, Pagination, Card, Dialog } from 'element-ui';

Vue.use(Icon);
Vue.use(Pagination);
Vue.use(Card);
Vue.use(Dialog);

Vue.prototype.$message = Message;
Vue.prototype.$msgbox = MessageBox;
Vue.prototype.$alert = MessageBox.alert;
Vue.prototype.$confirm = MessageBox.confirm;
Vue.prototype.$prompt = MessageBox.prompt;

import 'normalize.css';
import './assets/css/github-markdown.css';

import store from './store';
import router from './routes';

import App from './pages/App';

new Vue({
  el: '#app',
  store,
  router,
  render: h => h(App)
});
