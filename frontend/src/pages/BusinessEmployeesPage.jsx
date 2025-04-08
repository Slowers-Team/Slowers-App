import { useState, useEffect } from "react"
//import userService from "../services/users"
import businessService from "../services/business"
import EmployeesList from "../components/EmployeesList"
import { useTranslation } from 'react-i18next'


const BusinessEmployeesPage = () => {
  const [employees, setEmployees] = useState({})
  const { t, i18n } = useTranslation()
  //const [errorMessage, setErrorMessage] = useState("")

  useEffect(() => {
    businessService.getAllMembers().then((employees) => setEmployees(employees))
  }, []);

  return (
    <div className="m-3">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-8">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              <h2 className="mx-3 my-3">{ t("menu.employees") }</h2>
              <EmployeesList employees={employees} />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default BusinessEmployeesPage;
