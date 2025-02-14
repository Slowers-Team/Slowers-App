import { Nav } from "react-bootstrap"

import i18n from "../i18n"

const LangSelect = () => {
  const changeLanguage = lang => {
    document.cookie = `lang=${lang}; expires=${new Date(Date.now().valueOf() + 2592000000).toUTCString()}; path=/`
    i18n.changeLanguage(lang)
  }

  return (
    <>
      <Nav.Link href="#" onClick={() => changeLanguage('en')}>
        🇬🇧 <span className="lang-link">English</span>
      </Nav.Link>
      <Nav.Link href="#" onClick={() => changeLanguage('fi')}>
        🇫🇮 <span className="lang-link">Suomi</span>
      </Nav.Link>
    </>
  )
}

export default LangSelect
