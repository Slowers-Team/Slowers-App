/* eslint-disable react/prop-types */
import axios from 'axios'
import { useEffect, useState } from 'react'
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

const FlowerForm = ({ event, name, handleFlowerNameChange, latin_name, handleFlowerLatinNameChange }) => {
  return (
    <div>
      <form onSubmit={event}>
        <div>
          Name: <input value={name} onChange={handleFlowerNameChange} />
        </div>
        <div>
          Latin name: <input value={latin_name} onChange={handleFlowerLatinNameChange}/>
        </div>
        <div>
          <button type='submit'>Save</button>
        </div>
      </form>
    </div>
  )
}

const FlowerList = ({ flowers, handleDelete }) => {
  return (
    <table id="flowerList">
      <thead>
        <tr>
          <th>Name</th>
          <th>Latin name</th>
          <th>Added time</th>
        </tr>
      </thead>
      <tbody>
        {flowers.map(flower => {
          let addedTime = new Date(flower.added_time)

          let date = addedTime.toLocaleDateString('fi')
          let hour = addedTime.toLocaleString('fi', {hour: 'numeric'})
          let minute = addedTime.toLocaleString('fi', {minute: '2-digit'})
          let addedTimeStr = `${date} ${hour}:${minute}`

          return (
            <tr key={flower.id}>
              <td>{ flower.name }</td>
              <td><em>{ flower.latin_name }</em></td>
              <td>{ addedTimeStr }</td>
              <td>
                <button onClick={() => handleDelete(flower)}>Delete</button>
              </td>
            </tr>
          )
        })}
      </tbody>
    </table>
  )
}

export default App
