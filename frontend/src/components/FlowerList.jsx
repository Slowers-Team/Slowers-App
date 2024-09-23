const FlowerList = ({ flowers, deleteFlower }) => {
  return (
    <table id="flowerList">
      <thead>
        <tr>
          <th>Name</th>
          <th>Latin name</th>
          <th>Added time</th>
        </tr>
      </thead>
      <tbody>
        {flowers.map(flower => {
          let addedTime = new Date(flower.added_time)

          let date = addedTime.toLocaleDateString('fi')
          let hour = addedTime.toLocaleString('fi', {hour: 'numeric'})
          let minute = addedTime.toLocaleString('fi', {minute: '2-digit'})
          let addedTimeStr = `${date} ${hour}:${minute}`

          return (
            <tr key={flower._id}>
              <td>{ flower.name }</td>
              <td><em>{ flower.latin_name }</em></td>
              <td>{ addedTimeStr }</td>
              <td>
                <button id="deleteFlowerButton" onClick={() => deleteFlower(flower)}>Delete</button>
              </td>
            </tr>
          )
        })}
      </tbody>
    </table>
  )
}

export default FlowerList
