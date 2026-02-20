import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "10s", target: 50 },
    { duration: "20s", target: 50 },
    { duration: "10s", target: 0 },
  ],
};

export default function () {
  const res = http.get("http://localhost:8080/profile?userId=1");

  check(res, {
    "status is 200": (r) => r.status === 200,
    "has userId": (r) => r.body.includes('"userId":1'),
  });

  sleep(1);
}

