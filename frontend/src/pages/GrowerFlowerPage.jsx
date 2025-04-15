import { useParams } from 'react-router-dom'
import GrowerFlowerList from '../components/grower/GrowerFlowerList'
import flowerService from '../services/flowers'
import siteService from '../services/sites'
import AddFlower from '../components/grower/AddFlower'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'

const GrowerFlowerPage = () => {
  const params = useParams()
  const [flowers, setFlowers] = useState()
  const [checkedFlowers, setCheckedFlowers] = useState([])
  const [searchTerm, setSearchTerm] = useState("")
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
    <div className="m-3">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-12">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              {params.siteId ? (
                <h2>{site?.name} {t('title.siteflowers')}</h2>
                ) : (
                  <h2>{t('title.allflowers')}</h2>
                )}
              <div className="d-flex gap-2 mt-3">
                <div className="d-flex justify-content-start input-wrapper">
                  <input
                    type="text"
                    placeholder={t('button.Search')}
                    value={searchTerm}
                    onChange={(e) => setSearchTerm(e.target.value)}
                  />
                </div>
                {params.siteId && <AddFlower createFlower={addFlower} siteID={params.siteId} />}
                <button className="custom-delete-button" onClick={() => deleteMultipleFlowers(checkedFlowers)}>
                  <i className="bi bi-trash3-fill"> </i>
                  {t('button.delete')}
                </button>
              </div>
              { flowers ? (<GrowerFlowerList flowers={flowers} deleteFlower={deleteFlower} modifyFlower={modifyFlower} setCheckedFlowers={setCheckedFlowers} updateFlower={updateFlower} searchTerm={searchTerm}/>) : 
                          (<GrowerFlowerList flowers={[]} deleteFlower={deleteFlower} modifyFlower={modifyFlower} setCheckedFlowers={setCheckedFlowers} updateFlower={updateFlower} searchTerm={searchTerm}/>) }
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default GrowerFlowerPage
