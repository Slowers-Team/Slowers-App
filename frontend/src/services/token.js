const fetchToken = () => {
    return localStorage.getItem('token')
}

export default { fetchToken }