import axios from 'axios'
import tokenService from './token'

const baseUrl = '/api/business'

const create = business => {
    console.log(business)
    const config = {
        headers: { Authorization: tokenService.fetchToken() }
    }
    const request = axios.post(baseUrl, business, config)
    return request.then(response => response.data)
}


export default {
    create
}