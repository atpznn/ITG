import { check } from "k6";
import http from "k6/http";
import { pl } from "../const.js";
import { sleep } from "k6";
export { options } from "../const.js";
export default function () {
  const res = http.post("http://localhost:8081/dime", pl);

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
        "has results": (r) => body.results,
        "valid content": (r) =>
          body.results.DividendLogs.some((x) => x.Symbol == "SGOV"),
      });
    } catch (e) {
      console.log("Failed to parse JSON. Raw body: " + res.body);
    }
  } else {
    console.log(`Error! Status: ${res.status}, Body: ${res.body}`);
  }

  sleep(2);
}
