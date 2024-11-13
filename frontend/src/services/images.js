import axios from 'axios'
const baseUrl = '/api/images'
import tokenService from './token'

// get URL for an imageObject
const get = imageObject => {
  const filename = getFilename(imageObject);
  console.log("Generated filename:", filename);
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
    responseType: "blob"
  };
  return axios.get(`${baseUrl}/${filename}`, config)
    .then(response => {
      const imageUrl = URL.createObjectURL(response.data);
      console.log("Image URL:", imageUrl); 
      return { _id: imageObject._id, url: imageUrl };
    })
    .catch(error => console.error("Error fetching image blob:", error));
};

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

// get list of URLs for an entity
const getImagesByEntity = entityId => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
    responseType: "json" 
  }
  return axios.get(`${baseUrl}/entity/${entityId}`, config)
    .then( response => 
      Promise.all(response.data.map(object => get(object)))
    )
}
const deleteImage = id => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
  }
  return axios.delete(`${baseUrl}/${id}`, config)
    .then(response => response.data)
    .catch( error => {
      console.error("Error deleting image:", error)
      throw error;
    })
}

const getFilename = image => image._id + "." + image.file_format 

export default {
  get,
  create,
  getImagesByEntity,
  deleteImage,
}