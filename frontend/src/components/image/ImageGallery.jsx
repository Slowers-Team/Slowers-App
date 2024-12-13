import { useState, useEffect } from "react"
import { useTranslation } from "react-i18next"
import { Button } from "react-bootstrap"
import Masonry from "react-masonry-css"
import './ImageGallery.css'

const ImageGallery = ({ isGrower, images, deleteImage, favoriteImage, type }) => {
  const { t } = useTranslation() 
	const [selectedFavoriteID, setSelectedFavoriteID] = useState(null)

	useEffect(()=> {
		const favID = images.find((img) => img?.favorite)?._id
		setSelectedFavoriteID(favID)
	}, [images])
	
	const handleFavoriteSelect = (imageObject) => {
		favoriteImage(imageObject)
	}

	let breakpointColumnsObj

	if (type === "flower") {
		breakpointColumnsObj = {default: 2, 991: 1,}
	} else {
		breakpointColumnsObj = {default: 4, 1500: 3, 950: 2, 550: 1,}
	}

  return (
    <div className="my-2">
			{(!images || images.length === 0) ? (
        <p>
				{type === "flower"
					? t('image.noflowerimages')
					: t('image.nositeimages')}
				</p>
			) : (
				<div>
					<Masonry breakpointCols={breakpointColumnsObj} className="my-masonry-grid" columnClassName="my-masonry-grid_column">
						{images.map((image) => (
						<div className="image-box" key={image._id}>
							<img src={image.url}/>
							{isGrower && (
							<div className="image-buttons">
								<Button variant="dark" onClick={() => deleteImage(image)} className="delete-button" aria-label="Delete">
									<i className="bi bi-trash"></i>
								</Button>
								<Button variant="dark"
									      onClick={() => handleFavoriteSelect(image._id)}
									      className={`favourite-button ${selectedFavoriteID === image._id ? "selected" : ""}`} 
									      disabled={selectedFavoriteID !== null && selectedFavoriteID == image._id} aria-label="Favorite">
									<i className={`bi bi-star-fill ${selectedFavoriteID === image._id ? "text-warning" : ""}`}></i>
								</Button>
							</div>
							)}
						</div>
						))}
					</Masonry>
				</div>
			)}
    </div>
  )
}

export default ImageGallery
