const FlowerForm = ({ event, name, handleFlowerNameChange, latin_name, handleFlowerLatinNameChange }) => {
  return (
    <div>
      <form onSubmit={event}>
        <div>
          Name: <input value={name} onChange={handleFlowerNameChange} />
        </div>
        <div>
          Latin name: <input value={latin_name} onChange={handleFlowerLatinNameChange}/>
        </div>
        <div>
          <button type='submit'>Save</button>
        </div>
      </form>
    </div>
  )
}

export default FlowerForm
