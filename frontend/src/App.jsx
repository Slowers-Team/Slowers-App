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

  return (
    <>
      <button onClick={() => setShowAddNewFlower(!showAddNewFlower)}>Add a new flower</button>
      {showAddNewFlower && <FlowerForm event={addFlower} name={newFlowerName} handleFlowerNameChange={handleFlowerNameChange} latin_name={newFlowerLatinName} handleFlowerLatinNameChange={handleFlowerLatinNameChange}/>}
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

export default App