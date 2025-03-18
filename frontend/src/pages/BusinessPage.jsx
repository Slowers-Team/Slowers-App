import { useTranslation } from "react-i18next";
import { Container } from 'react-bootstrap'


const BusinessPage = () => {
  const { t, i18n } = useTranslation();


  return (
    <Container>
      <center>
      <div className="business-name-container">
      <h1>Business name</h1>
      <h2>Business type: Grower</h2>
      <h2>Owner: Owner name</h2>
      </div>
      <br></br>

      <div className="members-container">
      <h3>Members:</h3>
      <ul>
        <li>Matti</li>
        <li>Teppo</li>
        <li>Maija</li>
      </ul> 
      <button className='add-member-button'>+Add member</button>
      </div>
      <br></br>

      <div className="about-business-container">
        <h3>About business:</h3>
        <br></br>
        <textarea id="businessinfo" name="businessinfo" rows="4" cols="50">About business</textarea>
        <br></br>
        <button className='add-member-button'>Update</button>
      </div>
      </center>
    </Container>
  );
};

export default BusinessPage;