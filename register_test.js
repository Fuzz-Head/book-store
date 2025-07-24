import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    vus: 150,            // virtual users
    duration: '1m',     // test duration
};

export default function () {
    const unique = Math.floor(1000000000 + Math.random() * 9000000);//Date.now() + Math.random().toString(36).substring(2, 12);
    const payload = JSON.stringify({
        username: `k6user_${unique}`,
        email: `k6_${unique}@test.com`,
        password: "password123",
        role: "user"
    });

    const headers = { 'Content-Type': 'application/json' };

    let res = http.post('http://localhost:8080/register', payload, { headers });

    check(res, {
        'status is 201': (r) => r.status === 201,
    });

    sleep(1); // simulate think time
}

