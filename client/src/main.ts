import { createApp } from "vue";
import App from "./App.vue";
import store from "./store/store";

export const backendAddress = "https://logogen.mbaraa.fun"

createApp(App)
    .use(store)
    .mount('#app')
