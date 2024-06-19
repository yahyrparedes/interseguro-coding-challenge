import { Router } from 'express'
const router = Router()


router.get('/', index)
router.get('/ping', ping)

function index(req, res){
    res.send('Is Running')
}

function ping(req, res){
    res.send('pong')
}

export default router