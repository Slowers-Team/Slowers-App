// import { useState, useEffect } from "react";
// import userService from "../services/users";
// import { useTranslation } from "react-i18next";
// import { Authenticator } from "../Authenticator";
// import UserPage from "../pages/UserPage";

// const CreateBusinessForm = () => {
//   const [businessName, setBusinessName] = useState('')
//   const [type, setType] = useState ('')
//   const { t, i18n } = useTranslation()
//   const [newRole, setNewRole] = useState('')

//   const handleCreateBusiness = async (event) => {
//     event.preventDefault()
//     const updatedRole = type === "grower" ? "growerowner" : "retailerowner";
//     userService.setRole(updatedRole).then((_) => {
//       Authenticator.setRole(updatedRole)

//     })
//    }


//   return (
//     <div>
//       <form onSubmit={handleCreateBusiness}>
//         <div>
//           <table>
//             <tbody>
//               <tr>
//                 <td>{t('businessform.fieldname.businessname')}</td>
//                 <td>
//                   <input 
//                     className="form-control"
//                     value={businessName}
//                     placeholder={t('businessform.input.businessname')}
//                     onChange={event => setBusinessName(event.target.value)}
//                     style={{ width: "400px"}}
//                   />
//                 </td>
//               </tr>
//               <tr>
//                 <td>{t('businessform.fieldname.businesstype')}</td>
//                 <td>
//                   <label className='btn btn-outline-secondary' style={{ marginRight: "0.5rem" }} >
//                     <input type="radio" className='btn-check' name="typeSelector" value="grower" onChange={event => setType(event.target.value)} />{t('button.grower')}
//                   </label>
//                   <label className='btn btn-outline-secondary' style={{ marginRight: "0.5rem" }} >
//                   <input type="radio" className='btn-check' name="typeSelector" value="retailer"  onChange={event => setType(event.target.value)} />{t('button.retailer')}
//                   </label>
//                 </td>
//               </tr>
//             </tbody>
//           </table>
//           <br/>
//           <button type="submit" className='custom-button'>{t('button.createbusiness')}</button>
//         </div>
//       </form>
//     </div> 
//   )
// }

// export default CreateBusinessForm