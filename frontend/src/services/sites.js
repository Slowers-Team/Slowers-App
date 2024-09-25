import axios from 'axios'
const baseUrl = '/api/---'

const create = newSite =>  {
    const request = axios.post(baseUrl, newSite)
    return request.then(response => response.data)
}

export default {
    create
}
