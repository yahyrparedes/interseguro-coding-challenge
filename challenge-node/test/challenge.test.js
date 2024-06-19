'use strict'

import {expect} from 'chai';
import request from 'request'

const url = 'http://localhost:3002'

describe('POST /api/v1/challenge', function () {
    it('should respond with status 200 and our posted data', function (done) {
        let data = {
            q: [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]],
            r: [[1, 2, 3, 4], [5, 6, 7, 8]],
        };
        request.post({url: url + '/api/v1/challenge', form: data}, function (err, res, body) {
            if (err) {
                return console.error('An error occurred:');
            } else {
                let parse = JSON.parse(body);
                expect(res.statusCode).to.equal(200);
                expect(parse.results.min).to.equal(1);
            }
            done();
        });
    });
});