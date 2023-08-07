export const clean = (text) => {
  var digits = text.replace(/[(). +\-]/g, "")
  if (/[a-zA-Z]/.test(digits)) {
    throw new Error('Letters not permitted')
  } else if (!/^\d+$/.test(digits)) {
    throw new Error('Punctuations not permitted')
  } else if (digits.length < 10) {
    throw new Error('Incorrect number of digits')
  } else if (digits.length > 11) {
    throw new Error('More than 11 digits')
  } else if (digits.length === 11) {
    if (digits[0] != '1') {
      throw new Error('11 digits must start with 1')
    } else {
      digits = digits.slice(1)
    }
  }
  if (digits[0] === '0') {
    throw new Error('Area code cannot start with zero')
  } else if (digits[0] === '1') {
    throw new Error('Area code cannot start with one')
  } else if (digits[3] === '0') {
    throw new Error('Exchange code cannot start with zero')
  } else if (digits[3] === '1') {
    throw new Error('Exchange code cannot start with one')
  }
  return digits
};
