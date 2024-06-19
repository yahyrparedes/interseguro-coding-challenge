// const winston = require('winston');
import winston from 'winston';


export function logger() {
    new (winston.Logger)({
        transports: [
            new winston.transports.File({
                level: 'info',
                json: false,
                handleExceptions: true,
                maxSize: 512000,
                maxFiles: 5,
                // filename: `${__dirname}/log-de-aplicacion.log`,
                filename: './errors.log',
                prettyPrint: object => {
                    return JSON.stringify(object)
                }
            }),
            new winston.transports.Console({
                level: 'debug',
                handleExceptions: true,
                json: false,
                colorize: true,
                prettyPrint: object => {
                    return JSON.stringify(object)
                }
            })
        ]
    })
}