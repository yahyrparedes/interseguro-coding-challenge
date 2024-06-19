import { Router } from 'express'
const router = Router()

import challenge from './challenge.route.js'


router.use('/v1', challenge)

export default router