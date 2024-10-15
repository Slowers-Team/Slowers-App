import { useState } from 'react'
import { useParams } from 'react-router-dom'
import { useTranslation } from 'react-i18next'

const SiteForm = ({ createSite }) => {
  const { id: parentId } = useParams()
  const [newSiteName, setNewSiteName] = useState('')
  const [newSiteNote, setNewSiteNote] = useState('')
  const { t, i18n } = useTranslation()

  const addSite = event => {
    event.preventDefault()
    createSite({
      name: newSiteName,
      note: newSiteNote,
      ...(parentId && { parent: parentId }),
    })

    
    setNewSiteName('')
    setNewSiteNote('')
  }

  return (
    <div>
      <form onSubmit={addSite}>
        <div>
          <label htmlFor="newSiteNameInput">{t("site.data.name")}:</label>
          <input id="newSiteNameInput" value={newSiteName} onChange={event => setNewSiteName(event.target.value)} />
        </div>
        <div>
          <label htmlFor="newSiteNoteInput">{t("site.data.note")}:</label>
          <input id="newSiteNoteInput" value={newSiteNote} onChange={event => setNewSiteNote(event.target.value)} />
        </div>
        <div>
          <button id="saveNewSiteButton" type="submit">{t("button.save")}</button>
        </div>
      </form>
    </div>
  )
}

export default SiteForm
