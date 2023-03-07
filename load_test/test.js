import http from 'k6/http';

import { check, group, sleep } from 'k6';


export const options = {

    stages: [

        { duration: '30s', target: 1000 },
        { duration: '1m', target: 1000 },
        { duration: '30s', target: 0 },

    ],

    thresholds: {

        'http_req_duration': ['p(99)<500'],


    },

};


const BASE_URL = 'http://booking:8088';

export default () => {

    const maximize = http.post(`${BASE_URL}/maximize`,
        `[
{
"request_id":"bookata_XY123",
"check_in":"2020-01-01",
"nights":5,
"selling_rate":200,
"margin":20
},
{
"request_id":"kayete_PP234",
"check_in":"2020-01-04",
"nights":4,
"selling_rate":156,
"margin":5
},
{
"request_id":"atropote_AA930",
"check_in":"2020-01-04",
"nights":4,
"selling_rate":150,
"margin":6
},
{
"request_id":"acme_AAAAA",
"check_in":"2020-01-10",
"nights":4,
"selling_rate":160,
"margin":30
}
]`, {headers: {'Content-Type': 'application/json'}});


    check(maximize, {

        'maximize is successfully': (r) => r.status === 200,

    });

    const stats = http.post(`${BASE_URL}/stats`,
        `[
    {
    "request_id":"bookata_XY123",
    "check_in":"2020-01-01",
    "nights":5,
    "selling_rate":200,
    "margin":20
    },
    {
    "request_id":"kayete_PP234",
    "check_in":"2020-01-04",
    "nights":4,
    "selling_rate":156,
    "margin":5
    }
]`, {headers: {'Content-Type': 'application/json'}});


    check(stats, {

        'maximize is successfully': (r) => r.status === 200,

    });


};