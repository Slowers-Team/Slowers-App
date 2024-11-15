import { useParams } from 'react-router-dom'
import GrowerFlowerList from '../components/grower/GrowerFlowerList'
import flowerService from '../services/flowers'
import siteService from '../services/sites'
import FlowerForm from '../components/FlowerForm'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
const GrowerFlowerPage = () => {
  const params = useParams()
  const [flowers, setFlowers] = useState()
  const [showAddNewFlower, setShowAddNewFlower] = useState(false)
  const [site, setSite] = useState()
  const { t, i18n } = useTranslation()

  useEffect(() => {
    if (params.siteId) {
      siteService.get(params.siteId).then(initialSite => setSite(initialSite.site))
      flowerService
          .getFlowersBySite(params.siteId)
          .then(flowers => {
            setFlowers(flowers)
          })
          .catch(error => {
            console.error("Error fetching flowers:", error)
          })
    } else {
      flowerService.getUserFlowers().then(initialFlowers => setFlowers(initialFlowers))
    }
  }, [params.siteId])

  const addFlower = flowerObject => {
    flowerService
      .create(flowerObject)
      .then(returnedFlower => 
        setFlowers(flowers ? flowers.concat(returnedFlower) :
          [returnedFlower])
      )
      .catch(error => {
      console.log(error)
      alert(t("error.addingfailed"))
    })
  }

  const deleteFlower = flowerObject => {
    if (window.confirm(`${t('label.confirmflowerdeletion')} ${flowerObject.name}?`)) {
      flowerService.remove(flowerObject._id).then(response => {
        console.log(response)
        setFlowers(l => l.filter(item => item._id !== flowerObject._id))
      })
    }
  }

  const updateFlower = flowerObject => {
    setFlowers(flowers.map((flower) => 
      flower._id === flowerObject._id 
        ? flowerObject 
        : flower
    ))
  }

  return (
    <>
    {params.siteId ? (
      <div>
        <h2>{site?.name} {t('title.siteflowers')}</h2>
        <button id="showFlowerAddingFormButton" onClick={() => setShowAddNewFlower(!showAddNewFlower)} className="btn btn-light">
          {t("button.addflower")}
        </button>
        {showAddNewFlower && <FlowerForm createFlower={addFlower} siteID={params.siteId} />}
      </div>
    ) : (
      <h2>{t('title.allflowers')}</h2>
    )}
      { flowers ? (<GrowerFlowerList flowers={flowers} deleteFlower={deleteFlower} updateFlower={updateFlower}/>) : (<GrowerFlowerList flowers={[]} deleteFlower={deleteFlower} />) }
    </>
  )
}

export default GrowerFlowerPage
