import { useState } from 'react'
import NewSiteForm from './NewSiteForm'

const SiteFlexbox = ( createSite) => {
  const [showAddNewSite, setShowAddNewSite] = useState(false)

  const containerStyle = {
    display: 'flex',
    flexDirection: 'row',
    flexWrap: 'wrap',
    width: '100%',
    height: '100%',
    border: '5px solid red'
  }

  const boxStyle = {
    textAlign: 'center',
    border: '2px solid #ccc',
    padding: '20px',
    width: '175px',
    height: '190px',
  }

  return (
    <div style={containerStyle}>
      <div style={boxStyle}>1</div>
      <div style={boxStyle}>2</div>
      <div style={boxStyle}>3</div>
      <div style={boxStyle}>4</div>
      <div style={boxStyle}>
        <button id="addNewSiteButton" onClick={() => setShowAddNewSite(!showAddNewSite)}>Add new site</button>
        {showAddNewSite && <NewSiteForm createSite={createSite} />}
      </div>
    </div>
    )
}

export default SiteFlexbox
