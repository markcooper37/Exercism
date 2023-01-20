export const hey = (message) => {
  let isUpperCase = /[a-zA-Z]/.test(message) && message === message.toUpperCase();
  let isQuestion = message.trim().endsWith('?');
  if (isUpperCase && isQuestion) {
    return 'Calm down, I know what I\'m doing!';
  } else if (isUpperCase) {
    return 'Whoa, chill out!';
  } else if (isQuestion) {
    return 'Sure.';
  } else if (/\S/.test(message)) {
    return 'Whatever.';
  } else {
    return 'Fine. Be that way!';
  };
};
