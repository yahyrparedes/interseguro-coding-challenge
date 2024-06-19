import {initialize, app} from './server.service.js'


export async function run() {
    try {
        await initialize()
    } catch (err) {
        console.log(`An error has happend: ${err}`)
    }
}

export function getApp() {
    return app
}

// const start = async ()  => {
//     try {
//         await initialize()
//     }catch (err) {
//         console.log(`An error has happend: ${err}`)
//     }
// }
//
// export default start