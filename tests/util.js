function makeid(length) {
  let result = [];
  const characters =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  for (let i = 0; i < length; i++) {
    result.push(
      characters.charAt(Math.floor(Math.random() * characters.length))
    );
  }
  return result.join("");
}

const setParam = (context, events, done) => {
  context.vars["payload"] = makeid(Math.floor(Math.random() * 10) + 1)
  return done();
};

exports.setParam = setParam;
