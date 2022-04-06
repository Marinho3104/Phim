import http from 'k6/http';
import { randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.1.0/index.js';
import { sleep } from 'k6';
export const options = {
    vus: 120,
    duration: "10s"
};

export default function() {
    let data = { "addressFrom": "addressFrom", "addressTo": "addressTo", "amount": 1, "c": randomIntBetween(1, 5000), "sign": "", "fee": 1 };

    http.post("http://localhost:8000/addTransation", JSON.stringify(data), {
        headers: { 'Content-Type': 'application/json' },
    });

}