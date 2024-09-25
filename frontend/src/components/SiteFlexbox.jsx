import { useState } from 'react'

const SiteFlexbox = () => {
    const containerStyle = {
        display: 'flex',
        flexDirection: 'row',
        flexWrap: 'wrap',
        justifyContent: 'spacebetween',
        width: '100%',
        height: '100%',
        border: '5px solid red'
    }

    const boxStyle = {
        width: '50px',
        height: '100px',
        colour: 'black',
        padding: '10px',
        border: '5px solid green'
    }

    return (
        <div style={containerStyle}>
            <div style={boxStyle}>1</div>
            <div style={boxStyle}>2</div>
            <div style={boxStyle}>3</div>
            <div style={boxStyle}>4</div>
            <div style={boxStyle}>5</div>
        </div>
    )
}

export default SiteFlexbox
