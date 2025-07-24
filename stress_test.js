import http from 'k6/http';
import {check, sleep} from 'k6';

export let options = {
  stages: [
    { duration: '30s', target: 20 },
    { duration: '1m30s', target: 10 },
    { duration: '20s', target: 2 },
  ],
};

export default function () {
  const res = http.get('http://localhost:8080/books', {
    headers: {
      Authorization: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNyIsInJvbGUiOiJhZG1pbiIsInNjb3BlcyI6WyJjYW46cmVhZDpib29rcyIsImNhbjpyZWFkOmJvb2siLCJjYW46Y3JlYXRlOmJvb2siLCJjYW46dXBkYXRlOmJvb2siLCJjYW46ZGVsZXRlOmJvb2siXSwidHlwZSI6ImFjY2Vzc190b2tlbiIsImV4cCI6MTc1MzMzNTIzMywiaWF0IjoxNzUzMzMxNjMzfQ.sSyXCD8InuWC7mVls-RfxZQvnyp4bLb0FYCqVW_thdQ', 
    },
  });

  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  sleep(1);
}
