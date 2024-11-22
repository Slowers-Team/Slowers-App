import AddImage from './AddImage'

const FlowerImageTab = ({ isGrower, flower }) => {

    return (
        <div>
            {isGrower
            ? <AddImage entity={flower}/>
            : <></> }
        </div>
    )
}

export default FlowerImageTab
