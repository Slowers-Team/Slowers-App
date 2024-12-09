import { useTranslation } from "react-i18next"
import { useState } from "react"
import VisibilityButton from './VisibilityButton'
import ModifyFlowerForm from './ModifyFlowerForm'

const FlowerInfoTab = ({isGrower, flower, deleteFlower, updateFlower, modifyFlower, handleClose}) => {
    const [isModifyFormVisible, setIsModifyFormVisible] = useState(false)
    const { t } = useTranslation()

    const handleFormVisibility = () => {
      setIsModifyFormVisible((prev) => !prev)
    }

		const handleFlowerDelete = (flower) => {
			if (deleteFlower) {
				deleteFlower(flower)
			}
			handleClose()
		}

		const addedTime = (flower) => {
			let addedTime = new Date(flower.added_time)
	
			let date = addedTime.toLocaleDateString('fi')
			let time = addedTime.toLocaleTimeString('fi', { hour: '2-digit', minute: '2-digit' })
			let addedTimeStr = `${date} ${time}`
	
			return addedTimeStr
		}
    return (
      <div>
				<h3>{t('menu.info')}</h3>
        {isGrower && isModifyFormVisible ? (
            <div>
              <ModifyFlowerForm flower={flower} modifyFlower={modifyFlower} handleFlowerModify={updateFlower} handleFormVisibility={handleFormVisibility}/>
            </div> 
          ) : (
						<div>
						<table className="table">
							<tbody>
								<tr>
									<th>{t('flower.data.name')}</th>
									<td>{flower.name}</td>
								</tr>
								<tr>
									<th>{t('flower.data.latinname')}</th>
									<td>{flower.latin_name}</td>
								</tr>
								<tr>
									<th>{t('flower.data.addedtime')}</th>
									<td>{addedTime(flower)}</td>
								</tr>
								<tr>
									<th>{t('flower.data.site')}</th>
									<td>{flower.site_name}</td>
								</tr>
								<tr>
									<th>{t('flower.data.qty')}</th>
									<td>{flower.quantity}</td>
								</tr>
								<tr>
									<th>{t('flower.visible.long')}</th>
									<td>
										{isGrower && (
											<VisibilityButton flower={flower} updateFlower={updateFlower} visible={flower.visible}/>
										)}
									</td>
								</tr>
							</tbody>
						</table>
					</div>
				)}
        {deleteFlower && (
          <button id="deleteFlowerButton" onClick={() => handleFlowerDelete(flower)}>
            {t('button.delete')}
          </button>
        )}
        {isGrower && !isModifyFormVisible && (
          <button id="modifyFlowerButton" onClick={handleFormVisibility}>
            {t('button.modify')}
          </button>
        )}
    	</div>
    )
}

export default FlowerInfoTab
