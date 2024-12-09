import { useParams } from 'react-router-dom'
import GrowerFlowerList from '../components/grower/GrowerFlowerList'
import flowerService from '../services/flowers'
import siteService from '../services/sites'
import AddFlower from '../components/grower/AddFlower'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Button } from 'react-bootstrap'

const GrowerFlowerPage = () => {
  const params = useParams()
  const [flowers, setFlowers] = useState()
  const [checkedFlowers, setCheckedFlowers] = useState([])
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

  const deleteMultipleFlowers = checkedFlowers => {
    if (checkedFlowers.length === 0) {
      alert(t('label.noflowersselected'))
      return
    }
    if (window.confirm(t('label.confirmmultipleflowerdeletion'))) {
      flowerService.removeMultipleFlowers(checkedFlowers).then(response => {
        console.log(response)
        setFlowers(l => l.filter(item => !checkedFlowers.includes(item._id)))
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

  const modifyFlower = flowerObject => {
    flowerService
      .modify(flowerObject)
      .then(updateFlower(flowerObject))
      .catch(error => {
        console.log(error)
        alert(t("error.modifyingfailed"))
      })
  }

  return (
    <>
    {params.siteId ? (
      <div>
        <h2>{site?.name} {t('title.siteflowers')}</h2>
        <AddFlower createFlower={addFlower} siteID={params.siteId} />
        <Button variant="light" onClick={() => deleteMultipleFlowers(checkedFlowers)}>
          <i className="bi bi-trash3-fill"></i>
        </Button>
      </div>
    ) : (
      <div>
        <h2>{t('title.allflowers')}</h2>
        <Button variant="light" onClick={() => deleteMultipleFlowers(checkedFlowers)}>
          <i className="bi bi-trash3-fill"></i>
        </Button>
      </div>
    )}
      { flowers ? (<GrowerFlowerList flowers={flowers} deleteFlower={deleteFlower} modifyFlower={modifyFlower} setCheckedFlowers={setCheckedFlowers} updateFlower={updateFlower}/>) : 
                  (<GrowerFlowerList flowers={[]} deleteFlower={deleteFlower} modifyFlower={modifyFlower} setCheckedFlowers={setCheckedFlowers} updateFlower={updateFlower}/>) }
    </>
  )
}

export default GrowerFlowerPage
