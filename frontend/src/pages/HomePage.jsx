import { useState, useEffect } from 'react'
import flowerService from '../services/flowers'
import FlowerForm from '../components/FlowerForm'
import FlowerList from '../components/FlowerList'

const HomePage = () => {
    const [flowers, setFlowers] = useState([])
    const [showAddNewFlower, setShowAddNewFlower] = useState(false)
  
    useEffect(() => {
      flowerService
        .getAll()
        .then(initialFlowers => setFlowers(initialFlowers))
    }, [])
  
    const addFlower = flowerObject => {
      flowerService
        .create(flowerObject)
        .then(returnedFlower => setFlowers(flowers.concat(returnedFlower)))
        .catch(error => {
          console.log(error)
          alert(`Adding failed`)
        })
    }
  
    const deleteFlower = flowerObject => {
      if (window.confirm(`Are you sure you want to delete flower ${flowerObject.name}?`)) {
        flowerService
          .remove(flowerObject._id)
          .then(response => {
            console.log(response)
            setFlowers(l => l.filter(item => item.name !== flowerObject.name));
          })
      }
    }

  return (
    <div>
      <button id="showFlowerAddingFormButton" onClick={() => setShowAddNewFlower(!showAddNewFlower)}>Add a new flower</button>
      {showAddNewFlower && <FlowerForm createFlower={addFlower}/>}
      {flowers && <FlowerList flowers={flowers} deleteFlower={deleteFlower} />}
    </div>
  )
}

export default HomePage