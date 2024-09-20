import { useState } from 'react'

const RegisterForm = ({ createNewUser }) => {
    const [newUsername, setNewUsername] = useState('')
    const [newPassword, setNewPassword] = useState('')
    const [newEmail, setNewEmail] = useState('')

    const addUser = event => {
        event.preventDefault()
        createNewUser({ 
            username: newUsername,
            password: newPassword,
            email: newEmail 
        })

        setNewUsername('')
        setNewPassword('')
        setNewEmail('')
    }

    return (
        <div>
            <form onSubmit={addUser}>
                <div>
                    Username: <input id="newUsernameInput" value={newUsername} onChange={event => setNewUsername(event.target.value)} />
                </div>
                <div>
                    Password: <input type="password" id="newPasswordInput" value={newPassword} onChange={event => setNewPassword(event.target.value)} />
                </div>
                <div>
                    Email: <input id="newEmailInput" value={newEmail} onChange={event => setNewEmail(event.target.value)} />
                </div>
                <div>
                    <button id="createNewUserButton" type="submit">Register</button>
                </div>
            </form>
        </div>
    )
}

export default RegisterForm
