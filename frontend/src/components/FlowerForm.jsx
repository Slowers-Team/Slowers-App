import { useState } from "react"

const FlowerForm = ({ createFlower, siteID }) => {
  const [newFlowerName, setNewFlowerName] = useState("")
  const [newFlowerLatinName, setNewFlowerLatinName] = useState("")

  const addFlower = event => {
    event.preventDefault()
    createFlower({
      name: newFlowerName,
      latin_name: newFlowerLatinName,
      site: siteID,
    })

    setNewFlowerName("")
    setNewFlowerLatinName("")
  }

  return (
    <div>
      <form onSubmit={addFlower}>
        <div>
          <label htmlFor="newFlowerNameInput">Name:</label>
          <input
            id="newFlowerNameInput"
            value={newFlowerName}
            onChange={event => setNewFlowerName(event.target.value)}
          />
        </div>
        <div>
          <label htmlFor="newFlowerLatinNameInput">Latin name:</label>
          <input
            id="newFlowerLatinNameInput"
            value={newFlowerLatinName}
            onChange={event => setNewFlowerLatinName(event.target.value)}
          />
        </div>
        <div>
          <button id="saveNewFlowerButton" type="submit">
            Save
          </button>
        </div>
      </form>
    </div>
  )
}

export default FlowerForm
