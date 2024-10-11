import i18n from "../i18n"

const LangSelect = () => {
  const changeLanguage = lang => {
    document.cookie = `lang=${lang}; expires=${new Date(Date.now().valueOf() + 2592000000).toUTCString()}; path=/`
    i18n.changeLanguage(lang)
  }

  return (
    <div style={{position: "absolute", top: "0", right: "0", padding: "8px"}}>
      <a href="#" onClick={() => changeLanguage('en')} style={{paddingRight: "0.8rem"}}>
        en
      </a>
      <a href="#" onClick={() => changeLanguage('fi')} style={{paddingRight: "0.8rem"}}>
        fi
      </a>
    </div>
  )
}

export default LangSelect
