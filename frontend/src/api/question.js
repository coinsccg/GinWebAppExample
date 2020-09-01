import request from "@/utils/request";

export function saveQuestion(data) {
  return request({
    url: "/question",
    method: "post",
    data
  });
}

export function updateQuestion(data) {
  return request({
    url: "/updateQuestion",
    method: "post",
    data
  });
}

export function deleteQuestion(params) {
  return request({
    url: "/deleteQuestion",
    method: "get",
    params
  });
}

export function queryQuestions(params) {
  return request({
    url: "/questions",
    method: "get",
    params
  });
}

export function getQuestion(id) {
  return request({
    url: "/question?question_id="+id,
    method: "get"
  });
}

export function saveSubscribe(params) {
  return request({
    url: "/saveSubscribe",
    method: "get",
    params
  });
}

export function deleteSubscribe(params) {
  return request({
    url: "/deleteSubscribe",
    method: "get",
    params
  });
}
