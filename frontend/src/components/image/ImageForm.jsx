import { useState } from "react"
import { useTranslation } from "react-i18next"

const ImageForm = ({ createImage }) => {
  const [newImage, setNewImage] = useState(null)
  const [newImageNote, setNewImageNote] = useState("")
  const { t, i18n } = useTranslation()

  const handleSubmit = event => {
    event.preventDefault()

    createImage({
      note: newImageNote,
      image: newImage
    })

    setNewImage(null)
    setNewImageNote("")

    document.getElementById("image-form").reset()
  }

  const handleFileSelect = (event) => {
    setNewImage(event.target.files[0])
  }

  return (
    <div className="text-left">
      <form onSubmit={handleSubmit} id="image-form">
        <div className="form-group">
          <label htmlFor="newImage">{t("image.select")}:</label>
          <input
            id="newImageInput"
            className="form-control"
            type="file"
            accept="image/*"
            required={true}
            onChange={handleFileSelect}
          />
        </div>
        <div className="form-group">
          <label htmlFor="newImageNoteInput">{t("image.note")}:</label>
          <input
            id="newImageNoteInput"
            className="form-control"
            value={newImageNote}
            required={true}
            onChange={event => setNewImageNote(event.target.value)}
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
