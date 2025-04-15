import { useEffect, useState } from "react"
import { useTranslation } from "react-i18next"
import businessService from "../services/business"
import userService from "../services/users"


const ShowEmployee = ({ employee, handleEditEmployee, handleDeletion, currentUser }) => {
  console.log(employee)
  
  return (
    <tr>
      <td>{employee[0]}</td> 
      <td>{employee[1]}</td>
      {!(employee[0] === currentUser.Email) && (
        <>
        <td>
          <EditEmployeeForm employee={employee} handleEditEmployee={handleEditEmployee} />
        </td>
        <td>
          <DeleteEmployeeForm employee={employee} handleDeletion={handleDeletion} />
        </td>
        </>
        )}
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

const DeleteEmployeeForm = ({ employee, handleDeletion }) => {
  const { t, i18n } = useTranslation()
  return (
    <form onSubmit={handleDeletion}>
      <input type="hidden" name="email" value={employee[0]} />
      <button type="submit" className="custom-delete-button" color="red">
      <i className="bi bi-trash3-fill"> </i>
      {t("button.deletemember")}</button>
    </form>
  )
}

const EditEmployees = ({ employees, onEmployeeEdited }) => {
  const { t, i18n } = useTranslation()
  const [currentUser, setCurrentUser] = useState(null)

  useEffect(() => {
    const fetchUser = async () => {
      const user = await userService.get()
      setCurrentUser(user)
    }
    fetchUser()
  }, [])

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
  const handleDeletion = async (event) => {
    event.preventDefault()
    console.log("Trying to delete employee");
    
    const currentUser = await userService.get()

    const formData = new FormData(event.target)
    const userEmail = formData.get("email")
    const businessId = (await businessService.get()).ID

    if (userEmail === currentUser.Email) {
      alert(t("alert.cannotdeleteyourself"))
      return
    }

    const confirmed = window.confirm((t("alert.deletemember")) + userEmail + "?") 

    if (!confirmed) {
      return
    }

    try {
      await businessService.deleteMembership(userEmail, businessId)
      console.log("Deleted:", userEmail);
      onEmployeeEdited()
    } catch (error) {
      console.error("Error deleting employee:", error);
    }
	}

  return (
    <div>
      <h2>Edit employees</h2>
      <table id='employeeList' className='table table-hover align-middle'>
        <tbody>
          {Array.isArray(employees) && employees.length > 0 ? (
          employees.map(employee => (
            <ShowEmployee 
              employee={employee}
              currentUser={currentUser}
              handleEditEmployee={handleEditEmployee}
              handleDeletion={handleDeletion}
              key={employee[0]}
            />
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