import { useParams, useRouteLoaderData, useLoaderData } from 'react-router-dom'
import GrowerFlowerList from '../components/grower/GrowerFlowerList'
import flowerService from '../services/flowers'
import AddFlower from '../components/grower/AddFlower'
import AddFlowerUpdate from '../components/grower/AddFlowerUpdate'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'

const GrowerFlowerPage = () => {
  const params = useParams()
  const [flowers, setFlowers] = useState()
  const [checkedFlowers, setCheckedFlowers] = useState([])
  const { t, i18n } = useTranslation()
  const site = useRouteLoaderData("site")?.site

  const data = useLoaderData()
  console.log("data", data)


  useEffect(() => {
    if (params.siteId) {
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
    {site ? (
      <div>
        <h2>{site?.name} {t('title.siteflowers')}</h2>
        <AddFlower createFlower={addFlower} siteID={site._id} />
        <AddFlowerUpdate checkedFlowers={checkedFlowers} />
      </div>
    ) : (
      <div>
        <h2>{t('title.allflowers')}</h2>
        <AddFlowerUpdate checkedFlowers={checkedFlowers} />
      </div>
    )}
      { flowers ? (<GrowerFlowerList flowers={flowers} deleteFlower={deleteFlower} setCheckedFlowers={setCheckedFlowers} updateFlower={updateFlower}/>) : 
                  (<GrowerFlowerList flowers={[]} deleteFlower={deleteFlower} setCheckedFlowers={setCheckedFlowers} updateFlower={updateFlower}/>) }
    </>
  )
}

export default GrowerFlowerPage
