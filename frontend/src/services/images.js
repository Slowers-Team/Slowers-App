import axios from 'axios'
const baseUrl = '/api/images'
import tokenService from './token'

// get URL for an imageObject
const get = imageObject => {
  const filename = getFilename(imageObject);
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
    responseType: "blob"
  };
  return axios.get(`${baseUrl}/${filename}`, config)
    .then(response => {
      const imageUrl = URL.createObjectURL(response.data);
      return { _id: imageObject._id, url: imageUrl };
    })
    .catch(error => console.error("Error fetching image blob:", error));
};

const getByID = id => {
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
    'Content-Type': "application/json", 
    responseType: "blob"
  };
  return axios.get(`${baseUrl}/id/${id}`, config)
    .then(response => {
      const imageUrl = URL.createObjectURL(response.data);
      return imageUrl;
    })
    .catch(error => {
      // console.error("Error fetching image blob:", error); 
      throw error});
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

const setFavorite = (entityID, entityType, imageID) => {
  if (!(entityType === "flower" || entityType === "site")) {
    throw "Invalid entity type"
  }
  
  const url = `${baseUrl}/favorite`
  const config = {
    headers: { Authorization: tokenService.fetchToken() },
    'Content-Type': 'application/json'
  }

  const data = {
    EntityID: entityID,
    EntityType: entityType,
    ImageID: imageID
  }

  return axios.post(url, data, config)
    .then(_ => {
      return true
    })
    .catch(error => {
        console.error("Failed to set favorite image of", entityType, entityID, "set to", imageID, ":\n", error?.error)
      throw error.response.data
    })
}

const getFilename = image => image._id + "." + image.file_format 

export default {
  get,
  create,
  getImagesByEntity,
  deleteImage,
  setFavorite,
  getByID
}
