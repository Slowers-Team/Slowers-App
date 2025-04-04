import { useTranslation } from "react-i18next"


const BusinessInfo = ({ business }) => {
  const { t, i18n } = useTranslation()
  const dateOnly = business?.CreatedAt ? new Date(business.CreatedAt).toISOString().split("T")[0] : "N/A"

  let type
  if (business.Type === "grower") {
    type = t('role.grower')
  } else if (business.Type === "retailer") {
    type = t('role.retailer')
  }
  
  let delivery
  if (business.Delivery === "yes") {
    delivery = t('businessform.input.yesdelivery')
  } else if (business.Delivery === "no") {
    delivery = t('businessform.input.nodelivery')
  }

  return (
    <div>
      <h2>{ business.BusinessName }</h2>
      <br/>
      <h5>{t('businessform.fieldname.businesstype')}: { type } </h5>
      <h5>{t('businessform.fieldname.phonenumber')}: { business.PhoneNumber }</h5>
      <h5>{t('businessform.fieldname.email')}: { business.Email }</h5>
      <h5>{t('businessform.fieldname.postalcode')}: { business.PostalCode }</h5>
      <h5>{t('businessform.fieldname.city')}: { business.City }</h5>
      <h5>{t('businessform.fieldname.address')}: { business.Address }</h5>
      <h5>{t('businessform.fieldname.businessidcode')}: { business.BusinessIdCode }</h5>
      <h5>{t('businessform.fieldname.created')}: { dateOnly }</h5>
      {business.Type === "grower" && (
        <h5>{t('businesspage.delivery')}: { delivery }</h5>
      )}
      <br/>
      <h4>{t('businesspage.additionalinfo')}:</h4>
      <h5>{ business.AdditionalInfo }</h5>
    </div>
  )
}

export default BusinessInfo