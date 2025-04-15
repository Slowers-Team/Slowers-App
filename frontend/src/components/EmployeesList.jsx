import { useTranslation } from "react-i18next";

const ShowEmployee = ({ employee }) => {
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
    </tr>
  )
}

const EmployeesList = ({ employees }) => {
  const { t, i18n } = useTranslation()
  
  return (
    <div>
      <h2>{ t("menu.employees") }</h2>
      <table id='employeeList' className='table table-hover align-middle'>
        <tbody>
          {Array.isArray(employees) && employees.length > 0 ? (
          employees.map(employee => (
            <ShowEmployee 
              employee={employee}
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

export default EmployeesList
