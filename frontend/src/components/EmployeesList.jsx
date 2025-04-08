import { useTranslation } from "react-i18next";
import { Button } from "react-bootstrap"
import businessService from "../services/business"
import { useEffect } from "react";
import { useState } from "react";


const EmployeesList = ({ employees }) => { //tai = () => {
  const { t, i18n } = useTranslation()

  // kovakoodattu lista ihan vaan että napin saa näkymään
  // lähtee kun branchi on valmis
  
  // const [employees, setEmployees] = useState([
  //   [1, "John Doe"],
  //   [2, "Jane Smith"],
  //   [3, "Michael Johnson"]
  // ]);

  // tämä ja Button-kohta alhaalla uusia. Testaa miten toimii!
  const handleDeletion = (employee) => {
		businessService.deleteMembership(employee)
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
              {/* <td>
              <Button onClick={() => handleDeletion(employee)} >{t('button.deletemember')}
								</Button>
                </td> */}
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
