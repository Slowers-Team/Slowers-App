import GrowerFlowerList from '../components/grower/GrowerFlowerList'
import flowerService from '../services/flowers'
import { useEffect, useState } from 'react'
const GrowerFlowerPage = () => {
  const [flowers, setFlowers] = useState([])

  useEffect(() => {
    flowerService.getUserFlowers().then(initialFlowers => setFlowers(initialFlowers))
  }, [])

  const deleteFlower = flowerObject => {
    if (window.confirm(`Are you sure you want to delete flower ${flowerObject.name}?`)) {
      flowerService.remove(flowerObject._id).then(response => {
        console.log(response)
        setFlowers(l => l.filter(item => item._id !== flowerObject._id))
      })
    }
  }

  return (
    <>
      <h2>Flowers</h2>
      {flowers && <GrowerFlowerList flowers={flowers} deleteFlower={deleteFlower} />}
    </>
  )
}

export default GrowerFlowerPage
