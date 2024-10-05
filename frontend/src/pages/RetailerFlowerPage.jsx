import RetailerFlowerList from '../components/retailer/RetailerFlowerList'
import flowerService from '../services/flowers'
import { useEffect, useState } from 'react'
const RetailerFlowerPage = () => {
  const [flowers, setFlowers] = useState([])

  useEffect(() => {
    flowerService.getAll().then(initialFlowers => setFlowers(initialFlowers))
  }, [])

  return (
    <>
      <h2>Flowers</h2>
      {flowers && <RetailerFlowerList flowers={flowers} />}
    </>
  )
}

export default RetailerFlowerPage
