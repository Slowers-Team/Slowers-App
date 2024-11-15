import flowerService from '../services/flowers'
import Button from 'react-bootstrap/ToggleButton'
import { useState } from 'react' 
import { useTranslation } from 'react-i18next'

const VisibilityButton = ({ flower, updateFlower }) => {
  const [current, setCurrent] = useState(flower.visible)
  const {t, _ } = useTranslation()

  const handleClick = () => {
    flowerService.toggleVisibility(flower._id)
    .then((_)=> {
      const newVis = !current
      const newFlower = {...flower, visible: newVis}
      setCurrent(newVis)
      updateFlower(newFlower)
    })
    .catch(error => alert(error))
  }
  
  return (
    <Button onClick={handleClick}>
      {current
        ? t("button.hideFlower")
        : t("button.showFlower")}
    </Button>
  )
}

export default VisibilityButton
