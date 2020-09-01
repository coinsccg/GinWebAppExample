import request from "@/utils/request";

export function queryReviews(params) {
  return request({
    url: "/queryReviews",
    method: "get",
    params
  });
}

export function deleteReview(params) {
  return request({
    url: "/deleteReview",
    method: "get",
    params
  });
}

export function saveReview(data) {
  return request({
    url: "/saveReview",
    method: "post",
    data
  });
}

export function updateApprove(params) {
  return request({
    url: "/updateApprove",
    method: "get",
    params
  });
}
