import { useState } from "react"
import { useTranslation } from "react-i18next"

const ImageForm = ({ createImage, entityID }) => {
  const [newImage, setNewImage] = useState(null)
  const [newImageNote, setNewImageNote] = useState("")
  const { t, i18n } = useTranslation()

  const addImage = event => {
    if (!newImage) {
      
    }
    event.preventDefault()
    createImage({
      note: newImageNote,
      entity: entityID,
      image: newImage
    })

    setNewImage(null)
    setNewImageNote("")
  }

  const handleFileSelect = (event) => {
    setNewImage(event.target.files[0])
  }

  return (
    <div className="text-left">
      <form onSubmit={addImage}>
        <div className="form-group">
          <label htmlFor="newImage">{t("image.select")}:</label>
          <input
            id="newImageInput"
            type="file"
            accept="image/*"
            onChange={handleFileSelect}
            className="form-control"
          />
        </div>
        <div className="form-group">
          <label htmlFor="newImageNoteInput">{t("image.note")}:</label>
          <input
            id="newImageNoteInput"
            value={newImageNote}
            onChange={event => setNewImageNote(event.target.value)}
            className="form-control"
          />
        </div>
        <div>
          <button id="saveNewImageButton" type="submit" className="btn btn-light">
            {t("button.save")}
          </button>
        </div>
      </form>
    </div>
  )
}

export default ImageForm
