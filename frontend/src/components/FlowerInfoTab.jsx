import { useTranslation } from "react-i18next"
import { useState } from "react"
import VisibilitySwitch from './VisibilitySwitch'
import ModifyFlowerForm from './ModifyFlowerForm'
import "./FlowerModal.css"
import { formatTime } from "../utils.js"

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

	return (
		<div>
			{isGrower && isModifyFormVisible ? (
					<div>
						<ModifyFlowerForm flower={flower} modifyFlower={modifyFlower} handleFlowerModify={updateFlower} handleFormVisibility={handleFormVisibility} handleFlowerDelete={handleFlowerDelete} addedTime={formatTime(flower.added_time)}/>
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
								<td>
									<em>{flower.latin_name}</em>
								</td>
							</tr>
							<tr>
								<th>{t('flower.data.addedtime')}</th>
								<td>{formatTime(flower.added_time)}</td>
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
