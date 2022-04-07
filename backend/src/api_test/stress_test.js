import http from 'k6/http';
import { randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.1.0/index.js';
import { sleep } from 'k6';
export const options = {
    vus: 1,
    duration: "10s"
};

export default function() {
    let data = { "addressFrom": "addressFrom", "addressTo": "addressTo", "amount": 1, "c": 1, "sign": "dtgyhjhgr", "fee": 1 };

    http.request('POST', "http://localhost:8000/addTransation", JSON.stringify(data), {
        headers: { 'Content-Type': 'application/json' },
    });

    sleep(1)

}