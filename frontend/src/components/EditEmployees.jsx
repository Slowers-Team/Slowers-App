import { useTranslation } from "react-i18next"
import businessService from "../services/business"


const ShowEmployee = ({ employee, handleEditEmployee }) => {
  console.log(employee)
  return (
    <tr>
      <td>{employee[0]}</td> 
      <td>{employee[1]}</td>
      <td>
        <EditEmployeeForm employee={employee} handleEditEmployee={handleEditEmployee} />
      </td>
    </tr>
  )
}

const EditEmployeeForm = ({ employee, handleEditEmployee }) => {
  let buttonLabel = ""
  if (employee[1] === "owner") {
    buttonLabel = "Change to employee"
  } else if (employee[1] === "employee") {
    buttonLabel = "Change to owner"
  }

  return (
    <form onSubmit={handleEditEmployee}>
      <input type="hidden" name="email" value={employee[0]} />
      <input type="hidden" name="designation" value={employee[1]} />
      <button type="submit" className="custom-button">{buttonLabel}</button>
    </form>
  )
}

const EditEmployees = ({ employees, onEmployeeEdited }) => {
  const { t, i18n } = useTranslation()

  const handleEditEmployee = async (event) => {
    event.preventDefault()
    console.log('handleEditEmployee')
    const formData = new FormData(event.target)
    const userEmail = formData.get("email")
    const businessId = (await businessService.get()).ID
    const designation = (formData.get("designation") === "employee") ? "owner" : "employee"
    await businessService.editMember({userEmail, businessId, designation})
    onEmployeeEdited()
  }

  return (
    <div>
      <h2>Edit employees</h2>
      <table id='employeeList' className='table table-hover align-middle'>
        <tbody>
          {Array.isArray(employees) && employees.length > 0 ? (
          employees.map(employee => (
            <ShowEmployee employee={employee} handleEditEmployee={handleEditEmployee} key={employee[0]} />
          ))
        ) : (
        <tr>
          <td colSpan="2">No data available</td>
        </tr>
        )}
        </tbody>
      </table>
    </div>
  )
}

export default EditEmployees