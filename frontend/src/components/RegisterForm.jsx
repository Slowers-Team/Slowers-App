import { useState } from 'react'
import { useTranslation } from 'react-i18next'

const RegisterForm = ({ createNewUser }) => {
    const [newUsername, setNewUsername] = useState('')
    const [newPassword, setNewPassword] = useState('')
    const [newEmail, setNewEmail] = useState('')
    const [newRole, setNewRole] = useState('')
    const [termsAccepted, setTermsAccepted] = useState(false)
    const { t, i18n } = useTranslation()

    const addUser = event => {
        event.preventDefault()
        if (!termsAccepted) {
            alert(t('error.acceptterms'))
            return
        }
        createNewUser({ 
            username: newUsername,
            password: newPassword,
            email: newEmail,
            role: newRole
        })

        setNewUsername('')
        setNewPassword('')
        setNewEmail('')
        setNewRole('')
        setTermsAccepted(false)

        alert(t('alert.usercreated'))
    }

    return (
        <div className='text-left'>
            <form onSubmit={addUser}>
                <div className='form-group mb-4'>
                  <label htmlFor="newUsernameInput">{t('user.data.username')}</label>
                  <input id="newUsernameInput" value={newUsername} placeholder={t('user.input.username')} onChange={event => setNewUsername(event.target.value)} className='form-control' required/>
                </div>
                <div className='form-group mb-4'>
                  <label htmlFor="newEmailInput">{t('user.data.email')}</label>
                  <input type="email" id="newEmailInput" value={newEmail} placeholder={t('user.input.email')} onChange={event => setNewEmail(event.target.value)} className='form-control' required/>
                </div>
                <div className='form-group mb-4'>
                  <label htmlFor="newPasswordInput">{t('user.data.password')}</label>
                  <input type="password" id="newPasswordInput" value={newPassword} placeholder={t('user.input.password')} onChange={event => setNewPassword(event.target.value)} className='form-control' required/>
                </div>
                <div className='form-group mb-4'>
                    <label htmlFor="roleSelector">{t('label.defaultrole')}:</label>
                    <div>
                      <input type="radio" className='btn-check' name="roleSelector" id="growerSelector" autoComplete="off" value="grower" onChange={event => setNewRole(event.target.value)} required />
                      <label className='btn btn-light' htmlFor="growerSelector" style={{ marginRight: "0.5rem" }}>{t('role.grower')}</label>
                      <input type="radio" className='btn-check' name="roleSelector" id="retailerSelector" autoComplete='off' value="retailer" onChange={event => setNewRole(event.target.value)} />
                      <label className='btn btn-light' htmlFor="retailerSelector"style={{ marginRight: "0.5rem" }}>{t('role.retailer')}</label>
                    </div>
                </div>
                <div className='form-check form-group mb-4'>
                  <input type='checkbox' className='form-check-input' id='termsCheckbox' checked={termsAccepted} onChange={() => setTermsAccepted(!termsAccepted)}/>
                  <label className='form-check-label' htmlFor="termsCheckbox">{t('label.iagreeto')} <a href='/terms' target="_blank" rel="noopener noreferrer">{t('label.terms')}</a></label>
                </div>
                <div>
                  <button type="submit" id="createNewUserButton" className='btn btn-primary'>{t('button.register')}</button>
                </div>
            </form>
        </div>
    )
}

export default RegisterForm
