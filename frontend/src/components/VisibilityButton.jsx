import flowerService from '../services/flowers'
import Button from 'react-bootstrap/ToggleButton'

const VisibilityButton = ({id}) => {
  return (
    <Button onClick={() => flowerService.toggleVisibility(id)}>
      Toggle visibility
    </Button>
  )
}

export default VisibilityButton
