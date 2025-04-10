import { useTranslation } from "react-i18next";
import { Button } from "react-bootstrap"
import businessService from "../services/business"
import { useEffect } from "react";
import { useState } from "react";


const EmployeesList = () => {
  const { t, i18n } = useTranslation()
  //const [employeeList, setEmployeeList] = useState(employees)
  const [employees, setEmployees] = useState([])

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

  return (
    <div>
      Here is the list of employees
      <table id='employeeList' className='table table-hover align-middle'>
        <tbody>
          {Array.isArray(employees) && employees.length > 0 ? (
          employees.map(employee => (
            <tr key={employee[0]}>
              <td>{employee[0]}</td>
              <td>{employee[1]}</td>
              <td>
                <button type="submit" className='custom-button' onClick={() => handleDeletion(employee)} >{t('button.deletemember')}</button>
              </td>
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
