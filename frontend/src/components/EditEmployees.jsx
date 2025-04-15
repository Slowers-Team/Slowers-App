import { useState } from "react"
import { useTranslation } from "react-i18next"
import businessService from "../services/business"


const ShowEmployee = ({ employee, handleEditEmployee, handleDeletion }) => {
  console.log(employee)
  const { t, i18n } = useTranslation()
  let designation
  if (employee[1] === "owner") {
    designation = t("designation.owner")
  } else if (employee[1] === "employee") {
    designation = t("designation.employee")
  }
  return (
    <tr>
      <td>{employee[0]}</td> 
      <td>{designation}</td>
      <td>
        <EditEmployeeForm employee={employee} handleEditEmployee={handleEditEmployee} />
      </td>
      <td>
        <DeleteEmployeeForm employee={employee} handleDeletion={handleDeletion} />
      </td>
    </tr>
  )
}

const EditEmployeeForm = ({ employee, handleEditEmployee }) => {
  const { t, i18n } = useTranslation()

  let buttonLabel = ""
  if (employee[1] === "owner") {
    buttonLabel = t("button.changetoemployee")
  } else if (employee[1] === "employee") {
    buttonLabel = t("button.changetoowner")
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
    const formData = new FormData(event.target)
    const userEmail = formData.get("email")
    const businessId = (await businessService.get()).ID

    try {
      await businessService.deleteMembership(userEmail, businessId)
      console.log("Deleted:", userEmail);
      console.log("Fetching updated employee list...");
      onEmployeeEdited()
    } catch (error) {
      console.error("Error deleting employee:", error);
    }
	}

  return (
    <div>
      <h2>{t("title.editemployees")}</h2>
      <table id='employeeList' className='table table-hover align-middle'>
        <tbody>
          {Array.isArray(employees) && employees.length > 0 ? (
          employees.map(employee => (
            <ShowEmployee 
              employee={employee}
              handleEditEmployee={handleEditEmployee}
              handleDeletion={handleDeletion}
              key={employee[0]}
            />
          ))
        ) : (
        <tr>
          <td colSpan="2">{t("error.nodataavailable")}</td>
        </tr>
        )}
        </tbody>
      </table>
    </div>
  )
}

export default EditEmployees