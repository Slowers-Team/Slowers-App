import { useState } from "react"
import { useTranslation } from "react-i18next"

const ImageForm = ({ createImage }) => {
  const [newImage, setNewImage] = useState()
  const [preview, setPreview] = useState()
  const [newImageNote, setNewImageNote] = useState("")
  const { t, i18n } = useTranslation()

  const handleSubmit = event => {
    event.preventDefault()

    createImage({
      note: newImageNote,
      image: newImage
    })

    setNewImage(undefined)
    setNewImageNote("")

    document.getElementById("image-form").reset()
  }

  
  const handleFileSelect = (event) => {
    if (!event.target.files || event.target.files.length === 0) {
      setSelectedFile(undefined)
      return
    }
    const file = event.target.files[0]
    setNewImage(file)
    // const objectUrl = URL.createObjectURL(file)

    const reader = new FileReader()
    reader.onload = (e) => {
      setPreview(e.target.result)
    }
    reader.readAsDataURL(file)

    // setPreview(objectUrl)
    // setNewImage(file)

    // return () => URL.revokeObjectURL(objectUrl)
  }

  return (
    <div className="text-left mx-4">
      <form onSubmit={handleSubmit} id="image-form">
        <div className="form-group mt-2">
          <label htmlFor="newImageInput">{t("image.select")}:</label>
          <input
            id="newImageInput"
            className="form-control"
            type="file"
            accept="image/jpg,image/png,image/jpeg"
            required
            onChange={handleFileSelect}
          />
        </div>
        <div className="form-group mt-2">
          <label htmlFor="newImageNoteInput">{t("image.note")}:</label>
          <input
            id="newImageNoteInput"
            className="form-control"
            value={newImageNote}
            required
            onChange={event => setNewImageNote(event.target.value)}
          />
        </div>
        <div>
          <button id="saveNewImageButton" type="submit" className="custom-button my-3">
            <i className="bi bi-floppy2-fill"> </i>
            {t("button.save")}
          </button>
        </div>
      </form>
      {newImage && <img width={100} src={preview} alt="Preview" />}
    </div>
  )
}

export default ImageForm
