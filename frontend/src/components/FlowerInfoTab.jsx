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
              <p>{t('flower.data.name')}: {flower.name}</p>
              <p>{t('flower.data.latinname')}: {flower.latin_name}</p>
              <p>{t('flower.data.addedtime')}: {addedTime(flower)}</p>
              <p>{t('flower.data.site')}: {flower.site_name}</p>
              <p>{t('flower.data.qty')}: {flower.quantity}</p>
            </div>
          )}
        {isGrower ?
        <p>{t('flower.visible.long')}: {flower.visible 
              ? t('flower.visible.true') 
              : t('flower.visible.false')}
        <VisibilityButton flower={flower} updateFlower={updateFlower}/>
        </p> : <></>}
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
