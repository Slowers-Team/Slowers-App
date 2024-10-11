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
        <div className='text-left'>
            <form onSubmit={addUser}>
                <div className='form-group mb-4'>
                    <label htmlFor="newUsernameInput">Username</label>
                    <input id="newUsernameInput" value={newUsername} placeholder='username' onChange={event => setNewUsername(event.target.value)} className='form-control' required/>
                </div>
                <div className='form-group mb-4'>
                    <label htmlFor="newPasswordInput">Password</label>
                    <input type="password" id="newPasswordInput" value={newPassword} placeholder='password' onChange={event => setNewPassword(event.target.value)} className='form-control' required/>
                </div>
                <div className='form-group mb-4'>
                    <label htmlFor="newEmailInput">Email address</label>
                    <input type="email" id="newEmailInput" value={newEmail} placeholder='email' onChange={event => setNewEmail(event.target.value)} className='form-control' required/>
                </div>
                <div>
                    <button type="submit" id="createNewUserButton" className='btn btn-primary' >Register</button>
                </div>
            </form>
        </div>
    )
}

export default RegisterForm
