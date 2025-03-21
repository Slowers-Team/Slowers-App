import { Nav } from "react-bootstrap"

import i18n from "../i18n"

const LangSelect = () => {
  const changeLanguage = lang => {
    document.cookie = `lang=${lang}; expires=${new Date(Date.now().valueOf() + 2592000000).toUTCString()}; path=/`
    i18n.changeLanguage(lang)
  }

  return (
    <>
      <Nav.Link id="en" href="#" onClick={() => changeLanguage('en')}>
        🇬🇧 <span className="lang-link">English</span>
      </Nav.Link>
      <Nav.Link id="fi" href="#" onClick={() => changeLanguage('fi')}>
        🇫🇮 <span className="lang-link">Suomi</span>
      </Nav.Link>
      <Nav.Link id="sv" href="#" onClick={() => changeLanguage('sv')}>
        🇸🇪 <span className="lang-link">Svenska</span>
      </Nav.Link>
    </>
  )
}

export default LangSelect
