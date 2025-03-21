import { useTranslation } from "react-i18next";
import { Container } from 'react-bootstrap'


const BusinessPage = () => {
  const { t, i18n } = useTranslation();


  return (
    <Container>
      <center>
      <div className="business-name-container">
      <h2>Business name</h2>
      <h5>Business type: Grower</h5>
      <h5>Phone number: +358 45 2789654</h5>
      <h5>Email: business@email.com</h5>
      <h5>Postal code: 00000</h5>
      <h5>City: Kerava</h5> 
      <h5>Created: 23.5.2021</h5>

      <br></br>
      <br></br>

      <h4>About business:</h4>
        <br></br>
        <div class="about-business-container">About business</div>
        <br></br>
      </div>
      </center>
      <br></br>


    </Container>
  );
};

export default BusinessPage;