import { useState } from "react"
import { useTranslation } from "react-i18next"
import businessService from "../services/business"


const ShowEmployee = ({ employee, handleEditEmployee, handleDeletion }) => {
  console.log(employee)
  return (
    <tr>
      <td>{employee[0]}</td> 
      <td>{employee[1]}</td>
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
  return (
    <form onSubmit={handleDeletion}>
      <input type="hidden" name="email" value={employee[0]} />
      <button type="submit" className="custom-button">Delete</button>
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
  const handleDeletion = (employee) => {
    console.log("Trying to delete employee:", employee);

    businessService.deleteMembership(employee)
      .then(() => {
        console.log("Fetching updated employee list...");
        employeeGetter();
      })
      .catch(error => {
        //täältä tulee vielä error, korjaantuu ma tai ti!
        console.error("Error deleting employee:", error);
      });
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