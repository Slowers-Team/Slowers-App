import { useTranslation } from "react-i18next";
import { Button } from "react-bootstrap"
import businessService from "../services/business"
import { useEffect } from "react";


const EmployeesList = ({ employees }) => {
  const { t, i18n } = useTranslation()

  const handleDeletion = (employees) => {
		businessService.deleteMembership(employees)
	}

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
              <Button variant="dark" onClick={() => handleDeletion(employees)} className="delete-button">{t('button.deletemember')}
									<i className="bi bi-trash"></i>
								</Button>

              {/* <td>
                <button type="submit" className='custom-button'>{t('button.deletemember')}</button>
              </td> */}
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
