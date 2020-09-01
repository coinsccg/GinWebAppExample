import request from "@/utils/request";

export function restore() {
  return request({
    url: "/restore",
    method: "get"
  });
}
export function flush() {
  return request({
    url: "/flush",
    method: "get"
  });
}
export function backup() {
  return request({
    url: "/backup",
    method: "get"
  });
}

export function createIndex() {
  return request({
    url: "/createIndex",
    method: "get"
  });
}
