import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import ImageService from '../services/images'
import NewSiteForm from './NewSiteForm'
import placeholderImage from '../assets/images/site-placeholder-image.jpg'
import '../Misc.css'

const SiteFlexbox = ({ createSite, sites }) => {
  const [showAddNewSite, setShowAddNewSite] = useState(false)
  const [images, setImages] = useState([])
  const { t, i18n } = useTranslation()

  useEffect(() => {
    if (sites) {
    const newImages = Promise.all(sites.map((site) => {
      if (site.favorite_image) {
        return ImageService.getByID(site.favorite_image)
          .then((url) => (
            {site: site._id, url: url}
          ))
          .catch((_) => {})
      }
    }))

    newImages.then((imgs) => setImages(imgs.filter((x)=>x)))
    }
  }, [sites])
    
  return (
    <div className='flexbox'>
      <div className='flexGap'>
        {sites &&
          sites.map(site => (
            <div className="box" key={site._id}>
              <div className='image-flexbox-container'>
                {images.find((o) => o.site === site._id)?.url && 
                <img src={images.find((o) => o.site === site._id)?.url} alt={site.name} />
                }
              </div>
              <h3><Link to={`/grower/${site._id}`} className="link-success">{site.name}</Link></h3>
            </div>
          ))}
        <div className='box'>
          <button id="addNewSiteButton" onClick={() => setShowAddNewSite(!showAddNewSite)} className='btn btn-light'>
            {t('button.addsite')}
          </button>
          {showAddNewSite && <NewSiteForm createSite={createSite} />}
        </div>
      </div>
    </div>
  )
}

export default SiteFlexbox