import e from "express";

export const httpStatus = {
    OK: 200,
    CREATED: 201,
    NO_CONTENT: 204,
    BAD_REQUEST: 400,
    UNAUTHORIZED: 401,
    FORBIDDEN: 403,
    NOT_FOUND: 404,
    INTERNAL_SERVER_ERROR: 500,
};

export  const responseResult = {
    general(res, status, results) {
        console.log('status', status, 'results', results)
        res.status(status).json({
            status,
            results,
        });
    },
    error(res, status, message, stack = undefined) {
        const response = { status, message };

        if (stack) {
            response.stack = stack;
        }
        console.log('status', status, 'response', response)
        res.status(status).json(response);
    },
};

// module.exports = responseResult;