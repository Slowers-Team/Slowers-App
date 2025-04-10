import { useTranslation } from "react-i18next";


const EmployeesList = ({ employees }) => {
  const { t, i18n } = useTranslation()

  return (
    <div>
      Here is the list of employees
      <table id='employeeList' className='table table-hover align-middle'>
        <tbody>
          {Array.isArray(employees) && employees.length > 0 ? (
          employees.map(employees => (
            <tr key={employees[0]}>
              <td>{employees[0]}</td> 
              <td>{employees[1]}</td>
            </tr>
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

export default EmployeesList
