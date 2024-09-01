import http from 'k6/http';
import { check } from 'k6';

export const options = {
    thresholds: {
        http_req_failed: ['rate==0'],
        http_req_duration: ['p(95)<1'], // 95% should respond in less than 1ms
    },
    scenarios: {
        constant_request_rate: {
            executor: 'constant-arrival-rate',
            rate: 5000,     // 5000 requests per second
            timeUnit: '1s',
            duration: '10s',
            preAllocatedVUs: 50,
            maxVUs: 100,
        },
    }
};

const searchTerms = ['apple', 'banana', 'prejudice', 'zapara']

export default function () {

    const searchTerm = searchTerms[Math.floor(Math.random() * searchTerms.length)];

    const url = `http://localhost:8080/?q=${searchTerm}`
    const res = http.get(url);

    check(res, {
        'is status 200': (r) => r.status === 200,
    });
}
