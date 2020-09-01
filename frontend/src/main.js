// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import App from "./App";
import router from "./router";
import store from "./store";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";

Vue.config.productionTip = false;
Vue.use(ElementUI);

import {
  getTimeString,
  getCountString,
  nowTimestamp,
  userInfo,
  avaterUrl,
  questionInfo2ActvityInfo
} from "./js/utils";
Vue.prototype.$getTimeString = getTimeString;
Vue.prototype.$getCountString = getCountString;
Vue.prototype.$nowTimestamp = nowTimestamp;
Vue.prototype.$userInfo = userInfo;
Vue.prototype.$avaterUrl = avaterUrl;
Vue.prototype.$questionInfo2ActvityInfo = questionInfo2ActvityInfo;
/* eslint-disable no-new */
new Vue({
  el: "#app",
  router,
  store,
  components: { App },
  template: "<App/>"
});
