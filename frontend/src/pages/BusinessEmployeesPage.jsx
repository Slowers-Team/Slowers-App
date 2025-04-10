import { useState, useEffect } from "react"
//import userService from "../services/users"
import businessService from "../services/business"
import EmployeesList from "../components/EmployeesList"
import { useTranslation } from 'react-i18next'
import AddEmployeeForm from "../components/AddEmployeeForm"



const BusinessEmployeesPage = () => {
  const [employees, setEmployees] = useState({})
  const { t, i18n } = useTranslation()
  //const [errorMessage, setErrorMessage] = useState("")

//   useEffect(() => {
//     businessService.get().then((business) => setEmployees(business))
//   }, []);

  return (
    <div className="m-3">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-8">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              <h2>{ t("menu.employees") }</h2>
              <EmployeesList employees={employees} />
              <br/>
              <AddEmployeeForm />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default BusinessEmployeesPage;