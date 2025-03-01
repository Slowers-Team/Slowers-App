import { useState } from "react"
import { useTranslation } from "react-i18next"
import { Button } from "react-bootstrap"

const ModifyFlowerForm = ({ flower, modifyFlower, handleFlowerModify, handleFormVisibility, handleFlowerDelete, addedTime}) => {
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
      <div>
        <form onSubmit={updateFlower}>
          <table className="table custom-table">
            <tbody>
              <tr>
                <th>{t('flower.data.name')}</th>
                <td>
                  <input
                    id="modifiedFlowerNameInput"
                    value={modifiedFlowerName}
                    onChange={event => setModifiedFlowerName(event.target.value)}
                    className="form-control"
                    aria-label="Name"
                    required
                  />
                </td>
              </tr>
              <tr>
                <th>{t('flower.data.latinname')}</th>
                <td>
                  <input
                    id="modifiedFlowerLatinNameInput"
                    value={modifiedFlowerLatinName}
                    onChange={event => setModifiedFlowerLatinName(event.target.value)}
                    className="form-control"
                    aria-label="Latin name"
                  />
                </td>
              </tr>
              <tr>
								<th>{t('flower.data.addedtime')}</th>
                <td>{addedTime}</td>
              </tr>
              <tr>
                <th>{t('flower.data.site')}</th>
                <td>{flower.site_name}</td>
              </tr>
              <tr>
                <th>{t('flower.data.qty')}</th>
                <td>
                  <input
                    type="number"
                    id="modifiedFlowerQtyInput"
                    value={modifiedFlowerQty}
                    onChange={event => setModifiedFlowerQty(event.target.value)}
                    className="form-control"
                    aria-label="Quantity"
                    min="0"
                    max="1000000"
                    required
                  />
                </td>
              </tr>
              <tr>
              <th>{t('flower.visible.long')}</th>
              <td>
                <div>
                  {flower.visible 
                      ? t('flower.visible.true') 
                      : t('flower.visible.false')}
                </div>
              </td>
            </tr>
            </tbody>
          </table>
            <button className="custom-button" id="deleteFlowerButton" onClick={() => handleFlowerDelete(flower)} type="button">
            <i className="bi bi-trash3-fill"> </i>
              {t('button.delete')}
            </button>
            <button variant="light" className="custom-button" id="saveModifiedFlowerButton" type="submit">
              <i className="bi bi-floppy2-fill"> </i>
              {t("button.save")}
            </button>
            <button variant="dark" className="custom-button" id="modifyFlowerCancelButton" onClick={handleFormVisibility}>
              <i className="bi bi-x-lg"> </i>
              {t("button.cancel")}
            </button>
        </form>
      </div>
    )
}

export default ModifyFlowerForm
