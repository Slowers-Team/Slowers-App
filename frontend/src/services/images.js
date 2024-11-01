import axios from 'axios'
const baseUrl = '/api/images'
import tokenService from './token'

const get = filename => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
    responseType: "blob"
  }
  const request = axios.get(`${baseUrl}/${filename}`, config)
  return request.then(response => {
    return URL.createObjectURL(response.data)
  })
}

const create = imageObject => {
  const config = {
    headers: { 
      Authorization: tokenService.fetchToken(),
      'Content-Type': 'multipart/form-data'
       },
  }
  const request = axios.post(baseUrl, imageObject, config)
  return request.then(response => response.data)
}

const getFilename = image => image._id + "." + image.file_format 


export default {
  get,
  getFilename,
  create,
}
