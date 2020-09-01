import axios from "axios";
import store from "@/store";
import { Message } from "element-ui";
import router from "../router";

const service = axios.create({
  baseURL: "/api/", // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 60000 // request timeout
});

//发起请求的拦截器
service.interceptors.request.use(
  config => {
    const token = store.getters.token;;
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  error => {
    console.log(error); // for debug
    return Promise.reject(error);
  }
);

service.interceptors.response.use(
  response => {
    let code = response.data.code;
    if (code === 1008) {
      router.push("/login");
    } else {
      return response.data;
    }
  },
  error => {
    console.log(error);
    Message({
      message: error,
      type: "error",
      duration: 5 * 1000
    });
  }
);
export default service;
