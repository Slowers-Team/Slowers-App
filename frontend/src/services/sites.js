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
  const request = axios.post(baseUrl, newSite)
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
const get = id =>  {
  const request = axios.get(`${baseUrl}/${id}`)
  return request.then(response => response.data)
}

/*
  Get all root sites (sites that have no parents)
  output
  [
    {
      _id: ID,
    	name: string,
    	added_time: Date,
    	note: string,
    	parent: ID
    	flowers: [flower],
    	owner: ID
    }
  ]
  
*/
const getRoot = () => {
  const request = axios.get(baseUrl)
  return request.then(response => response.data)
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
  return axios.delete(`${baseUrl}/${id}`)
}

export default {
  create, getRoot, get, remove
}
