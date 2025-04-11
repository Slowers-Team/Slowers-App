import { useState } from "react"
import { useTranslation } from "react-i18next"

const ShowEmployee = ({ employee, handleEditEmployee }) => {
  return (
    <tr>
      <td>{employee[3]}</td> 
      <td>{employee[5]}</td>
      <td>
        <EditEmployeeForm employee={employee} handleEditEmployee={handleEditEmployee} />
      </td>
    </tr>
  )
}

const EditEmployeeForm = ({ employee, handleEditEmployee }) => {
  let buttonLabel = ""
  if (employee[5] === "owner") {
    buttonLabel = "Change to employee"
  } else if (employee[5] === "employee") {
    buttonLabel = "Change to owner"
  }

  return (
    <form onSubmit={handleEditEmployee}>
      <button type="submit" className="custom-button">{buttonLabel}</button>
    </form>
  )
}

const EditEmployees = ({ employees }) => {
  const { t, i18n } = useTranslation()

  const handleEditEmployee = (event) => {
    console.log('handleEditEmployee')
  }

const DeleteEmployees = ({ employees }) => {
  const { t, i18n } = useTranslation()
  //const [employees, setEmployees] = useState() //mitä pitää laittaa

  const employeeGetter = async () => {
    const business = await businessService.get();
    setEmployees(await businessService.getAllMembers(business.ID));
  }

  useEffect(() => {
    employeeGetter();
  }, []);
  
  const handleDeletion = (employee) => {
    console.log("Trying to delete employee:", employee);
  
    businessService.deleteMembership(employee)
      .then(() => {
        console.log("Fetching updated employee list...");
        employeeGetter();
      })
      .catch(error => {
        console.error("Error deleting employee:", error);
      });
  }
}

  return (
    <div>
      <h2>Edit employees</h2>
      <table id='employeeList' className='table table-hover align-middle'>
        <tbody>
          {Array.isArray(employees) && employees.length > 0 ? (
          employees.map(employee => (
            <ShowEmployee employee={employee} handleEditEmployee={handleEditEmployee} key={employee[3]} />
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