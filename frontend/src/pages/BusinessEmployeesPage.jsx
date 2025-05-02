import { useState, useEffect } from "react"
import businessService from "../services/business"
import EmployeesList from "../components/EmployeesList"
import { useTranslation } from 'react-i18next'
import AddEmployeeForm from "../components/AddEmployeeForm"
import EditEmployees from "../components/EditEmployees"
import { Authenticator } from "../Authenticator"
import CenteredCard from "../components/CenteredCard"


const BusinessEmployeesPage = () => {
  const [employees, setEmployees] = useState([])
  const { t, i18n } = useTranslation()
  const designation = Authenticator.designation

  const employeeGetter = async () => {
    const business = await businessService.get();
    setEmployees(await businessService.getAllMembers(business.ID));
  };
  useEffect(() => {
    employeeGetter();
  }, []);
  

  return (
    <CenteredCard>
      {designation === 'owner' && (
        <AddEmployeeForm onEmployeeAdded={employeeGetter}/>
      )}
      <br/>
      {(designation === "owner") && (
        <EditEmployees employees={employees} onEmployeeEdited={employeeGetter} />
      )}              
      {(designation === "employee") && (
        <EmployeesList employees={employees} />
      )}
    </CenteredCard>
  )
}

export default BusinessEmployeesPage;
