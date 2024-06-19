// const commonRepository = require('./../repositories/common.repository')
import {responseResult, httpStatus} from '../utils/response.util.js';
// import { BAD_REQUEST, OK } from 'http-status-codes';
import { isMatrix, getMax, getMin, calculateAverage, sum, isDiagonal } from '../utils/matrix.util.js';

export const ChallengeController = {

    async challenge(req, res, next) {
        const body = req.body

        console.log('body', body)

        // validate is matrix R and Q
        if (!body.q || !body.r || !isMatrix(body.q) || !isMatrix(body.r)) {
            return responseResult.error(res, httpStatus.BAD_REQUEST,
                'Bad request, check the data sent');
        }

        const maxQ = getMax(body.q)
        const minQ = getMin(body.q)
        const minR = getMin(body.r)
        const maxR = getMax(body.r)

        let max = maxQ > maxR ? maxQ : maxR;
        let min = minQ < minR ? minQ : minR;

        const response = {
            min: min,
            max: max,
            avg: calculateAverage(body.q) + calculateAverage(body.r) / 2,
            sum: sum(body.q) + sum(body.r),
            isDiagonalQ: isDiagonal(body.q),
            isDiagonalR: isDiagonal(body.r),
        }

        return responseResult.general(res, httpStatus.OK, response)
    },
}
 
