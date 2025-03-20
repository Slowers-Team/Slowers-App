import axios from 'axios'
import tokenService from './token'

const baseUrl = '/api/business/create'

const create = newBusiness => {
    console.log(newBusiness)
    const config = {
        headers: { Authorization: tokenService.fetchToken() }
    }
    const request = axios.post(baseUrl, newBusiness, config)
    return request.then(response => response.data)
}


export default {
    create
}