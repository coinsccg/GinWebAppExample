import Vue from "vue";
import Vuex from "vuex";
Vue.use(Vuex);

const defaultLoginResult = {
  token: null,
  userID: null,
  username: null,
}

export default new Vuex.Store({
  state: {
    loginResult: defaultLoginResult,
  },
  mutations: {
    init(state){
      const loginResult = JSON.parse(localStorage.getItem("loginResult"));
      console.log(localStorage.getItem("loginResult"))
      if (loginResult != null){
        state.loginResult = loginResult;
      }
    },
    login(state, userInfo){
      console.log(userInfo)
      state.loginResult = userInfo;
    },
    logout(state){
      localStorage.removeItem("loginResult");
      state.loginResult = defaultLoginResult;
    }
  },
  actions: {},
  getters: {
    isLogin: state => state.loginResult.userID !== null,
    userID: state => state.loginResult.userID,
    username: state => state.loginResult.username,
    token: state => state.loginResult.token,
  }
});
