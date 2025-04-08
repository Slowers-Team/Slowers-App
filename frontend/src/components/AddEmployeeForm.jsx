import { useState } from "react"
import { useTranslation } from "react-i18next"
import businessService from "../services/business"

const AddEmployeeForm = () => {
  const { t, i18n } = useTranslation()
  const [ email, setEmail] = useState('')
  const [errorMessage, setErrorMessage] = useState('')

  const addEmployee = async (event) => {
    try {
      event.preventDefault()
      const designation = "employee"
      const business = await businessService.get()
      const business_id = business.ID
      console.log(business_id)
      console.log(email)
      await businessService.addMembership({"UserEmail": email, "BusinessID": business_id, "Designation": designation})
    } catch (error) {
        setErrorMessage(t('error.errocurred'))
    }
  }

  return (
    <div>
      <h2>Add employee</h2>
      <form onSubmit={addEmployee}>
        <table>
          <tbody>
            <tr>
              <td>
                Email:
              </td>
              <td>
                <input
                  className="form-control"
                  type="email"
                  value={email}
                  placeholder="Enter email"
                  onChange={event => setEmail(event.target.value)}
                  required
                />
              </td>
              <td>
                <button type="submit" className="custom-button">Add employee</button>
              </td>
            </tr>
          </tbody>
        </table>
      </form>
    </div>
  )
}

export default AddEmployeeForm