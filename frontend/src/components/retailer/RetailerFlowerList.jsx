import '../../layouts/Retailer.css'
import { useTranslation } from 'react-i18next'

const RetailerFlowerList = ({ flowers }) => {
  const { t, i18n } = useTranslation()
  return (
    <div className="retailerFlowerList">
      <table id="retailerFlowerList">
        <thead>
          <tr>
            <th>{t('flower.data.name')}</th>
            <th>{t('flower.data.latinname')}</th>
            <th>{t('flower.data.addedtime')}</th>
            <th>{t('flower.data.grower')}</th>
          </tr>
        </thead>
        <tbody>
          {flowers.map(flower => {
            let addedTime = new Date(flower.added_time)

            let date = addedTime.toLocaleDateString('fi')
            let hour = addedTime.toLocaleString('fi', { hour: 'numeric' })
            let minute = addedTime.toLocaleString('fi', { minute: '2-digit' })
            let addedTimeStr = `${date} ${hour}:${minute}`

            return (
              <tr key={flower._id}>
                <td>{flower.name}</td>
                <td>
                  <em>{flower.latin_name}</em>
                </td>
                <td>{addedTimeStr}</td>
                <td>{flower.grower_email}</td>
              </tr>
            )
          })}
        </tbody>
      </table>
    </div>
  )
}

export default RetailerFlowerList
