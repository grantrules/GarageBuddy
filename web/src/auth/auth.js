const unauthorized = () => { throw new Error('Unauthorized'); };

const authWithUser = (func) => ({ session }, ...args) => ((session
    && session.teacherId) ? func(session.teacherId, ...args) : unauthorized());

const auth = (func) => authWithUser((_, ...args) => func(...args));

export { auth, authWithUser };