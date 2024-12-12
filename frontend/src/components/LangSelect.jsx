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
        🇬🇧 en
      </Nav.Link>
      <Nav.Link id="fi" href="#" onClick={() => changeLanguage('fi')}>
        🇫🇮 fi
      </Nav.Link>
    </>
  )
}

export default LangSelect
