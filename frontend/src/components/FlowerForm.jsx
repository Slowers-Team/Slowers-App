import { Container } from 'react-bootstrap'
import { useState } from "react"
import { useTranslation } from "react-i18next"

const FlowerForm = ({ createFlower, siteID, handleClose }) => {
  const [newFlowerName, setNewFlowerName] = useState("")
  const [newFlowerLatinName, setNewFlowerLatinName] = useState("")
  const [newFlowerQty, setNewFlowerQty] = useState("0")
  const { t, i18n } = useTranslation()

  const addFlower = event => {
    event.preventDefault()
    createFlower({
      name: newFlowerName,
      latin_name: newFlowerLatinName,
      site: siteID,
      quantity: Number(newFlowerQty),
    })

    setNewFlowerName("")
    setNewFlowerLatinName("")
    setNewFlowerQty("0")
    handleClose()
  }

  return (
    <Container>
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
        <div className="form-group">
          <label htmlFor="newFlowerQtyInput">{t("flower.data.qty")}:</label>
          <input
            type="number"
            id="newFlowerQtyInput"
            value={newFlowerQty}
            onChange={event => setNewFlowerQty(event.target.value)}
            className="form-control"
            min="0"
            required
          />
        </div>

        <div>
          <button id="saveNewFlowerButton" type="submit" className="btn btn-light my-2">
            <i className="bi bi-floppy2-fill"> </i>
            {t("button.save")}
          </button>
        </div>
      </form>
    </Container>
  )
}

export default FlowerForm
