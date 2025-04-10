import { useState, useEffect } from "react"
import businessService from "../services/business"
import EmployeesList from "../components/EmployeesList"
import { useTranslation } from 'react-i18next'
import AddEmployeeForm from "../components/AddEmployeeForm"
import EditEmployees from "../components/EditEmployees"


const BusinessEmployeesPage = () => {
  const [employees, setEmployees] = useState([])
  const { t, i18n } = useTranslation()
  const employeeGetter = async () => {
    const business = await businessService.get();
    setEmployees(await businessService.getAllMembers(business.ID));
  };
  useEffect(() => {
    employeeGetter();
  }, []);
  


  return (
    <div className="m-3">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-8">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              <AddEmployeeForm onEmployeeAdded={employeeGetter}/>
              <br/>
              <h2>{ t("menu.employees") }</h2>
              <EmployeesList employees={employees} />              
              <EditEmployees employees={employees} />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default BusinessEmployeesPage;
