import http from "k6/http";
import { sleep, check } from "k6";
import { paramsFile, payloadFiles } from "../const.js";
export { options } from "../const.js";

export default function () {
  const res = http.post(
    "http://localhost:8082/ocr-single",
    fd.body(),
    paramsFile
  );
  check(res, {
    "status is 200": (r) => r.status === 200,
  });

  sleep(1);
}
