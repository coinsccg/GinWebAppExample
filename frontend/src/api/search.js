import request from "@/utils/request";

export function searchUserInfos(params) {
  return request({
    url: "/searchUserInfos",
    method: "get",
    params
  });
}

export function searchQuestionInfos(params) {
  return request({
    url: "/searchQuestionInfos",
    method: "get",
    params
  });
}
