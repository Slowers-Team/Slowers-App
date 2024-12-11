import { Button, Container } from "react-bootstrap"
import { useState } from "react"
import { useTranslation } from "react-i18next"

const ModifyFlowerForm = ({ flower, modifyFlower, handleFlowerModify, handleFormVisibility }) => {
  const [modifiedFlowerName, setModifiedFlowerName] = useState(flower.name)
  const [modifiedFlowerLatinName, setModifiedFlowerLatinName] = useState(flower.latin_name)
  const [modifiedFlowerQty, setModifiedFlowerQty] = useState(flower.quantity)
  const { t, i18n } = useTranslation()
  
  const updateFlower = event => {
    event.preventDefault()
    const newFlower = {
      ...flower, 
      name: modifiedFlowerName, 
      latin_name: modifiedFlowerLatinName, 
      quantity: Number(modifiedFlowerQty),
    }
    modifyFlower(newFlower)
    setModifiedFlowerName(modifiedFlowerName)
    setModifiedFlowerLatinName(modifiedFlowerLatinName)
    setModifiedFlowerQty(modifiedFlowerQty)
    handleFlowerModify(newFlower)
  }
    return (
      <Container>
        <form onSubmit={updateFlower}>
          <div className="form-group">
            <label htmlFor="modifiedFlowerNameInput">{t("flower.data.name")}:</label>
            <input
              id="modifiedFlowerNameInput"
              value={modifiedFlowerName}
              onChange={event => setModifiedFlowerName(event.target.value)}
              className="form-control"
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="modifiedFlowerLatinNameInput">{t("flower.data.latinname")}:</label>
            <input
              id="modifiedFlowerLatinNameInput"
              value={modifiedFlowerLatinName}
              onChange={event => setModifiedFlowerLatinName(event.target.value)}
              className="form-control"
            />
          </div>
          <div className="form-group">
            <label htmlFor="modifiedFlowerQtyInput">{t("flower.data.qty")}:</label>
            <input
              type="number"
              id="modifiedFlowerQtyInput"
              value={modifiedFlowerQty}
              onChange={event => setModifiedFlowerQty(event.target.value)}
              className="form-control"
              min="0"
              max="1000000"
              required
            />
          </div>
          <div>
            <Button variant="light" id="saveModifiedFlowerButton" type="submit">
              {t("button.save")}
            </Button>
            <Button variant="dark" className="custom-button" onClick={handleFormVisibility}>
              {t('button.cancel')}
            </Button>
          </div>

        </form>
      </Container>
    )
}

export default ModifyFlowerForm
