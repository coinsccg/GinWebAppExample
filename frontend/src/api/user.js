import request from "@/utils/request";

export function login(data) {
  return request({
    url: "/login",
    method: "post",
    data
  });
}

export function logout(params) {
  return request({
    url: "/logout",
    method: "get",
    params
  });
}

export function getUserInfo(params) {
  return request({
    url: "/getUserInfo",
    method: "get",
    params
  });
}

export function queryUserInfos(data) {
  return request({
    url: "/queryUserInfos",
    method: "post",
    data
  });
}

export function saveFollow(params) {
  return request({
    url: "/saveFollow",
    method: "get",
    params
  });
}

export function deleteFollow(params) {
  return request({
    url: "/deleteFollow",
    method: "get",
    params
  });
}

export function queryUserInfosByFollowerId(params) {
  return request({
    url: "/queryUserInfosByFollowerId",
    method: "get",
    params
  });
}

export function queryFollowerInfosByPeopleId(params) {
  return request({
    url: "/queryFollowerInfosByPeopleId",
    method: "get",
    params
  });
}

export function signup(data) {
  return request({
    url: "/signup",
    method: "post",
    data
  });
}

export function updateAuthority(data) {
  return request({
    url: "/updateAuthority",
    method: "post",
    data
  });
}

export function updateUser(data) {
  return request({
    url: "/updateUser",
    method: "post",
    data
  });
}

export function updatePassword(data) {
  return request({
    url: "/updatePassword",
    method: "post",
    data
  });
}

export function updateAvater(data) {
  return request({
    url: "/updateAvater",
    method: "post",
    data
  });
}

export function queryPeopleActivities(params) {
  return request({
    url: "/queryPeopleActivities",
    method: "get",
    params
  });
}

export function queryFollowActivities(params) {
  return request({
    url: "/queryFollowActivities",
    method: "get",
    params
  });
}

export function queryAnswersByUserId(params) {
  return request({
    url: "/queryAnswersByUserId",
    method: "get",
    params
  });
}

export function queryQuestionsByUserId(params) {
  return request({
    url: "/queryQuestionsByUserId",
    method: "get",
    params
  });
}

export function queryTopicsByUserId(params) {
  return request({
    url: "/queryTopicsByUserId",
    method: "get",
    params
  });
}

export function querySubscribeQuestionsByUserId(params) {
  return request({
    url: "/querySubscribeQuestionsByUserId",
    method: "get",
    params
  });
}
