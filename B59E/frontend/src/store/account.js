let account = {
    namespaced: true,
    state: {
        token: "",
        user: "",
    },
    mutations: {
        toggleLoginToken(state, token) {
            console.log("set login!!!")
            state.token = token;
        },
        setUserInfo(state, user) {
            console.log("set user!!!")
            state.user = user;
        }
    },
    getters: {
        token: state => state.token,
        user: state => state.user,
    }
}

export default account