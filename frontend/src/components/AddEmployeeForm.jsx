import { useState } from "react"
import { useTranslation } from "react-i18next"
import businessService from "../services/business"

const AddEmployeeForm = ({ onEmployeeAdded }) => {
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
      setEmail("");
      onEmployeeAdded();
    } catch (error) {
        setErrorMessage(t('error.errocurred'))
    }
  }

  return (
    <div>
      <h2>{t('title.addemployees')}</h2>
      <form onSubmit={addEmployee}>
        <table>
          <tbody>
            <tr>
              <td>
                {t('user.data.email')}
              </td>
              <td>
                <input
                  className="form-control"
                  id="employeeEmailInput"
                  type="email"
                  value={email}
                  placeholder={t('user.input.email')}
                  size="40"
                  onChange={event => setEmail(event.target.value)}
                  required
                />
              </td>
              <td>
                <button type="submit" className="custom-button" id="addEmployeeButton">{t('button.addemployee')}</button>
              </td>
            </tr>
          </tbody>
        </table>
      </form>
    </div>
  )
}

export default AddEmployeeForm