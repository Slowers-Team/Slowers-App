const formatTime = (timeStr) => {
  let dateObj = new Date(timeStr);

  let date = dateObj.toLocaleDateString("fi");
  let hour = dateObj.getHours().toString();
  let minute = dateObj.getMinutes().toString().padStart(2, "0");

  return `${date} ${hour}:${minute}`;
};

const validateEmail = (email) => {
  // Use same regex as in backend (utils.go)
  const re = /^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$/;
  return re.test(email);
};

const validateBusinessIdCode = (idCode) => {
  // Use same regex as in backend (utils.go)
  const re = /^\d{7}-\d$/;
  return re.test(idCode);
};

const validatePostalCode = (postalCode) => {
  // Use same regex as in backend (utils.go)
  const re = /^\d{5}$/;
  return re.test(postalCode);
};

const validatePhoneNumber = (phoneNumber) => {
  // Use same regex as in backend (utils.go)
  const re = /^\d{10,13}$/;
  return re.test(phoneNumber);
};

export {
  formatTime,
  validateEmail,
  validateBusinessIdCode,
  validatePostalCode,
  validatePhoneNumber,
};
