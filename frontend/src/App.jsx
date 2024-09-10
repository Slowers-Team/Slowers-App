import axios from 'axios'
import { useEffect, useState } from 'react'
import './App.css'

const App = () => {
  const [flowers, setFlowers] = useState([])

  useEffect(() => {
    axios
      .get('/api/flowers')
      .then(response => setFlowers(response.data))
  }, [])

  return (
    <>
      {flowers && <FlowerList flowers={flowers} />}
    </>
  )
}

const FlowerList = ({ flowers }) => {
  return (
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Latin name</th>
          <th>Added time</th>
        </tr>
      </thead>
      <tbody>
        {flowers.map(flower => (
          <tr key={flower.id}>
            <td>{ flower.name }</td>
            <td><em>{ flower.latin_name }</em></td>
            <td>{ new Date(flower.added_time).toDateString() }</td>
          </tr>
        ))}
      </tbody>
    </table>
  )
}

export default App
