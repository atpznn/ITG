import { check } from "k6";
import http from "k6/http";
import { paramsFile, payloadFiles } from "../const.js";
import { sleep } from "k6";
export { options } from "../const.js";
export default function () {
  console.log(payloadFiles);
  const res = http.post(
    "https://itg-go.zeabur.app/ocr-single-safe",
    payloadFiles.body(),
    paramsFile
  );

  const isOk = check(res, {
    "status is 200": (r) => r.status === 200,
    "is json": (r) =>
      r.headers["Content-Type"] &&
      r.headers["Content-Type"].includes("application/json"),
  });

  if (isOk) {
    try {
      const body = res.json();
      check(res, {
        "has results": (r) => body.results && body.results.length > 0,
        "valid content": (r) => body.results[0].length > 10,
      });
    } catch (e) {
      console.log("Failed to parse JSON. Raw body: " + res.body);
    }
  } else {
    console.log(`Error! Status: ${res.status}, Body: ${res.body}`);
  }

  sleep(2);
}
