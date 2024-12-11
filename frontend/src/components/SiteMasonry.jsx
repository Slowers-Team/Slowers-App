import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import Masonry from 'react-masonry-css'
import ImageService from '../services/images'
import '../Misc.css'

const SiteMasonry = ({ sites }) => {
  const [images, setImages] = useState([])

  const breakpointColumnsObj = {
    default: 5,
    1700: 4,
    1300: 3,
    900: 2,
    500: 1,
  }

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
    <Masonry breakpointCols={breakpointColumnsObj} className="my-masonry-grid" columnClassName="my-masonry-grid_column">
      {sites &&
        sites.map(site => (
          <div className="masonry-box" key={site._id}>
            <div className='image-flexbox-container'>
              {images.find((o) => o.site === site._id)?.url && 
              <img src={images.find((o) => o.site === site._id)?.url} alt={site.name} />
              }
            </div>
            <h3 className="my-2"><Link to={`/grower/${site._id}`} className="link-success">{site.name}</Link></h3>
            <p>{site.note}</p>
          </div>
        ))}
    </Masonry>
  )
}

export default SiteMasonry