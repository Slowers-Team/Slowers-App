import GrowerFlowerList from '../components/grower/GrowerFlowerList'
import flowerService from '../services/flowers'
import { useEffect, useState } from 'react'
const GrowerFlowerPage = () => {
  const [flowers, setFlowers] = useState([])

  useEffect(() => {
    flowerService.getUserFlowers().then(initialFlowers => setFlowers(initialFlowers))
  }, [])

  return (
    <>
      <h2>Flowers</h2>
      {flowers && <GrowerFlowerList flowers={flowers} />}
    </>
  )
}

export default GrowerFlowerPage
