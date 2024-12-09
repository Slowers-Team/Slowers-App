import { useState } from "react"
import { useTranslation } from "react-i18next"

const ModifyFlowerForm = ({ flower, modifyFlower, handleFlowerModify, handleFormVisibility, addedTime}) => {
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
          <table className="table table">
            <tbody>
              <tr>
                <th>{t('flower.data.name')}</th>
                <td>
                  <input
                    id="modifiedFlowerNameInput"
                    value={modifiedFlowerName}
                    onChange={event => setModifiedFlowerName(event.target.value)}
                    className="form-control"
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
        </form>
      </div>
    )
}

export default ModifyFlowerForm
