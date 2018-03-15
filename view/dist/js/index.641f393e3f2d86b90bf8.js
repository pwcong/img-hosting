webpackJsonp([0],{106:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var u=n(97),r=function(t){return t&&t.__esModule?t:{default:t}}(u);e.default={data:function(){return{}},computed:{counts:function(){return this.$store.state.sample.counts},doubleCounts:function(){return this.$store.getters.doubleCounts}},methods:{handleCounterPlus:function(t){this.$store.dispatch(r.default.SAMPLE_ACTIONS_PLUS)},handleCounterPlusAsync:function(t){this.$store.dispatch(r.default.SAMPLE_ACTIONS_PLUS_ASYNC,{time:1e3})}}}},107:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var u=n(177),r=function(t){return t&&t.__esModule?t:{default:t}}(u);e.default={data:function(){return{logo:r.default}}}},142:function(t,e,n){"use strict";function u(t){return t&&t.__esModule?t:{default:t}}var r=n(51),o=u(r),s=n(145),a=u(s),i=n(147),c=u(i),l=n(175),f=u(l);new o.default({el:"#app",store:a.default,router:c.default,render:function(t){return t(f.default)}})},145:function(t,e,n){"use strict";function u(t){return t&&t.__esModule?t:{default:t}}Object.defineProperty(e,"__esModule",{value:!0});var r=n(51),o=u(r),s=n(96),a=u(s),i=n(146),c=u(i);o.default.use(a.default);var l=new a.default.Store({modules:{sample:c.default}});e.default=l},146:function(t,e,n){"use strict";function u(t,e,n){return e in t?Object.defineProperty(t,e,{value:n,enumerable:!0,configurable:!0,writable:!0}):t[e]=n,t}Object.defineProperty(e,"__esModule",{value:!0});var r,o=n(97),s=function(t){return t&&t.__esModule?t:{default:t}}(o),a={state:{counts:0},getters:{doubleCounts:function(t){return 2*t.counts}},mutations:u({},s.default.SAMPLE_MUTATIONS_PLUS,function(t,e){t.counts++}),actions:(r={},u(r,s.default.SAMPLE_ACTIONS_PLUS,function(t){var e=t.commit;t.state;e(s.default.SAMPLE_MUTATIONS_PLUS)}),u(r,s.default.SAMPLE_ACTIONS_PLUS_ASYNC,function(t,e){var n=t.dispatch;t.commit,t.state;return new Promise(function(t,u){setTimeout(function(){n(s.default.SAMPLE_ACTIONS_PLUS)},e.time)})}),r)};e.default=a},147:function(t,e,n){"use strict";function u(t){return t&&t.__esModule?t:{default:t}}Object.defineProperty(e,"__esModule",{value:!0});var r=n(51),o=u(r),s=n(98),a=u(s),i=n(148),c=u(i),l=n(172),f=u(l);o.default.use(a.default);var d=[{path:"/",component:c.default},{path:"/counter",component:f.default}],_=new a.default({routes:d});e.default=_},148:function(t,e,n){"use strict";function u(t){n(149)}Object.defineProperty(e,"__esModule",{value:!0});var r=n(99),o=n.n(r);for(var s in r)"default"!==s&&function(t){n.d(e,t,function(){return r[t]})}(s);var a=n(171),i=n(66),c=u,l=i(o.a,a.a,!1,c,null,null);e.default=l.exports},149:function(t,e){},150:function(t,e,n){"use strict";function u(t){return t&&t.__esModule?t:{default:t}}function r(){return(0,s.default)(i.default.sample.test.url(),i.default.sample.test.method,{},{})}Object.defineProperty(e,"__esModule",{value:!0});var o=n(151),s=u(o),a=n(170),i=u(a);e.default={test:r}},151:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=function(t,e,n,u){return new Promise(function(s,a){(0,r.default)({url:t,method:e,data:n,headers:Object.assign({},o,u)}).then(function(t){200===t.status?s(t):a(t)}).catch(function(t){a(t)})})};var u=n(100),r=function(t){return t&&t.__esModule?t:{default:t}}(u),o={}},170:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var u={sample:{test:{url:function(){return"#"},method:"GET"}}};e.default=u},171:function(t,e,n){"use strict";var u=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"home"},[n("button",{attrs:{type:"button"},on:{click:t.handleSendRequest}},[t._v("Send Request")])])},r=[],o={render:u,staticRenderFns:r};e.a=o},172:function(t,e,n){"use strict";function u(t){n(173)}Object.defineProperty(e,"__esModule",{value:!0});var r=n(106),o=n.n(r);for(var s in r)"default"!==s&&function(t){n.d(e,t,function(){return r[t]})}(s);var a=n(174),i=n(66),c=u,l=i(o.a,a.a,!1,c,null,null);e.default=l.exports},173:function(t,e){},174:function(t,e,n){"use strict";var u=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"counter"},[n("div",{staticClass:"counts"},[n("div",[t._v("Counts: ")]),t._v("\n        "+t._s(t.counts)+"\n        "),n("div",[t._v("Double Counts: ")]),t._v("\n        "+t._s(t.doubleCounts)+"\n    ")]),t._v(" "),n("div",{staticClass:"tools"},[n("button",{attrs:{type:"button"},on:{click:t.handleCounterPlus}},[t._v("Plus")]),t._v(" "),n("button",{attrs:{type:"button"},on:{click:t.handleCounterPlusAsync}},[t._v("Plus Async")])])])},r=[],o={render:u,staticRenderFns:r};e.a=o},175:function(t,e,n){"use strict";function u(t){n(176)}Object.defineProperty(e,"__esModule",{value:!0});var r=n(107),o=n.n(r);for(var s in r)"default"!==s&&function(t){n.d(e,t,function(){return r[t]})}(s);var a=n(178),i=n(66),c=u,l=i(o.a,a.a,!1,c,null,null);e.default=l.exports},176:function(t,e){},177:function(t,e,n){t.exports=n.p+"imgs/vue.png?161bdff073e8e30e52d59f99871cb5b7"},178:function(t,e,n){"use strict";var u=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"container"},[n("img",{attrs:{src:t.logo,alt:"Logo"}}),t._v(" "),n("div",{staticClass:"navs"},[n("router-link",{staticClass:"nav",attrs:{to:"/"}},[t._v("Home")]),t._v(" "),n("router-link",{staticClass:"nav",attrs:{to:"/counter"}},[t._v("Counter")])],1),t._v(" "),n("div",{staticClass:"content"},[n("router-view")],1)])},r=[],o={render:u,staticRenderFns:r};e.a=o},66:function(t,e){t.exports=function(t,e,n,u,r,o){var s,a=t=t||{},i=typeof t.default;"object"!==i&&"function"!==i||(s=t,a=t.default);var c="function"==typeof a?a.options:a;e&&(c.render=e.render,c.staticRenderFns=e.staticRenderFns,c._compiled=!0),n&&(c.functional=!0),r&&(c._scopeId=r);var l;if(o?(l=function(t){t=t||this.$vnode&&this.$vnode.ssrContext||this.parent&&this.parent.$vnode&&this.parent.$vnode.ssrContext,t||"undefined"==typeof __VUE_SSR_CONTEXT__||(t=__VUE_SSR_CONTEXT__),u&&u.call(this,t),t&&t._registeredComponents&&t._registeredComponents.add(o)},c._ssrRegister=l):u&&(l=u),l){var f=c.functional,d=f?c.render:c.beforeCreate;f?(c._injectStyles=l,c.render=function(t,e){return l.call(e),d(t,e)}):c.beforeCreate=d?[].concat(d,l):[l]}return{esModule:s,exports:a,options:c}}},97:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});e.default={SAMPLE_MUTATIONS_PLUS:"SAMPLE_MUTATIONS_PLUS",SAMPLE_ACTIONS_PLUS:"SAMPLE_ACTIONS_PLUS",SAMPLE_ACTIONS_PLUS_ASYNC:"SAMPLE_ACTIONS_PLUS_ASYNC"}},99:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var u=n(150),r=function(t){return t&&t.__esModule?t:{default:t}}(u);e.default={methods:{handleSendRequest:function(t){r.default.test().then(function(t){alert("send request successfully !")}).catch(function(t){alert("send request fail !")})}}}}},[142]);
//# sourceMappingURL=index.641f393e3f2d86b90bf8.js.map