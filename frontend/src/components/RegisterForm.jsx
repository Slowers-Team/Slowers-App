import { useState } from 'react'
import { useTranslation } from 'react-i18next'

const RegisterForm = ({ createNewUser }) => {
    const [newUsername, setNewUsername] = useState('')
    const [newPassword, setNewPassword] = useState('')
    const [newEmail, setNewEmail] = useState('')
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
            email: newEmail 
        })

        setNewUsername('')
        setNewPassword('')
        setNewEmail('')
        setTermsAccepted(false)
    }

    return (
        <div>
            <form onSubmit={addUser}>
                <div>
                  <label htmlFor="newUsernameInput">{t('user.data.username')}:</label>
                  <input id="newUsernameInput" value={newUsername} onChange={event => setNewUsername(event.target.value)} />
                </div>
                <div>
                  <label htmlFor="newPasswordInput">{t('user.data.password')}:</label>
                  <input type="password" id="newPasswordInput" value={newPassword} onChange={event => setNewPassword(event.target.value)} />
                </div>
                <div>
                  <label htmlFor="newEmailInput">{t('user.data.email')}:</label>
                  <input type="email" id="newEmailInput" value={newEmail} onChange={event => setNewEmail(event.target.value)} />
                </div>
                <div>
                  <input type='checkbox' id='termsCheckbox' checked={termsAccepted} onChange={() => setTermsAccepted(!termsAccepted)}/>
                  <label htmlFor="termsCheckbox">{t('label.iagreeto')} <a href='/terms' target="_blank" rel="noopener noreferrer">{t('label.terms')}</a></label>
                </div>
                <div>
                  <button id="createNewUserButton" type="submit">{t('button.register')}</button>
                </div>
            </form>
        </div>
    )
}

export default RegisterForm
