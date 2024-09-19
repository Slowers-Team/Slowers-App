/* eslint-disable react/prop-types */
import axios from 'axios'
import { useEffect, useState } from 'react'
import FlowerForm from './components/FlowerForm'
import FlowerList from './components/FlowerList'
import './App.css'

const App = () => {
  const [flowers, setFlowers] = useState([])
  const [newFlowerName, setNewFlowerName] = useState('')
  const [newFlowerLatinName, setNewFlowerLatinName] = useState('')
  const [showAddNewFlower, setShowAddNewFlower] = useState(false)

  useEffect(() => {
    axios
      .get('/api/flowers')
      .then(response => setFlowers(response.data))
  }, [])

  const addFlower = event => {
    event.preventDefault()
    const flowerObject = {
      name: newFlowerName,
      latin_name: newFlowerLatinName
    }

    axios
      .post('/api/flowers', flowerObject)
      .then(response => {
        console.log(response)
        if (!flowers) {
          setFlowers([response.data])
        } else {
          setFlowers(flowers.concat(response.data))
        }
        setNewFlowerName('')
        setNewFlowerLatinName('')
      })
      .catch(error => {
        console.log(error)
        alert(`Adding failed`)
      })
  }

  const handleDelete = (flower) => {
    if (window.confirm(`Are you sure you want to delete flower ${flower.name}?`)) {
      axios
      .delete(`/api/flowers/${flower._id}`)
      .then(response => {
        console.log(response)
        setFlowers(l => l.filter(item => item.name !== flower.name));
      })
    }
  }

  const handleFlowerNameChange = (event) => {
    setNewFlowerName(event.target.value)
  }

  const handleFlowerLatinNameChange = (event) => {
    setNewFlowerLatinName(event.target.value)
  }

  return (
    <>
      <button onClick={() => setShowAddNewFlower(!showAddNewFlower)}>Add a new flower</button>
      {showAddNewFlower && <FlowerForm event={addFlower} name={newFlowerName} handleFlowerNameChange={handleFlowerNameChange} latin_name={newFlowerLatinName} handleFlowerLatinNameChange={handleFlowerLatinNameChange}/>}
      {flowers && <FlowerList flowers={flowers} handleDelete={handleDelete} />}
    </>
  )
}

export default App
