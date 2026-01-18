import http from "k6/http";
import { sleep, check } from "k6";
import { params, paramsFile, payload, payloadFiles, pl } from "./const.js";
export { options } from "./const.js";
export default function () {
  const res = http.post("https://itg-ts.onrender.com/image-process", pl);
  const isOk = check(res, {
    "status is 200": (r) => r.status === 200,
    "is json": (r) =>
      r.headers["Content-Type"] &&
      r.headers["Content-Type"].includes("application/json"),
  });

  if (isOk) {
    try {
      check(res, {
        "has results": (r) => res.body.length,
        "valid content": (r) => res.body.includes("SGOV 1.72 USD"),
      });
    } catch (e) {
      console.log("Failed to parse JSON. Raw body: " + res.body);
    }
  } else {
    console.log(`Error! Status: ${res.status}, Body: ${res.body}`);
  }

  sleep(2);
}
