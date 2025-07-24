import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  stages: [
    { duration: '1m0s', target: 80 }, // ramp-up to 10 users
    { duration: '2m0s', target: 80 }, // hold at 10 users
    { duration: '45s', target: 0 },  // ramp-down
  ],
};

const BASE_URL = 'http://localhost:8080';

export default function () {
  // Unique suffix per virtual user
  const unique = Math.floor(Math.random() * 1000000);
  const username = `user_${unique}`;
  const email = `user_${unique}@example.com`;
  const password = 'secretpass';
  const role = 'admin';

  // Register user
  const registerPayload = JSON.stringify({
    username: username,
    email: email,
    password: password,
    role: role,
  });

  const registerHeaders = { 'Content-Type': 'application/json' };
  let regRes = http.post(`${BASE_URL}/register`, registerPayload, { headers: registerHeaders });
  check(regRes, {
    'register status is 201': (r) => r.status === 201 || r.status === 409,
  });

  // Login
  const loginPayload = JSON.stringify({
    email: email,
    password: password,
  });

  let loginRes = http.post(`${BASE_URL}/login`, loginPayload, { headers: registerHeaders });
  check(loginRes, {
    'login status is 200': (r) => r.status === 200,
  });

  if (loginRes.status === 200) {
    const tokens = loginRes.json();
    const authHeaders = {
      Authorization: `Bearer ${tokens.access_token}`,
    };

    // Get all books
    let booksRes = http.get(`${BASE_URL}/books`, { headers: authHeaders });
    check(booksRes, {
      'books status is 200': (r) => r.status === 200,
    });

    // Create a book
    const bookPayload = JSON.stringify({
      title: `Stress Test Book ${unique}`,
      author: 'Test Author',
      price: 10.99,
      isbn: `9780002011232`,
    });
 
    //if (tokens.access_token != ""){
    const bookRes = http.post(`${BASE_URL}/book`, bookPayload, { 
      headers: {...authHeaders, 'Content-Type': 'application/json'}
    });

    check(bookRes, {
      'book creation is 201 or 400 (duplicate)': (r) => r.status === 201 || r.status === 400,
    });
   // }
  }

  sleep(1); // Pause between iterations
}

