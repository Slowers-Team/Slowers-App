const translations = require('../src/translations.json');

function compareStructures(obj1, obj2) {
    const keys1 = Object.keys(obj1);
    const keys2 = Object.keys(obj2);
  
    if (keys1.length !== keys2.length) {
      return false;
    }
  
    for (let key of keys1) {
      if (!keys2.includes(key)) {
        return false;
      }
  
      if (typeof obj1[key] === 'object' && typeof obj2[key] === 'object') {
        if (!compareStructures(obj1[key], obj2[key])) {
          return false;
        }
      }
    }
  
    return true;
  }
  
test('All translation objects have equal structures', () => {
    const enTranslation = translations.en.translation;
    const svTranslation = translations.sv.translation;
    const fiTranslation = translations.fi.translation;

    expect(compareStructures(enTranslation, svTranslation)).toBe(true);
    expect(compareStructures(enTranslation, fiTranslation)).toBe(true);
});