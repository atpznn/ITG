import http from "k6/http";
import { sleep, check } from "k6";
import { params, paramsFile, payload, payloadFiles, pl } from "./const.js";
export { options } from "/const.js";
export default function () {
  const res = http.post("https://itg-go.onrender.com/ocr-single-safe", pl);
  const isOk = check(res, {
    "status is 200": (r) => r.status === 200,
  });

  if (isOk) {
    try {
      const body = res.json();
      check(res, {
        "has results": (r) => body.results[0].length,
        "valid content": (r) => body.results[0].includes("SGOV 1.72 USD"),
      });
    } catch (e) {
      console.log("Failed to parse JSON. Raw body: " + res.body);
    }
  } else {
    console.log(`Error! Status: ${res.status}, Body: ${res.body}`);
  }

  sleep(2);
}
