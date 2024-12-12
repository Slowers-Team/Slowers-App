const formatTime = timeStr => {
  let dateObj = new Date(timeStr)

  let date = dateObj.toLocaleDateString('fi')
  let hour = dateObj.getHours().toString()
  let minute = dateObj.getMinutes().toString().padStart(2, '0')

  return `${date} ${hour}:${minute}`
}

export {
  formatTime,
}
