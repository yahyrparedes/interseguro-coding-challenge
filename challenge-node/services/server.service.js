// Imports
import { createServer } from 'http'
import express from 'express'
import bodyParser from 'body-parser'
import cors from 'cors'
import {logger} from './../utils/logger/index.js'
import morgan from 'morgan';

import routeWeb from '../routes/index.js'
import routeAPI from '../routes/api/index.js'

const app = express()
app.disable('x-powered-by')
const server = createServer(app)

// cors
app.use(cors())

// Implements Logger using Morgan
if (process.env.APP_ENV === 'prod') {
    app.use(morgan('dev', {
        stream: {
            write: message => logger.info(message.trim()),
        }
    }));
}

// Middleware
// app.use(json({limit: '50mb'}))
// app.use(urlencoded({limit: '50mb', extended: true}))
app.use(bodyParser.urlencoded());
app.use(bodyParser.json());


// routes
app.use('/', routeWeb)
app.use('/api', routeAPI)


const initialize = async () => {
    const promise = new Promise((resolve, reject) => {
        app.listen(process.env.PORT || 3002)
            .on('listening', () => {
                console.log(`Server is running on port ${process.env.PORT || '3000'} `)
                resolve()
            })
            .on('error', (err) => {
                console.log(`Error in server ${err}`)
                reject(err)
            })

        server.listen()
        Promise.resolve()
    })

    await promise
}

export { initialize, app}
// export default {initialize, app}