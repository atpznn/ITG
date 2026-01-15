import http from "k6/http";
import { sleep, check } from "k6";
import { params, payload } from "./const.js";
export { options } from "/const.js";
export default function () {
  const res = http.post(
    "https://itg-typescript.zeabur.app/v1/dime/process-text",
    JSON.stringify(payload),
    params
  );
  check(res, {
    "status is 200": (r) => r.status === 200,
  });

  sleep(1);
}
