import { useState } from 'react'

const SiteForm = ({ createSite }) => {
  const [newSiteName, setNewSiteName] = useState('')
  const [newSiteNote, setNewSiteNote] = useState('')

  const addSite = event => {
    event.preventDefault()
    createSite({
      name: newSiteName,
      note: newSiteNote
    })

    setNewSiteName('')
    setNewSiteLocation('')
  }

  return (
    <div>
      <form onSubmit={addSite}>
        <div>
          <label htmlFor="newSiteNameInput">Name:</label>
          <input id="newSiteNameInput" value={newSiteName} onChange={event => setNewSiteName(event.target.value)} />
        </div>
        <div>
          <label htmlFor="newSiteNoteInput">Note:</label>
          <input id="newSiteNoteInput" value={newSiteNote} onChange={event => setNewSiteNote(event.target.value)} />
        </div>
        <div>
          <button id="saveNewSiteButton" type="submit">Save</button>
        </div>
      </form>
    </div>
  )
}

export default SiteForm
