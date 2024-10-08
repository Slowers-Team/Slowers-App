import { useState, useEffect } from "react"
import flowerService from "../services/flowers"
import FlowerList from "../components/FlowerList"

const HomePage = () => {
  const [flowers, setFlowers] = useState([])

  useEffect(() => {
    flowerService.getAll().then(initialFlowers => setFlowers(initialFlowers))
  }, [])

  const deleteFlower = flowerObject => {
    if (
      window.confirm(
        `Are you sure you want to delete flower ${flowerObject.name}?`
      )
    ) {
      flowerService.remove(flowerObject._id).then(response => {
        console.log(response)
        setFlowers(l => l.filter(item => item._id !== flowerObject._id))
      })
    }
  }

  return (
    <div>
      <h2>Flowers</h2>
      {flowers && <FlowerList flowers={flowers} deleteFlower={deleteFlower} />}
    </div>
  )
}

export default HomePage
