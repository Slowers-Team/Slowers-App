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
    <div className="text-left">
      <form onSubmit={addFlower}>
        <div className="form-group">
          <label htmlFor="newFlowerNameInput">{t("flower.data.name")}:</label>
          <input
            id="newFlowerNameInput"
            value={newFlowerName}
            onChange={event => setNewFlowerName(event.target.value)}
            className="form-control"
          />
        </div>
        <div className="form-group">
          <label htmlFor="newFlowerLatinNameInput">{t("flower.data.latinname")}:</label>
          <input
            id="newFlowerLatinNameInput"
            value={newFlowerLatinName}
            onChange={event => setNewFlowerLatinName(event.target.value)}
            className="form-control"
          />
        </div>
        <div>
          <button id="saveNewFlowerButton" type="submit" className="btn btn-light">
            {t("button.save")}
          </button>
        </div>
      </form>
    </div>
  )
}

export default FlowerForm
