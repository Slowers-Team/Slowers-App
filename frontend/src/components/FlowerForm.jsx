import { useState } from "react"
import { useTranslation } from "react-i18next"

const FlowerForm = ({ createFlower, siteID }) => {
  const [newFlowerName, setNewFlowerName] = useState("")
  const [newFlowerLatinName, setNewFlowerLatinName] = useState("")
  const { t, i18n } = useTranslation()

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
          <label htmlFor="newFlowerNameInput">{t("flower.data.name")}:</label>
          <input
            id="newFlowerNameInput"
            value={newFlowerName}
            onChange={event => setNewFlowerName(event.target.value)}
          />
        </div>
        <div>
          <label htmlFor="newFlowerLatinNameInput">{t("flower.data.latinname")}:</label>
          <input
            id="newFlowerLatinNameInput"
            value={newFlowerLatinName}
            onChange={event => setNewFlowerLatinName(event.target.value)}
          />
        </div>
        <div>
          <button id="saveNewFlowerButton" type="submit">
            {t("button.save")}
          </button>
        </div>
      </form>
    </div>
  )
}

export default FlowerForm
