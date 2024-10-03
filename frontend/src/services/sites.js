import axios from 'axios'
const baseUrl = '/api/sites'

/*
  Create a new site
  input
    newSite : {
      name: string,
      note: string,
      parent: ID
    }
    name is required
    if parent is null, site is attached as a root site

  output
    {
      _id: ID,
    	name: string,
    	added_time: Date,
    	note: string,
    	parent: ID
    	flowers: [flower],
    	owner: ID
    }
*/
const create = newSite =>  {
  const config = {
    headers: { Authorization: fetchToken() },
  }
  const request = axios.post(baseUrl, newSite, config)
  return request.then(response => response.data)
}

/*
  Get a site and its subsites
  input
    id: ID

  output
    {
      site: {
        _id: ID,
      	name: string,
      	added_time: Date,
      	note: string,
      	parent: ID
      	flowers: [flower],
      	owner: ID
    	},
    	subsites: [
    	  {
    	    _id: ID,
    	    name: string,
    	    note: string
    	  } ]
    }
*/
const get = (id = null) => {
  const config = {
    headers: { Authorization: fetchToken() },
  }
  const url = id ? `${baseUrl}/${id}` : baseUrl
  return axios.get(url, config).then(response => response.data)
}

/*
  Remove site from database
  input
    id: ID

  output
    {
      "Deleted count": int
    }
*/
const remove = id => {
  const config = {
    headers: { Authorization: fetchToken() },
  }
  return axios.delete(`${baseUrl}/${id}`, config)
}

const fetchToken = () => {
  return localStorage.getItem("token")
}

export default {
  create, getRoot, get, remove
}
