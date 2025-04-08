import { useState } from "react"
import { useTranslation } from "react-i18next"
import businessService from "../services/business"

const AddEmployeeForm = () => {
  const [ email, setEmail] = useState('')

  const addEmployee = async (event) => {
    event.preventDefault()
    const designation = "employee"
    const business = await businessService.get()
    console.log(business.ID)
    const business_id = business.ID
    businessService.addMembership({email, business_id, designation})
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