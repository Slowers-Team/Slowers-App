import { useTranslation } from "react-i18next"
import { useState } from "react"
import VisibilitySwitch from './VisibilitySwitch'
import ModifyFlowerForm from './ModifyFlowerForm'
import "./FlowerModal.css"

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
			{isGrower && isModifyFormVisible ? (
					<div>
						<ModifyFlowerForm flower={flower} modifyFlower={modifyFlower} handleFlowerModify={updateFlower} handleFormVisibility={handleFormVisibility} handleFlowerDelete={handleFlowerDelete} addedTime={addedTime(flower)}/>
					</div> 
				) : (
					<div>
					<table className="table custom-table">
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
							{isGrower && (
							<tr>
								<th>{t('flower.data.site')}</th>
								<td>{flower.site_name}</td>
							</tr>
							)}
							{!isGrower && (
							<tr>
								<th>{t('flower.data.grower')}</th>
								<td>{flower.grower_email}</td>
							</tr>
							)}
							<tr>
								<th>{t('flower.data.qty')}</th>
								<td>{flower.quantity}</td>
							</tr>
							{isGrower && (
								<tr>
									<th>{t('flower.visible.long')}</th>
									<td>
										<VisibilitySwitch flower={flower} updateFlower={updateFlower} visible={flower.visible}/>
									</td>
								</tr>
							)}
						</tbody>
					</table>
				</div>
			)}
			{isGrower && !isModifyFormVisible && (
				<div>
					<button className="custom-button" id="deleteFlowerButton" onClick={() => handleFlowerDelete(flower)}>
						<i className="bi bi-trash3-fill"> </i>
						{t('button.delete')}
					</button>
					<button className="custom-button" id="modifyFlowerButton" onClick={handleFormVisibility}>
						<i className="bi bi-pencil-fill"> </i>
						{t('button.modify')}
					</button>
				</div>
			)}
		</div>
	)
}

export default FlowerInfoTab
