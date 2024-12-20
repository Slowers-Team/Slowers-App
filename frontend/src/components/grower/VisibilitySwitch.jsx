import flowerService from '../../services/flowers'
import { useState } from 'react' 
import { useTranslation } from 'react-i18next'
import Form from 'react-bootstrap/Form'

const VisibilitySwitch = ({ flower, updateFlower }) => {
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
    <Form>
      <Form.Check type="switch" id="custom-switch" className='custom-switch' checked={current} onChange={handleClick} />
    </Form>
  )
}

export default VisibilitySwitch
