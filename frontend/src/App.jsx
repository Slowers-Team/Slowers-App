/* eslint-disable react/prop-types */
import axios from 'axios'
import { useState } from 'react'

const App = () => {
  const [newFlowerName, setNewFlowerName] = useState('')
  const [newFlowerLatinName, setNewFlowerLatinName] = useState('')
  const [showAddNewFlower, setShowAddNewFlower] = useState(false)

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
        setNewFlowerName('')
        setNewFlowerLatinName('')
      })
      .catch(error => {
        console.log(error)
        alert(`Adding failed`)
      })
  }

  const handleFlowerNameChange = (event) => {
    setNewFlowerName(event.target.value)
  }

  const handleFlowerLatinNameChange = (event) => {
    setNewFlowerLatinName(event.target.value)
  }
import axios from 'axios'
import { useEffect, useState } from 'react'
import './App.css'

const App = () => {
  const [flowers, setFlowers] = useState([])

  useEffect(() => {
    axios
      .get('/api/flowers')
      .then(response => setFlowers(response.data))
  }, [])

  return (
    <>
      <button onClick={() => setShowAddNewFlower(!showAddNewFlower)}>Add a new flower</button>
      {showAddNewFlower && <FlowerForm event={addFlower} name={newFlowerName} handleFlowerNameChange={handleFlowerNameChange} latin_name={newFlowerLatinName} handleFlowerLatinNameChange={handleFlowerLatinNameChange}/>}
      {flowers && <FlowerList flowers={flowers} />}
    </>
  )
}

const FlowerForm = ({ event, name, handleFlowerNameChange, latin_name, handleFlowerLatinNameChange }) => {
  return (
    <div>
      <form onSubmit={event}>
        <div>
          name: <input value={name} onChange={handleFlowerNameChange} />
        </div>
        <div>
          latin name: <input value={latin_name} onChange={handleFlowerLatinNameChange}/>
        </div>
        <div>
          <button type='submit'>save</button>
        </div>
      </form>
    </div>
  )
}

const FlowerList = ({ flowers }) => {
  return (
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Latin name</th>
          <th>Added time</th>
        </tr>
      </thead>
      <tbody>
        {flowers.map(flower => (
          <tr key={flower.id}>
            <td>{ flower.name }</td>
            <td><em>{ flower.latin_name }</em></td>
            <td>{ new Date(flower.added_time).toDateString() }</td>
          </tr>
        ))}
      </tbody>
    </table>
  )
}

export default App