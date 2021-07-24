import {createStore} from 'vuex';
import ThemeChooser from "../utils/ThemeChooser.js";
import Logo from "../models/Logo.js";
import Text from "../models/Text.js";

const store = createStore({
    state: {
        theme_chooser: new ThemeChooser(),
        logo: new Logo("color",
            new Text("Type to see magic!", "", "ProductSans", "1.58em", "60.625%"),
            "vertical"
        )
    },
    mutations: {
        CHANGE_THEME(state) {
            state.theme_chooser.changeTheme();
        },

        SET_LOGO(state, newLogo) {
            state.logo = newLogo;
        },
    },
    actions: {
        changeTheme: ({commit}) => {
            commit("CHANGE_THEME");
        },

        setLogo: ({commit}, newLogo) => {
            commit("SET_LOGO", newLogo);
        },
    },
    getters: {
        getTheme(state) {
            return state.theme_chooser.getTheme();
        },

        getLogo(state) {
            return state.logo;
        }
    }
})

export default store;