import { useTranslation } from "react-i18next";


const EmployeesList = ({ employees }) => {
  const { t, i18n } = useTranslation()

  return (
    <div>
      Here is the list of employees
    </div>
  )
}

export default EmployeesList
