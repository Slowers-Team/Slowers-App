import flowerService from '../services/flowers'
import { useState } from 'react' 
import { useTranslation } from 'react-i18next'

const VisibilityButton = ({ flower, updateFlower }) => {
  const [current, setCurrent] = useState(flower.visible)
  const [disabled, setDisabled] = useState(false)
  const {t, _ } = useTranslation()

  const handleClick = () => {
    flowerService.toggleVisibility(flower._id)
    .then((_)=> {
      const newVis = !current
      const newFlower = {...flower, visible: newVis}
      setCurrent(newVis)
      updateFlower(newFlower)
    })
    .catch(error => {
      if (error.startsWith("No image")) {
        setDisabled(true)
        alert(t("error.noimagesinflower"))
      } else {
       alert(error) 
      }
    })
  }
  
  return (
    <button onClick={handleClick} disabled={disabled} id="visibilityButton" className="flower-button">
      {current
        ? t("button.hideFlower")
        : t("button.showFlower")}
    </button>
  )
}

export default VisibilityButton
