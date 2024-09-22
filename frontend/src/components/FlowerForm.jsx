import { useState } from 'react'

const FlowerForm = ({ createFlower }) => {
  const [newFlowerName, setNewFlowerName] = useState('')
  const [newFlowerLatinName, setNewFlowerLatinName] = useState('')

  const addFlower = event => {
    event.preventDefault()
    createFlower({
      name: newFlowerName,
      latin_name: newFlowerLatinName,
    })

    setNewFlowerName('')
    setNewFlowerLatinName('')
  }

  return (
    <div>
      <form onSubmit={addFlower}>
        <div>
          Name: <input id="newFlowerNameInput" value={newFlowerName} onChange={event => setNewFlowerName(event.target.value)} />
        </div>
        <div>
          Latin name: <input id="newFlowerLatinNameInput" value={newFlowerLatinName} onChange={event => setNewFlowerLatinName(event.target.value)}/>
        </div>
        <div>
          <button id="saveNewFlowerButton" type="submit">Save</button>
        </div>
      </form>
    </div>
  )
}

export default FlowerForm
