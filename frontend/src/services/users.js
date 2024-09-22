import axios from 'axios'
const baseUrl = '/api/register'

const create = newUser =>  {
    const request = axios.post(baseUrl, newUser)
    return request.then(response => response.data)
  }

  export default {
    create
  }
