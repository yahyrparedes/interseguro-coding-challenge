import { Router } from 'express'
const router = Router()

import { ChallengeController } from './../../controllers/challenge.controller.js'

router.post('/challenge', ChallengeController.challenge)


export default router