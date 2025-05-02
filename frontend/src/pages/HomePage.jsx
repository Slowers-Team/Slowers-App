import { useTranslation } from "react-i18next";
import CenteredCard from "../components/CenteredCard";


const HomePage = () => {
  const { t, i18n } = useTranslation();

  return (
    <CenteredCard>
      <h2>{t('title.welcome')}</h2>
    </CenteredCard>
  );
};

export default HomePage;
