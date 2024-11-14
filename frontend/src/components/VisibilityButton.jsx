import flowerService from '../services/flowers'
import Button from 'react-bootstrap/ToggleButton'
import { useState } from 'react' 

const VisibilityButton = ({ flower }) => {
  const [current, setCurrent] = useState(flower.visible)

  const handleClick = () => {
    flowerService.toggleVisibility(flower.id).then(setCurrent(!current))
  }
  
  return (
    <Button onClick={handleClick}>
      {current
        ? "Hide from retailers"
        : "Show  to retailers"}
    </Button>
  )
}

export default VisibilityButton
